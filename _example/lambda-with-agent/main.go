package main

import (
	"context"
	"log"

	"github.com/Arthur1/openfeature-provider-go-aws-appconfig/appconfigprovider"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v11"
	"github.com/open-feature/go-sdk/openfeature"
)

type config struct {
	AppConfigApp string `env:"APPCONFIG_APPLICATION"`
	AppConfigEnv string `env:"APPCONFIG_ENVIRONMENT"`
	AppConfigCfg string `env:"APPCONFIG_CONFIGURATION"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	openfeature.SetProviderAndWait(appconfigprovider.New(cfg.AppConfigApp, cfg.AppConfigEnv, cfg.AppConfigCfg))
	client := openfeature.NewClient("app")
	evalCtxA := openfeature.NewTargetlessEvaluationContext(
		map[string]any{"userId": "userA"},
	)
	evalCtxB := openfeature.NewTargetlessEvaluationContext(
		map[string]any{"userId": "userB"},
	)
	result1, err := client.BooleanValueDetails(ctx, "feature1", false, evalCtxA)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result1 value: expected=false, actual=%t\n", result1.Value)
	log.Printf("result1 variant: expected=, actual=%s\n", result1.Variant)
	log.Printf("result1 flag metadata: expected=map[], actual=%+v\n", result1.FlagMetadata)

	result2, err := client.BooleanValueDetails(ctx, "feature2", false, evalCtxA)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result2 value: expected=true, actual=%t\n", result2.Value)
	log.Printf("result2 variant: expected=, actual=%s\n", result2.Variant)
	log.Printf("result2 flag metadata: expected=map[max_items:200], actual=%+v\n", result2.FlagMetadata)

	result3a, err := client.BooleanValueDetails(ctx, "feature3", false, evalCtxA)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result3a value: expected=false, actual=%t\n", result3a.Value)
	log.Printf("result3a variant: expected=default, actual=%s\n", result3a.Variant)
	log.Printf("result3a flag metadata: expected=map[], actual=%+v\n", result3a.FlagMetadata)

	result3b, err := client.BooleanValueDetails(ctx, "feature3", false, evalCtxB)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("result3b value: expected=true, actual=%t\n", result3b.Value)
	log.Printf("result3b variant: expected=users, actual=%s\n", result3b.Variant)
	log.Printf("result3b flag metadata: expected=map[], actual=%+v\n", result3b.FlagMetadata)
}
