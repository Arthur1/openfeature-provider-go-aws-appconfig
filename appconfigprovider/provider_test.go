package appconfigprovider

import (
	"context"
	"testing"

	"github.com/Arthur1/openfeature-provider-go-aws-appconfig/appconfig"
	"github.com/stretchr/testify/assert"
)

type dummyClient struct{}

var _ appconfig.Client = (*dummyClient)(nil)

func (c *dummyClient) GetFlag(_ context.Context, _, _, _, _ string, _ map[string]any) (*appconfig.GetFlagResult, error) {
	return &appconfig.GetFlagResult{}, nil
}

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("Default", func(t *testing.T) {
		t.Parallel()
		provider := New("app", "env", "conf")
		assert.Equal(t, "app", provider.application)
		assert.Equal(t, "env", provider.environment)
		assert.Equal(t, "conf", provider.configuration)
		assert.IsType(t, &appconfig.AgentClient{}, provider.client)
	})

	t.Run("WithClient", func(t *testing.T) {
		t.Parallel()
		client := &dummyClient{}
		provider := New("app", "env", "conf", WithClientOption(client))
		assert.Equal(t, "app", provider.application)
		assert.Equal(t, "env", provider.environment)
		assert.Equal(t, "conf", provider.configuration)
		assert.Equal(t, client, provider.client)
	})
}
