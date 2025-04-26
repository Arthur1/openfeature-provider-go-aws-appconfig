resource "aws_lambda_function" "this" {
  function_name    = "openfeature-appconfig-demo"
  package_type     = "Zip"
  role             = aws_iam_role.lambda.arn
  architectures    = ["arm64"]
  filename         = "${path.module}/../archive.zip"
  source_code_hash = filebase64sha256("${path.module}/../archive.zip")
  handler          = "main"
  memory_size      = 128
  timeout          = 7
  runtime          = "provided.al2023"
  // https://docs.aws.amazon.com/en_us/appconfig/latest/userguide/appconfig-integration-lambda-extensions-versions.html#appconfig-integration-lambda-extensions-enabling-ARM64
  layers = ["arn:aws:lambda:ap-northeast-1:980059726660:layer:AWS-AppConfig-Extension-Arm64:79"]

  environment {
    variables = {
      APPCONFIG_APPLICATION   = local.application_name
      APPCONFIG_ENVIRONMENT   = local.environment_name
      APPCONFIG_CONFIGURATION = local.configuration_name
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = "openfeature-appconfig-demo-lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role_lambda.json
}

data "aws_iam_policy_document" "assume_role_lambda" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role_policy" "lambda" {
  role   = aws_iam_role.lambda.id
  name   = "appconfig"
  policy = data.aws_iam_policy_document.appconfig.json
}

data "aws_iam_policy_document" "appconfig" {
  statement {
    effect = "Allow"
    actions = [
      "appconfig:StartConfigurationSession",
      "appconfig:GetLatestConfiguration",
    ]
    resources = [
      aws_appconfig_application.this.arn,
      "${aws_appconfig_application.this.arn}/*",
    ]
  }
}

resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda.id
  policy_arn = data.aws_iam_policy.lambda_basic_execution.arn
}

data "aws_iam_policy" "lambda_basic_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
