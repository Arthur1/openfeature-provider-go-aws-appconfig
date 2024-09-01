package appconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AgentClient struct {
	httpClient *http.Client
	baseURL    string
}

var _ Client = (*AgentClient)(nil)

type agentClientOptions struct {
	httpClient *http.Client
	baseURL    string
}

type AgentClientOption interface {
	apply(opts *agentClientOptions)
}

type httpClientOption struct {
	httpClient *http.Client
}

func (o httpClientOption) apply(opts *agentClientOptions) {
	opts.httpClient = o.httpClient
}

func WithHTTPClientOption(httpClient *http.Client) httpClientOption {
	return httpClientOption{httpClient: httpClient}
}

type baseURLOption string

func (o baseURLOption) apply(opts *agentClientOptions) {
	opts.baseURL = string(o)
}

func WithBaseURLOption(baseURL string) baseURLOption {
	return baseURLOption(baseURL)
}

func NewAgentClient(opts ...AgentClientOption) *AgentClient {
	options := &agentClientOptions{
		httpClient: http.DefaultClient,
		baseURL:    "http://localhost:2772",
	}
	for _, o := range opts {
		o.apply(options)
	}

	return &AgentClient{
		httpClient: options.httpClient,
		baseURL:    options.baseURL,
	}
}

func (c *AgentClient) GetFlag(_ context.Context, application, environment, configuration, flagName string, evalCtx map[string]any) (*GetFlagResult, error) {
	endpointURL, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	endpointURL = endpointURL.JoinPath("applications", application, "environments", environment, "configurations", configuration)
	query := endpointURL.Query()
	query.Set("flag", flagName)
	endpointURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, endpointURL.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range evalCtx {
		vs, err := evalCtxValueToString(v)
		if err != nil {
			continue
		}
		req.Header.Add("Context", fmt.Sprintf("%s=%s", k, vs))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server has returned invalid status code: %d", res.StatusCode)
	}
	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result, err := jsonToResult(resb)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func evalCtxValueToString(v any) (string, error) {
	switch vt := v.(type) {
	case string:
		return vt, nil
	case bool:
		return strconv.FormatBool(vt), nil
	case int64:
		return strconv.FormatInt(vt, 10), nil
	case float64:
		return strconv.FormatFloat(vt, 'E', -1, 64), nil
	case time.Time:
		return vt.String(), nil
	default:
		j, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(j), nil
	}
}

func jsonToResult(jsonb []byte) (*GetFlagResult, error) {
	var kvs map[string]any
	if err := json.Unmarshal(jsonb, &kvs); err != nil {
		return nil, err
	}

	result := &GetFlagResult{}
	for k, v := range kvs {
		switch k {
		case "enabled":
			enabled, ok := v.(bool)
			if !ok {
				return nil, fmt.Errorf("enabled must be bool")
			}
			result.Enabled = enabled
		case "_variant":
			variant, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("_variant must be string")
			}
			result.Variant = variant
		default:
			if result.Attributes == nil {
				result.Attributes = map[string]any{}
			}
			result.Attributes[k] = v
		}
	}
	return result, nil
}
