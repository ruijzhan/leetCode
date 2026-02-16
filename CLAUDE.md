# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

LeetCode practice repository in Go, containing solutions for LeetCode problems and "JianZhi Offer" (剑指 Offer - Chinese coding interview problems).

## Commands

### Testing
```bash
# Run all tests
go test -v ./...

# Run tests for a specific package (category)
go test -v ./problems/array
go test -v ./JianZhiOffer

# Run tests for a specific problem
go test -v ./problems/128
```

### Update Statistics
```bash
./gen_report.sh    # Regenerates README.md with problem counts by difficulty
```

## Code Organization

### Problem Locations

| Location | Description |
|----------|-------------|
| `problems/{category}/` | LeetCode problems grouped by algorithm category (array, dp, stack, tree, etc.) |
| `JianZhiOffer/` | 剑指 Offer problems |
| `pkg/` | Shared utility packages (stack, utils) |

### Problem File Patterns

**New format (problem-number directory):**
```
problems/128/
├── q.md        # Question description (Chinese)
├── a.go        # Solution implementation
└── a_test.go   # Test cases
```

**Legacy format (category directory):**
```
problems/array/E118_PascalTriangle.go
problems/array/E118_PascalTriangle_test.go
```

**JianZhi Offer format:**
```
JianZhiOffer/E06_从尾到头打印链表.go         # E=Easy, M=Medium, H=Hard
JianZhiOffer/E06_从尾到头打印链表_test.go
```

## Testing Conventions

Tests follow Go's table-driven testing pattern:

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name string
        input    InputType
        want     ExpectedType
    }{
        {name: "case_name", input: ..., want: ...},
        // ... more cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := functionToTest(tt.input); got != tt.want {
                t.Fatalf("functionToTest(%v) = %v, want %v", tt.input, got, tt.want)
            }
        })
    }
}
```

## Naming Conventions

- Functions use `camelCase`
- Test files: `*_test.go`
- Difficulty prefixes: `E` (Easy/简单), `M` (Medium/中等), `H` (Hard/困难)
- Package names match directory names (lowercase)

## Utilities

- `pkg/stack/` - Generic and int-specific stack implementations with `Interface` and `StackInt` types
- `pkg/utils/` - Utility functions like `GetFunctionName()` for runtime reflection
