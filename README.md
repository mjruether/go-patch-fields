# PatchPropertiesGo

A set of helper structs and an optional mapper service providing an opinionated and simple approach to HTTP PATCH operations. This library solves the challenge of inconsistent PATCH endpoint implementations by offering a type-safe and structured way to selectively update struct fields using wrapper types.

## Features

- Type-safe property patching with explicit value wrapping
- Built-in support for common Go types:
  - Basic types (bool, int, int64, float64, string)
  - time.Time with timezone handling
  - Nullable fields via pointer types
- Automatic type handler registration for built-in types
- Customizable type conversion system for custom types
- Configurable handling of unmatched properties
- Detailed warning and error reporting

## Example Usage

```go
// Define your entity
type User struct {
    ID        int
    Name      string
    CreatedAt time.Time
    Active    bool
}

// Define your patch model using wrapper types
type UserPatch struct {
    Name      *PatchPropertyString    // Can be nil if no update needed
    CreatedAt *PatchPropertyTime     
    Active    *PatchPropertyBool
}

// Create and use the service (handlers auto-registered)
service := services.NewPatchFieldService()

// Optional: Configure unmatched property handling
service.SetIgnoreUnmatchedProperties(true)

// Create a patch with selective updates
patch := &UserPatch{
    Name:   &PatchPropertyString{Value: "Updated Name"},
    Active: &PatchPropertyBool{Value: false},
    // CreatedAt omitted - won't be updated
}

// Apply the patch
result := service.SetValues(user, patch)

// Handle warnings if any occurred
if len(result.Warnings) > 0 {
    for field, warning := range result.Warnings {
        fmt.Printf("Warning for %s: %s\n", field, warning)
    }
}

// Cast back to concrete type if needed
updatedUser := result.Entity.(*User)
```

## Implementation Details

### Patch Property Pattern

The library uses a generic interface pattern for type-safe property wrapping:

```go
type PatchFieldBase interface {
    GetValue() interface{}
}

type PatchField[T any] interface {
    PatchPropertyBase
    SetValue(value T)
    GetTypedValue() T
}
```

Each property type (string, int, etc.) implements this interface to provide type safety and proper value handling.

### Type Handling System

The service uses a type registry and handler system for value conversion:

1. Built-in handlers are automatically registered for common types
2. Custom types can be supported by implementing the TypeHandler interface:

```go
type MyTypeHandler struct{}

func (h *MyTypeHandler) ConvertValue(value interface{}) (interface{}, error) {
    // Implement conversion logic
    return convertedValue, nil
}

// Register the handler
registry := service.GetTypeRegistry()
registry.RegisterHandler(reflect.TypeOf(MyType{}), &MyTypeHandler{})
```

### Key Features

1. **Selective Updates**: Only fields present in the patch model are updated
2. **Null Handling**: 
   - Nil pointer in patch model = no update
   - Nil value in property = sets field to zero/nil
3. **Type Safety**: Each field uses a specific wrapper type matching its target
4. **Validation**: Type compatibility checked during patch application
5. **Warning System**: Detailed feedback for unmatched fields or conversion issues

### Reflection Usage

The service uses reflection to:
- Match patch fields to entity fields by name
- Handle both pointer and value types appropriately
- Perform type-safe value conversion and assignment
- Support dynamic property access and modification

## Continuous Integration

This project uses GitHub Actions to run tests automatically on every push and pull request to the main branch. Repository members will automatically receive notifications through GitHub when tests fail.

To configure notifications:
1. Go to your GitHub Settings > Notifications
2. Under "Actions", ensure you have enabled notifications for:
   - "Send notifications for failed workflows on repositories you're watching"
   - "Watch repositories you contribute to"

### Manual Testing

Execute the following from the project's root to manually run test.
```base
go test -v ./tests/...   
```