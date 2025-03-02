# gopatchfields

[![Go Reference](https://pkg.go.dev/badge/github.com/mjruether/go-patch-fields.svg)](https://pkg.go.dev/github.com/mjruether/go-patch-fields)
[![Go Report Card](https://goreportcard.com/badge/github.com/mjruether/go-patch-fields)](https://goreportcard.com/report/github.com/mjruether/go-patch-fields)

## Overview

gopatchfields is a Go library that provides a type-safe and structured approach to HTTP PATCH operations. It solves the challenge of inconsistent PATCH endpoint implementations by offering a robust way to selectively update struct fields using wrapper types.

## Features

- 🔒 Type-safe field patching
- 🧩 Supports common Go types (bool, int, string, time.Time, etc.)
- 🚀 Automatic type handler registration
- 🛠 Configurable handling of unmatched fields
- 📝 Detailed warning and error reporting

## Installation

To install the package, use the following command:

```bash
go get github.com/mjruether/go-patch-fields@latest
```

### Import Specific Packages

```go
import (
    "github.com/mjruether/go-patch-fields/gopatchfields"
    "github.com/mjruether/go-patch-fields/gopatchservices"
)
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"

    "github.com/mjruether/go-patch-fields/gopatchfields"
    "github.com/mjruether/go-patch-fields/gopatchservices"
)

// Define your entity
type User struct {
    ID        int
    Name      string
    CreatedAt time.Time
    Active    bool
}

// Define your patch model using wrapper types
type UserPatch struct {
    Name      *gopatchfields.PatchFieldString
    CreatedAt *gopatchfields.PatchFieldTime
    Active    *gopatchfields.PatchFieldBool
}

func main() {
    // Create service (handlers are auto-registered)
    service := gopatchservices.NewPatchFieldService()

    // Create an example user
    user := &User{
        ID:        1,
        Name:      "Original Name",
        CreatedAt: time.Now(),
        Active:    true,
    }

    // Create a patch with some changes
    patch := &UserPatch{
        Name:   &gopatchfields.PatchFieldString{Value: "Updated Name"},
        Active: &gopatchfields.PatchFieldBool{Value: false},
    }

    // Apply the patch
    result := service.SetValues(user, patch)

    // Check for any warnings during patching
    if len(result.Warnings) > 0 {
        for field, warning := range result.Warnings {
            fmt.Printf("Warning for %s: %s\n", field, warning)
        }
    }

    // The user is now updated
    updatedUser := result.Entity.(*User)
    fmt.Printf("Updated User: %+v\n", updatedUser)
}
```

## Key Concepts

### Patch Fields

The library uses wrapper types (`PatchFieldString`, `PatchFieldInt`, etc.) to provide type-safe patching:
- `nil` wrapper means no update for that field
- Non-nil wrapper indicates a value to update

### Type Handlers

- Automatic registration for common types
- Extensible type conversion system
- Custom type handlers can be added via `service.GetTypeRegistry()`

## Configuration

### Ignoring Unmatched Fields

```go
service := gopatchservices.NewPatchFieldService()
service.SetIgnoreUnmatchedFields(true)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

See the LICENSE file for details.
