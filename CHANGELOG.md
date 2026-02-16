# Changelog

## [v0.4.0](https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/compare/v0.3.0...v0.4.0) - 2026-02-16
- chore: bump minimum supported Go version to 1.23 by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/32
- test: add test with running AppConfig Agent container by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/31
- test: use goleak to detect goroutine leaks by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/34
- ci: remove test cache by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/35
- docs: update lambda with agent example code by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/36
- test: remove the log output assertion for agent because it's flaky by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/51
- chore: bump minimum version of go to 1.25 by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/52
- chore: upgrade some GitHub Actions by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/53
- fix(deps): update module github.com/open-feature/go-sdk to v1.17.1 by @renovate[bot] in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/42
- fix(deps): update module github.com/testcontainers/testcontainers-go to v0.40.0 by @renovate[bot] in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/43
- chore: use tparse via go tools by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/54
- docs: update docs & example code fort Go 1.25 by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/55
- refactor: use t.Context() by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/56

## [v0.3.0](https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/compare/v0.2.1...v0.3.0) - 2025-04-25
- feat: Create request with go context by @roothybrid7 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/24
- test: fix SA1029: should not use built-in type string as key for value by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/27
- chore(deps): update golangci/golangci-lint-action action to v7 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/23
- ci: pin GitHub Actions Digets by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/28
- fix(deps): update module github.com/open-feature/go-sdk to v1.14.1 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/21
- chore(deps): pin dependencies by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/29
- fix(deps): update module github.com/google/go-cmp to v0.7.0 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/22
- chore(deps): update actions/create-github-app-token action to v2 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/25
- test: use compare functions with SortSlices in go-cmp by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/30

## [v0.2.1](https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/compare/v0.2.0...v0.2.1) - 2025-01-21
- docs: add context to correspondence table by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/12
- chore: call t.Helper() in test helper functions by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/18
- Update module github.com/open-feature/go-sdk to v1.14.0 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/16
- Update module github.com/stretchr/testify to v1.10.0 by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/17
- bugfix: avoid resource leak by closing response body in http requests by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/19
- chore: fix ignore paths on renovate by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/20

## [v0.2.0](https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/compare/v0.1.0...v0.2.0) - 2024-09-01
- BREAKING: refactor interfaces by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/8
- ignore sub module update by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/9
- ignore sub module update by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/11

## [v0.1.0](https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/commits/v0.1.0) - 2024-08-29
- Configure Renovate by @renovate in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/1
- add Example with Lambda AppConfig Agent by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/4
- add README.md by @Arthur1 in https://github.com/Arthur1/openfeature-provider-go-aws-appconfig/pull/5
