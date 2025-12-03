# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go template repository for Senzing Garage projects. It demonstrates best practices for Go CLI applications including project structure, testing, linting, and packaging.

## Common Commands

### Build

```bash
make build              # Build for current platform (output in target/)
make build-all          # Build for all platforms (darwin/linux/windows, amd64/arm64)
```

### Test

```bash
make test               # Run tests with formatted output
go test ./...           # Run tests directly
go test -v ./... -run TestName  # Run a specific test
```

### Lint

```bash
make lint               # Run all linters (golangci-lint, govulncheck, cspell)
make golangci-lint      # Run golangci-lint only
make fix                # Auto-fix linting issues
```

### Coverage

```bash
make coverage           # Generate coverage report and open in browser
make check-coverage     # Check coverage thresholds (80% file/package/total)
```

### Other

```bash
make run                # Run the program directly via go run
make clean              # Clean build artifacts and caches
make dependencies       # Update Go dependencies
make dependencies-for-development  # Install development tools
```

## Architecture

### Project Structure

- `main.go` - Entry point, calls `cmd.Execute()`
- `cmd/` - Cobra CLI commands and configuration
  - `root.go` - Main command definition with context variables for CLI flags/env vars
  - `github.go` - Auto-generated version info from GitHub Actions
  - `context_*.go` - OS-specific context variables
- `examplepackage/` - Example package demonstrating interface patterns
  - `main.go` - Interface definition (`ExamplePackage`)
  - `examplepackage_basic.go` - Interface implementation (`BasicExamplePackage`)

### Key Patterns

**CLI Configuration**: Uses `go-cmdhelping` library with `option.ContextVariable` for unified handling of CLI flags, environment variables, and config files via Viper.

**Interface Pattern**: Define interfaces in `main.go` of a package, implement in separate `*_basic.go` files.

**Error Handling**: Use `go-helpers/wraperror` for error wrapping.

### Makefile System

The Makefile uses OS detection with platform-specific includes:

- `makefiles/osdetect.mk` - Detects OS type and architecture
- `makefiles/{darwin,linux,windows}.mk` - OS-specific target implementations

## Linting Configuration

Golangci-lint config: `.github/linters/.golangci.yaml`

- Line length: 120 characters
- Coverage thresholds: 80% (configurable in `.github/coverage/testcoverage.yaml`)
- Uses extensive linter set including exhaustruct, wrapcheck, err113
