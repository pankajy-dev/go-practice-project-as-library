# go-practice-project-as-library

A Go library providing utility packages for string manipulation and integer operations.

## Installation

### Option 1: Using Version Tags (Recommended)

```bash
# Get latest version (easiest)
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@latest

# Or get specific version
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@v0.1.1
```

### Option 2: Using Local Path (For Development)

In your project's `go.mod`:

```go
require github.com/pankajy-dev/go-practice-project-as-library v0.1.1

replace github.com/pankajy-dev/go-practice-project-as-library => /Users/pyadav/all/workspace/codebase/pankaj/go/go-practice-project-as-library
```

Then run:
```bash
go mod tidy
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/pankajy-dev/go-practice-project-as-library/stringutilpankaj"
    "github.com/pankajy-dev/go-practice-project-as-library/intutil"
)

func main() {
    // String utilities
    text := "hello world"
    fmt.Println(stringutilpankaj.Capitalize(text))      // Hello world
    fmt.Println(stringutilpankaj.CapitalizeWords(text)) // Hello World
    fmt.Println(stringutilpankaj.ToUpper(text))         // HELLO WORLD
    fmt.Println(stringutilpankaj.ToLower("HELLO"))      // hello

    // Integer utilities
    fmt.Println(intutil.Random(1, 100))                 // Random number 1-100
    fmt.Println(intutil.IsPrime(17))                    // true
    fmt.Println(intutil.Factorial(5))                   // 120
    fmt.Println(intutil.GCD(48, 18))                    // 6
    fmt.Println(intutil.Abs(-42))                       // 42
}
```

## Available Packages

### stringutilpankaj
- `Capitalize(s string)` - Capitalize first character
- `CapitalizeWords(s string)` - Title case
- `ToUpper(s string)` - Convert to uppercase
- `ToLower(s string)` - Convert to lowercase

### intutil
- `Random(min, max int)` - Random number in range
- `RandomN(n int)` - Random number 0 to n
- `RandomSlice(length, min, max int)` - Generate random slice
- `Abs(n int)` - Absolute value
- `Max(a, b int)` - Maximum of two numbers
- `Min(a, b int)` - Minimum of two numbers
- `Clamp(value, min, max int)` - Restrict to range
- `IsEven(n int)` - Check if even
- `IsOdd(n int)` - Check if odd
- `IsPrime(n int)` - Check if prime
- `Factorial(n int)` - Calculate factorial
- `GCD(a, b int)` - Greatest common divisor
- `LCM(a, b int)` - Least common multiple

## Development Workflow

### Making Changes to This Library

1. **Make your changes** to the code
2. **Run tests** to ensure everything works:
   ```bash
   go test ./...
   ```

3. **Commit your changes**:
   ```bash
   git add .
   git commit -m "Add new feature or fix bug"
   ```

4. **Create a new version tag**:
   ```bash
   # Increment version appropriately
   git tag v0.1.2
   git push origin main
   git push origin v0.1.2
   ```

### Using Updated Library in Another Project

#### If using local path (Option 2):
Changes are immediately available. Just import and use.

#### If using GitHub tags (Option 1):
After pushing a new tag, update in your project:

```bash
# Get latest version (recommended)
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@latest

# Or get specific version
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@v0.1.2

# If cached, clear first
go clean -modcache
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@latest
```

Or temporarily for development, add to your shell config (`~/.zshrc`):

```bash
# Add this function
pankaj-go-get() {
    GOPROXY=direct GOSUMDB=off go get "$@"
}
```

Then use:
```bash
# Get latest version
pankaj-go-get github.com/pankajy-dev/go-practice-project-as-library@latest

# Or specific version
pankaj-go-get github.com/pankajy-dev/go-practice-project-as-library@v0.1.2
```

## Versioning

This project follows semantic versioning (semver):
- `v0.1.x` - Patch version (bug fixes)
- `v0.x.0` - Minor version (new features, backward compatible)
- `vx.0.0` - Major version (breaking changes)

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./stringutilpankaj
go test ./intutil

# Verbose output
go test -v ./...
```

## Building

```bash
# Build all packages
go build ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

## Common Issues

### "verifying module" 404 error
Use direct fetch bypassing proxy:
```bash
GOPROXY=direct GOSUMDB=off go get github.com/pankajy-dev/go-practice-project-as-library@vX.X.X
```

### Module path mismatch
Ensure `go.mod` contains:
```
module github.com/pankajy-dev/go-practice-project-as-library
```

### Changes not appearing
- If using local path: Changes are immediate
- If using GitHub: Must create and push a new tag

## License

This is a practice project for learning Go library development.
