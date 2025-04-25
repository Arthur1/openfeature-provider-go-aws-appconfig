module github.com/Arthur1/openfeature-provider-go-aws-appconfig/_example/lambda-with-agent

go 1.23.0

toolchain go1.24.2

require (
	github.com/Arthur1/openfeature-provider-go-aws-appconfig v0.0.0
	github.com/aws/aws-lambda-go v1.47.0
	github.com/caarlos0/env/v11 v11.2.2
	github.com/open-feature/go-sdk v1.14.1
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
)

replace github.com/Arthur1/openfeature-provider-go-aws-appconfig v0.0.0 => ../../
