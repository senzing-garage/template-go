# template-go

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: template-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The template-go repository serves as a starting point for new repositories hosting Go code.
It also shows best practices that can be retro-fitted into existing repositories hosting Go code.

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

[![golangci-lint.yaml Badge]][golangci-lint.yaml]

## Overview

Aspects of the template-go repository:

1. **Badges:** Example badges can be seen above.
1. **Makefile:** Simplifies development lifecycle commands.
1. **Sample code:** `main.go` and `examplepackage` code examples.
1. **Sample test cases:** `*_test.go` files showing how to write and document test cases.
1. **Sample documentation:** Documentation style conducive to the [Go Package library].
1. **Dockerfile:** Containerizing the Go program.
1. **RPM/DEB builds:** Using `package.Dockerfile` to build `RPM` and `DEB` files for installation.
1. **.github/workflows:** GitActions tailored to Go programming.
1. **.github/dependabot.yml** Specifications for keeping Go dependencies up-to-date.

## Use

See [main.go] for an example of use.

## References

1. [SDK documentation]
1. [Development]
1. [Errors]
1. [Examples]
1. [Package reference]
1. Related artifacts:
    1. [DockerHub]

[SDK documentation]: https://pkg.go.dev/github.com/senzing-garage/template-go
[Development]: docs/development.md
[DockerHub]: https://hub.docker.com/r/senzing/template-go
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[Go Package library]: https://pkg.go.dev
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/template-go.svg
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/template-go
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/template-go
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-darwin.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-darwin.yaml
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-linux.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-linux.yaml
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-windows.yaml/badge.svg
[go-test-windows.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/template-go/actions/workflows/golangci-lint.yaml/badge.svg
[golangci-lint.yaml]: https://github.com/senzing-garage/template-go/actions/workflows/golangci-lint.yaml
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[License]: https://github.com/senzing-garage/template-go/blob/main/LICENSE
[main.go]: main.go
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/template-go
[Senzing Garage]: https://github.com/senzing-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing]: https://senzing.com/
