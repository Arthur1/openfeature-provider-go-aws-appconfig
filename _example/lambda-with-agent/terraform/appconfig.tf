resource "aws_appconfig_application" "this" {
  name = local.application_name
}

resource "aws_appconfig_configuration_profile" "this" {
  application_id = aws_appconfig_application.this.id
  location_uri   = "hosted"
  name           = local.configuration_name
  type           = "AWS.AppConfig.FeatureFlags"
}

resource "aws_appconfig_environment" "this" {
  name           = local.environment_name
  application_id = aws_appconfig_application.this.id
}

resource "aws_appconfig_hosted_configuration_version" "this" {
  application_id           = aws_appconfig_application.this.id
  configuration_profile_id = aws_appconfig_configuration_profile.this.configuration_profile_id
  content_type             = "application/json"
  content = jsonencode({
    version = "1"
    flags = {
      feature1 = {
        name       = "feature 1"
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
      }
      feature2 = {
        name       = "feature 2"
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
        attributes = {
          max_items = {
            constraints = {
              type     = "number"
              required = true
            }
          }
        }
      }
      feature3 = {
        name       = "feature 3"
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
      }
    }
    values = {
      feature1 = {
        enabled    = false
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
      }
      feature2 = {
        enabled    = true
        max_items  = 200
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
      }
      feature3 = {
        _variants = [
          {
            name    = "users"
            rule    = <<-EOT
            (or
              (eq $userId "userB")
              (eq $userId "userC")
            )
            EOT
            enabled = true
          },

          {
            name    = "default"
            enabled = false
          },
        ]
        _createdAt = "2024-08-28T10:00:00.000Z"
        _updatedAt = "2024-08-28T10:00:00.000Z"
      }
    }
  })
}

resource "aws_appconfig_deployment_strategy" "immediately" {
  name                           = "demo-immediately"
  deployment_duration_in_minutes = 0
  final_bake_time_in_minutes     = 0
  growth_factor                  = 100
  growth_type                    = "LINEAR"
  replicate_to                   = "NONE"
}

resource "aws_appconfig_deployment" "this" {
  application_id           = aws_appconfig_application.this.id
  configuration_profile_id = aws_appconfig_configuration_profile.this.configuration_profile_id
  configuration_version    = aws_appconfig_hosted_configuration_version.this.version_number
  deployment_strategy_id   = aws_appconfig_deployment_strategy.immediately.id
  environment_id           = aws_appconfig_environment.this.environment_id
}
