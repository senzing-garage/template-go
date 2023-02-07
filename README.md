# template-go

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/template-go.svg)](https://pkg.go.dev/github.com/senzing/template-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/template-go)](https://goreportcard.com/report/github.com/senzing/template-go)
[![go-test.yaml](https://github.com/Senzing/template-go/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/template-go/actions/workflows/go-test.yaml)

## Synopsis

The template-go repository serves as a starting point for new repositories hosting Go code.
It also shows best practices that can be retro-fitted into existing repositories hosting Go code.

## Overview

Aspects of the template-go repository:

1. **Badges:** Example badges can be seen above.
1. **Makefile:** Simplifies development lifecycle commands.
1. **Sample code:** `main.go` and `examplepackage` code examples.
1. **Sample test cases:** `*_test.go` files showing how to write and document test cases.
1. **Sample documentation:** Documentation style conducive to the [Go Package library](https://pkg.go.dev).
1. **Dockerfile:** Containerizing the Go program.
1. **RPM/DEB builds:** Using `package.Dockerfile` to build `RPM` and `DEB` files for installation.
1. **.github/workflows:** GitActions tailored to Go programming.
1. **.github/dependabot.yml** Specifications for keeping Go dependencies up-to-date.

## Development

### Build

1. Identify repository.
   Example:

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=template-go
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.
1. Build the binaries.
   Example:

     ```console
     cd ${GIT_REPOSITORY_DIR}
     make build

     ```

1. The binaries will be found in ${GIT_REPOSITORY_DIR}/target.
   Example:

    ```console
    tree ${GIT_REPOSITORY_DIR}/target

    ```

1. Run the binary.
   Example:

    ```console
    ${GIT_REPOSITORY_DIR}/target/linux/template-go

    ```

1. Clean up.
   Example:

     ```console
     cd ${GIT_REPOSITORY_DIR}
     make clean

     ```

### Test

1. Run Go tests.
   Example:

     ```console
     cd ${GIT_REPOSITORY_DIR}
     make test

     ```

### Documentation

1. Start `godoc` documentation server.
   Example:

    ```console
     cd ${GIT_REPOSITORY_DIR}
     godoc

    ```

1. Visit [localhost:6060](http://localhost:6060)
1. Senzing documentation will be in the "Third party" section.
   `github.com` > `senzing` > `template-go`

1. When published, the reference can be found by clicking on the following badge:

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-observing.svg)](https://pkg.go.dev/github.com/senzing/go-observing)

### Docker

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make docker-build

    ```

1. Run docker container.
   Example:

    ```console
    docker run \
      --rm \
      senzing/template-go

    ```

### Package

#### Package RPM and DEB files

1. Use make target to run a docker images that builds RPM and DEB files.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make package

    ```

1. The results will be in the `${GIT_REPOSITORY_DIR}/target` directory.
   Example:

    ```console
    tree ${GIT_REPOSITORY_DIR}/target

    ```

#### Test DEB package on Ubuntu

1. Determine if `template-go` is installed.
   Example:

    ```console
    apt list --installed | grep template-go

    ```

1. :pencil2: Install `template-go`.
   Example:

    ```console
    cd ${GIT_REPOSITORY_DIR}/target
    sudo apt install ./template-go-0.0.0.deb

    ```

1. Run command.
   Example:

    ```console
    template-go

    ```

1. Remove `template-go` from system.
   Example:

    ```console
    sudo apt-get remove template-go

    ```
