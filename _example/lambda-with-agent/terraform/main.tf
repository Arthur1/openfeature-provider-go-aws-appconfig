terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

locals {
  application_name   = "demo-app"
  environment_name   = "demo-env"
  configuration_name = "demo-conf"
}
