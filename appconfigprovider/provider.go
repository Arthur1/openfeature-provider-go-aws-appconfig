package appconfigprovider

import (
	"context"

	"github.com/Arthur1/openfeature-provider-go-aws-appconfig/appconfig"
	"github.com/open-feature/go-sdk/openfeature"
)

type Provider struct {
	client appconfig.Client
}

var _ openfeature.FeatureProvider = (*Provider)(nil)

func NewProvider(client appconfig.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "AWSAppConfig",
	}
}

func (p *Provider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	res, err := p.client.GetFlag(ctx, flag, evalCtx)
	if err != nil {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				Reason: openfeature.ErrorReason,
				// TODO: Use ResolutionError appropriately depending on the type of error returned by client.
				ResolutionError: openfeature.NewGeneralResolutionError(err.Error()),
			},
		}
	}
	return openfeature.BoolResolutionDetail{
		Value: res.Enabled,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:       openfeature.UnknownReason,
			Variant:      res.Variant,
			FlagMetadata: res.Attributes,
		},
	}
}

func (p *Provider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	return openfeature.StringResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:          openfeature.ErrorReason,
			ResolutionError: openfeature.NewTypeMismatchResolutionError("StringEvaluation is not supported"),
		},
	}
}

func (p *Provider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:          openfeature.ErrorReason,
			ResolutionError: openfeature.NewTypeMismatchResolutionError("FloatEvaluation is not supported"),
		},
	}
}

func (p *Provider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:          openfeature.ErrorReason,
			ResolutionError: openfeature.NewTypeMismatchResolutionError("IntEvaluation is not supported"),
		},
	}
}

func (p *Provider) ObjectEvaluation(ctx context.Context, flag string, defaultValue any, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:          openfeature.ErrorReason,
			ResolutionError: openfeature.NewTypeMismatchResolutionError("ObjectEvaluation is not supported"),
		},
	}
}

func (p *Provider) Hooks() []openfeature.Hook {
	return nil
}
