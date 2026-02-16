module github.com/Arthur1/openfeature-provider-go-aws-appconfig/_example/lambda-with-agent

go 1.25.0

require (
	github.com/Arthur1/openfeature-provider-go-aws-appconfig v0.0.0
	github.com/aws/aws-lambda-go v1.47.0
	github.com/caarlos0/env/v11 v11.2.2
	github.com/open-feature/go-sdk v1.17.1
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	go.uber.org/mock v0.6.0 // indirect
)

replace github.com/Arthur1/openfeature-provider-go-aws-appconfig v0.0.0 => ../../
