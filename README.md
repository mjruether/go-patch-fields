# PatchPropertiesGo

A Go port of PatchProperties that provides a flexible property patching system. This library allows you to selectively update struct fields using patch models that wrap values in type-safe containers.

## Features

- Type-safe property patching
- Customizable type conversion system
- Support for basic Go types (bool, int, string, etc.)
- Support for time.Time and custom types
- Nullable support through pointers
- Detailed warnings and error handling

## Example Usage

```go
// Define your entity
type User struct {
    ID        int
    Name      string
    CreatedAt time.Time
    Active    bool
}

// Define your patch model
type UserPatch struct {
    Name      *PatchPropertyString
    CreatedAt *PatchPropertyTime
    Active    *PatchPropertyBool
}

// Create and configure the service
service := services.NewPatchPropertyService()

// Register type handlers
registry := service.GetTypeRegistry()
registry.RegisterHandler(reflect.TypeOf(""), &handlers.StringHandler{})
registry.RegisterHandler(reflect.TypeOf(time.Time{}), &handlers.TimeHandler{})
registry.RegisterHandler(reflect.TypeOf(true), &handlers.BoolHandler{})

// Create a patch
patch := &UserPatch{
    Name:   &PatchPropertyString{Value: "Updated Name"},
    Active: &PatchPropertyBool{Value: false},
}

// Apply the patch
result := service.SetValues(user, patch)

// Check for warnings
if len(result.Warnings) > 0 {
    for field, warning := range result.Warnings {
        fmt.Printf("%s: %s\n", field, warning)
    }
}
```

## Differences from C# Version

1. Uses Go interfaces and reflection instead of C# generics
2. Implements nullability through pointers rather than C#'s nullable types
3. Provides custom type conversion through registered handlers
4. More explicit error handling through return values
5. Uses Go naming conventions (e.g., GetValue() instead of Value property)

## Adding Custom Types

To add support for custom types:

1. Implement the TypeHandler interface:
```go
type MyTypeHandler struct{}

func (h *MyTypeHandler) ConvertValue(value interface{}) (interface{}, error) {
    // Implement conversion logic
}
```

2. Register the handler:
```go
registry.RegisterHandler(reflect.TypeOf(MyType{}), &MyTypeHandler{})
