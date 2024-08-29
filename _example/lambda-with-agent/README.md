# Example with Lambda AppConfig Agent

This is a demo of openfeature-provider-go-aws-appconfig along with AWS AppConfig Agent for AWS Lambda.

## Preparation

Resources such as AppConfig and a Lambda function containing the agent layer will be created.

```console
$ cd ./terraform
$ terraform init
$ terraform apply
```

## Expected Result

```
[appconfig agent] 2024/08/29 17:22:53 INFO AppConfig Lambda Extension 2.0.711
[appconfig agent] 2024/08/29 17:22:53 INFO serving on localhost:2772
EXTENSION	Name: AppConfigAgent	State: Ready	Events: [INVOKE, SHUTDOWN]
START RequestId: a2366c71-cf93-4866-9e74-9bfbffb8095b Version: $LATEST
[appconfig agent] 2024/08/29 17:22:54 INFO retrieved initial data for 'demo-app:demo-env:demo-conf' in 255ms
2024/08/29 17:22:54 result1 value: expected=false, actual=false
2024/08/29 17:22:54 result1 variant: expected=, actual=
2024/08/29 17:22:54 result1 flag metadata: expected=map[], actual=map[]
2024/08/29 17:22:54 result2 value: expected=true, actual=true
2024/08/29 17:22:54 result2 variant: expected=, actual=
2024/08/29 17:22:54 result2 flag metadata: expected=map[max_items:200], actual=map[max_items:200]
2024/08/29 17:22:54 result3a value: expected=false, actual=false
2024/08/29 17:22:54 result3a variant: expected=default, actual=default
2024/08/29 17:22:54 result3a flag metadata: expected=map[], actual=map[]
2024/08/29 17:22:54 result3b value: expected=true, actual=true
2024/08/29 17:22:54 result3b variant: expected=users, actual=users
2024/08/29 17:22:54 result3b flag metadata: expected=map[], actual=map[]
END RequestId: a2366c71-cf93-4866-9e74-9bfbffb8095b
REPORT RequestId: a2366c71-cf93-4866-9e74-9bfbffb8095b	Duration: 263.34 ms	Billed Duration: 467 ms	Memory Size: 128 MB	Max Memory Used: 38 MB	Init Duration: 203.47 ms	
```
