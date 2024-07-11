# Workflows

## add-labels-standardized.yaml

- [Add Labels Standardized GitHub action]
  - uses [senzing-factory/build-resources/.../add-labels-to-issue.yaml]

## add-to-project-garage-dependabot.yaml

- [Add to Project Garage Dependabot GitHub action]
  - uses [senzing-factory/build-resources/.../add-to-project-dependabot.yaml]

## add-to-project-garage.yaml

- [Add to Project Garage GitHub action]
  - uses [senzing-factory/build-resources/.../add-to-project.yaml]

## dependabot-approve-and-merge.yaml

- [Dependabot Approve and Merge GitHub action]
  - uses [senzing-factory/build-resources/.../dependabot-approve-and-merge.yaml]

## docker-build-container.yaml

- [Docker Build Container GitHub action]
  - uses [senzing-factory/github-action-docker-buildx-build]

## docker-push-containers-to-dockerhub.yaml

- [Docker Push Containers to DockerHub GitHub action]
  - uses [senzing-factory/github-action-docker-buildx-build]

## golangci-lint.yaml

- [Golangci Lint GitHub action]
  - uses:
    - [actions/checkout]
    - [senzing-factory/github-action-install-senzing-api]
    - [actions/setup-go]
    - [golangci/golangci-lint-action]

## go-proxy-pull.yaml

- [Go Proxy Pull GitHub action]
  - uses [andrewslotin/go-proxy-pull-action]

## go-test-darwin.yaml

- [Go Test Darwin GitHub action]
  - uses:
    - [actions/checkout]
    - [actions/setup-go]
    - [gotesttools/gotestfmt-action]
    - [senzing-factory/github-action-install-senzing-api]
    - [actions/upload-artifact]
    - [senzing-factory/build-resources/.../go-coverage.yaml]

## go-test-linux.yaml

- [Go Test Linux GitHub action]
  - uses:
    - [actions/checkout]
    - [actions/setup-go]
    - [gotesttools/gotestfmt-action]
    - [senzing-factory/github-action-install-senzing-api]
    - [actions/upload-artifact]
    - [senzing-factory/build-resources/.../go-coverage.yaml]

## go-test-windows.yaml

- [Go Test Windows GitHub action]
  - uses:
    - [actions/checkout]
    - [actions/setup-go]
    - [gotesttools/gotestfmt-action]
    - [senzing-factory/github-action-install-senzing-api]
    - [actions/upload-artifact]
    - [senzing-factory/build-resources/.../go-coverage.yaml]

## lint-workflows.yaml

- [Lint Workflows GitHub action]
  - uses [senzing-factory/build-resources/.../lint-workflows.yaml]

## make-go-github-file.yaml

- [Make Go GitHub File GitHub action]
  - uses [senzing-factory/build-resources/.../make-go-github-file.yaml]

## make-go-tag.yaml

- [Make Go Tag GitHub action]
  - uses:
    - [actions/checkout]
    - [senzing-factory/github-action-make-go-tag]

## move-pr-to-done-dependabot.yaml

- [Move PR to Done Dependabot GitHub action]
  - uses [senzing-factory/build-resources/.../move-pr-to-done-dependabot.yaml]

[actions/checkout]: https://github.com/actions/checkout
[actions/setup-go]: https://github.com/actions/setup-go
[actions/upload-artifact]: https://github.com/actions/upload-artifact
[Add Labels Standardized GitHub action]: add-labels-standardized.yaml
[Add to Project Garage Dependabot GitHub action]: add-to-project-garage-dependabot.yaml
[Add to Project Garage GitHub action]: add-to-project-garage.yaml
[andrewslotin/go-proxy-pull-action]: https://github.com/andrewslotin/go-proxy-pull-action
[Dependabot Approve and Merge GitHub action]: dependabot-approve-and-merge.yaml
[Docker Build Container GitHub action]: docker-build-container.yaml
[Docker Push Containers to DockerHub GitHub action]: docker-push-containers-to-dockerhub.yaml
[Go Proxy Pull GitHub action]: go-proxy-pull.yaml
[Go Test Darwin GitHub action]: go-test-darwin.yaml
[Go Test Linux GitHub action]: go-test-linux.yaml
[Go Test Windows GitHub action]: go-test-windows.yaml
[Golangci Lint GitHub action]: golangci-lint.yaml
[golangci/golangci-lint-action]: https://github.com/golangci/golangci-lint-action
[gotesttools/gotestfmt-action]: https://github.com/gotesttools/gotestfmt-action
[Lint Workflows GitHub action]: lint-workflows.yaml
[Make Go GitHub File GitHub action]: make-go-github-file.yaml
[Make Go Tag GitHub action]: make-go-tag.yaml
[Move PR to Done Dependabot GitHub action]: move-pr-to-done-dependabot.yaml
[senzing-factory/build-resources/.../add-labels-to-issue.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/add-labels-to-issue.yaml
[senzing-factory/build-resources/.../add-to-project-dependabot.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/add-to-project-dependabot.yaml
[senzing-factory/build-resources/.../add-to-project.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/add-to-project.yaml
[senzing-factory/build-resources/.../dependabot-approve-and-merge.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/dependabot-approve-and-merge.yaml
[senzing-factory/build-resources/.../go-coverage.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/go-coverage.yaml
[senzing-factory/build-resources/.../lint-workflows.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/lint-workflows.yaml
[senzing-factory/build-resources/.../make-go-github-file.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/make-go-github-file.yaml
[senzing-factory/build-resources/.../move-pr-to-done-dependabot.yaml]: https://github.com/senzing-factory/build-resources/blob/main/.github/workflows/move-pr-to-done-dependabot.yaml
[senzing-factory/github-action-docker-buildx-build]: https://github.com/senzing-factory/github-action-docker-buildx-build
[senzing-factory/github-action-install-senzing-api]: https://github.com/senzing-factory/github-action-install-senzing-api
[senzing-factory/github-action-make-go-tag]: https://github.com/senzing-factory/github-action-make-go-tag
