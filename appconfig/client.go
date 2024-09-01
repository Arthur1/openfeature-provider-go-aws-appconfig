package appconfig

import (
	"context"
)

type Client interface {
	GetFlag(ctx context.Context, application, environment, configuration, flagName string, evalCtx map[string]any) (*GetFlagResult, error)
}

type GetFlagResult struct {
	Enabled    bool
	Variant    string
	Attributes map[string]any
}
