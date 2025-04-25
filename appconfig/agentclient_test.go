package appconfig

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/Arthur1/openfeature-provider-go-aws-appconfig/internal/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestNewAgentClient(t *testing.T) {
	t.Parallel()
	t.Run("default", func(t *testing.T) {
		t.Parallel()
		client := NewAgentClient()
		assert.Equal(t, http.DefaultClient, client.httpClient)
		assert.Equal(t, "http://localhost:2772", client.baseURL)
	})

	t.Run("WithHTTPClientOption", func(t *testing.T) {
		t.Parallel()
		httpClient := new(http.Client)
		client := NewAgentClient(WithHTTPClientOption(httpClient))
		assert.Equal(t, httpClient, client.httpClient)
	})

	t.Run("WithBaseURLOption", func(t *testing.T) {
		t.Parallel()
		client := NewAgentClient(WithBaseURLOption("http://localhost:8080"))
		assert.Equal(t, "http://localhost:8080", client.baseURL)
	})
}

type roundTripperFunc func(r *http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestGetFlag(t *testing.T) {
	t.Parallel()
	t.Run("default", func(t *testing.T) {
		t.Parallel()
		var cnt int64
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&cnt, 1)
			fmt.Fprintln(w, `{"enabled": true}`)
			assert.Equal(t, "/applications/app/environments/env/configurations/conf", r.URL.Path)
			assert.Equal(t, "myflag", r.URL.Query().Get("flag"))
			assert.Equal(t, []string(nil), r.Header.Values("Context"))
		}))
		defer ts.Close()

		client := NewAgentClient(WithBaseURLOption(ts.URL))
		got, err := client.GetFlag(context.Background(), "app", "env", "conf", "myflag", nil)
		assert.NoError(t, err)
		testutil.NoDiff(t, &GetFlagResult{Enabled: true}, got, nil)
		assert.Equal(t, int64(1), cnt)
	})

	t.Run("with evaluation context", func(t *testing.T) {
		t.Parallel()
		var cnt int64
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&cnt, 1)
			fmt.Fprintln(w, `{"enabled": true}`)
			assert.Equal(t, "/applications/app/environments/env/configurations/conf", r.URL.Path)
			assert.Equal(t, "myflag", r.URL.Query().Get("flag"))
			testutil.NoDiff(t, []string{"attr1=1", "attr2=hoge", "attr3=true"}, r.Header.Values("Context"), []cmp.Option{cmpopts.SortSlices(func(i, j string) bool { return i < j })})
		}))
		defer ts.Close()

		client := NewAgentClient(WithBaseURLOption(ts.URL))
		got, err := client.GetFlag(context.Background(), "app", "env", "conf", "myflag", map[string]any{"attr1": 1, "attr2": "hoge", "attr3": true})
		assert.NoError(t, err)
		testutil.NoDiff(t, &GetFlagResult{Enabled: true}, got, nil)
		assert.Equal(t, int64(1), cnt)
	})

	t.Run("with go context", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "key", "value")
		var cnt int64
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&cnt, 1)
			fmt.Fprintln(w, `{"enabled": true}`)
			assert.Equal(t, "/applications/app/environments/env/configurations/conf", r.URL.Path)
			assert.Equal(t, "myflag", r.URL.Query().Get("flag"))
			assert.Equal(t, []string(nil), r.Header.Values("Context"))
		}))
		defer ts.Close()

		fn := func(roundTripper http.RoundTripper) roundTripperFunc {
			return func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "value", r.Context().Value("key"))
				return roundTripper.RoundTrip(r)
			}
		}

		httpClient := &http.Client{
			Transport: fn(http.DefaultTransport),
		}
		client := NewAgentClient(WithBaseURLOption(ts.URL), WithHTTPClientOption(httpClient))
		got, err := client.GetFlag(ctx, "app", "env", "conf", "myflag", nil)
		assert.NoError(t, err)
		testutil.NoDiff(t, &GetFlagResult{Enabled: true}, got, nil)
		assert.Equal(t, int64(1), cnt)
	})
}

func TestJsonToResult(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		in      string
		want    *GetFlagResult
		wantErr bool
	}{
		"set enabled true": {
			in: `{"enabled": true}`,
			want: &GetFlagResult{
				Enabled: true,
			},
		},
		"set enabled false": {
			in:   `{"enabled": false}`,
			want: &GetFlagResult{},
		},
		"set variant": {
			in: `{"enabled": true, "_variant": "some variant"}`,
			want: &GetFlagResult{
				Enabled: true,
				Variant: "some variant",
			},
		},
		"with an attribute": {
			in: `{"enabled": true, "attr1": "hoge"}`,
			want: &GetFlagResult{
				Enabled:    true,
				Attributes: map[string]any{"attr1": "hoge"},
			},
		},
		"with some attributes": {
			in: `{"enabled": true, "attr1": "hoge", "attr2": 2}`,
			want: &GetFlagResult{
				Enabled:    true,
				Attributes: map[string]any{"attr1": "hoge", "attr2": float64(2)},
			},
		},
		"invalid json": {
			in:      `{`,
			wantErr: true,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := jsonToResult([]byte(tt.in))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				testutil.NoDiff(t, tt.want, got, nil)
			}
		})
	}
}
