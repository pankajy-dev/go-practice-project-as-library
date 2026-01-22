# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a pure Go library project with no main executable. It provides reusable packages meant to be imported by other Go projects. The main example is a string utilities package (`stringutilpankaj`).

## Architecture

The project follows a standard Go library pattern:

- **`stringutilpankaj/`**: Package providing string manipulation utilities
  - `capitalize.go`: Functions for capitalizing strings (single character or all words)

This is a **library-only project** with no main.go. To use this library, other projects would import it via Go modules.

## Development Commands

### Building
```bash
go build ./...              # Build all packages
```

### Testing
```bash
go test ./...               # Run all tests in the project
go test ./stringutilpankaj  # Test specific package
go test -v ./...            # Verbose test output
```

### Module Management
```bash
go mod tidy                 # Clean up dependencies
go mod verify               # Verify dependencies
```

### Formatting and Linting
```bash
go fmt ./...                # Format all Go files
go vet ./...                # Run Go vet for common issues
```

## Package Usage Pattern

To use the `stringutilpankaj` package in other parts of the code:

```go
import "go-practice-project-as-library/stringutilpankaj"

capitalized := stringutilpankaj.Capitalize("hello")        // "Hello"
titleCase := stringutilpankaj.CapitalizeWords("hello world") // "Hello World"
```

Note: The module name is `go-practice-project-as-library` (from go.mod), so imports use this as the base path.

## Adding New Utilities

When adding new utility packages:
1. Create a new directory under the project root (e.g., `mathutil/`, `fileutil/`)
2. Package name should match the directory name
3. Export functions by capitalizing the first letter
4. Add corresponding tests in `*_test.go` files
5. Update this documentation if the package introduces new architectural patterns
