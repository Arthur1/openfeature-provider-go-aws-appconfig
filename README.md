# openfeature-provider-go-aws-appconfig

Go implementation of OpenFeature Provider for AWS AppConfig (Feature Flags).

## Requirements

Go 1.22 or higher is required.

We supports latest two major releases of Go.

## Usage

```go
appName := "app" // Application name of AppConfig
envName := "env" // Environment name of AppConfig
cfgName := "cfg" // Configuration profile name of AppConfig
openfeature.SetProvider(appconfigprovider.New(appName, envName, cfgName))
client := openfeature.NewClient("app")
evalCtx := openfeature.NewTargetlessEvaluationContext(
	map[string]any{"userId": "userA"},
)
result, err := client.BooleanValueDetails(ctx, "feature1", false, evalCtx)
```

For more specific usage, please see an [example](./_example/).

## Specifications

Currently, we ONLY support getting flags **via [AppConfig Agent](https://docs.aws.amazon.com/en_us/appconfig/latest/userguide/appconfig-agent.html)**. If you wish to obtain flags via AWS SDK, please contribute to this project.

This provider ONLY supports getting **[Boolean values](https://openfeature.dev/specification/types#boolean)**, unless AppConfig Feature Flags supports non-boolean flag values.

The correspondence between data of AppConfig Feature Flags and boolean value details evaluated by the OpenFeature provider is shown in the following table:

|Data of AppConfig Feature Flags|Evaluation Details of OpenFeature|
|:--|:--|
|feature flag key|flag key|
|enabled value|value (boolean)|
|attributes|flag metadata|
|variant (in multi variant flag)|variant|
|caller context|evaluation context|

## License

MIT License

## Contact

Please contact me in GitHub issues or [`@Arthur1__` on X](https://x.com/arthur1__).
