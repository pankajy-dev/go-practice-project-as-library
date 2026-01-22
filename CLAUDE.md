# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a pure Go library project with no main executable. It provides reusable packages meant to be imported by other Go projects. The main example is a string utilities package (`stringutilpankaj`).

## Architecture

The project follows a standard Go library pattern:

- **`stringutilpankaj/`**: Package providing string manipulation utilities
  - Functions for capitalizing, uppercasing, and lowercasing strings
- **`intutil/`**: Package providing integer utilities
  - Random number generation (single, range, slices)
  - Math operations (abs, max, min, clamp)
  - Number properties (even, odd, prime)
  - Advanced math (factorial, GCD, LCM)

This is a **library-only project** with no main.go. To use this library, other projects would import it via Go modules.

## Development Commands

### Building
```bash
go build ./...              # Build all packages
```

### Testing
```bash
go test ./...               # Run all tests in the project
go test ./stringutilpankaj  # Test string utilities package
go test ./intutil           # Test integer utilities package
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

Import packages from this library:

```go
import (
    "github.com/pankajy-dev/go-practice-project-as-library/stringutilpankaj"
    "github.com/pankajy-dev/go-practice-project-as-library/intutil"
)

// String utilities
capitalized := stringutilpankaj.Capitalize("hello")           // "Hello"
titleCase := stringutilpankaj.CapitalizeWords("hello world")  // "Hello World"
upper := stringutilpankaj.ToUpper("hello")                    // "HELLO"

// Integer utilities
random := intutil.Random(1, 100)                              // Random number between 1-100
isPrime := intutil.IsPrime(17)                                // true
factorial := intutil.Factorial(5)                             // 120
gcd := intutil.GCD(48, 18)                                    // 6
```

The `intutil` package has a short name for convenient use without an alias.

## Adding New Utilities

When adding new utility packages:
1. Create a new directory under the project root (e.g., `mathutil/`, `fileutil/`)
2. Package name should match the directory name
3. Export functions by capitalizing the first letter
4. Add corresponding tests in `*_test.go` files
5. Update this documentation if the package introduces new architectural patterns
