// Package gopatchfields provides a type-safe and structured approach to HTTP PATCH operations.
//
// This library offers helper structs and a mapper service that solve the challenge of
// inconsistent PATCH endpoint implementations. It enables selective and type-safe
// updates to struct fields using wrapper types.
//
// Key features:
//   - Type-safe property patching with explicit value wrapping
//   - Support for common Go types (bool, int, string, time.Time, etc.)
//   - Automatic type handler registration
//   - Configurable handling of unmatched properties
//   - Detailed warning and error reporting
//
// Example usage:
//
//	type User struct {
//	    Name      string
//	    Active    bool
//	}
//
//	type UserPatch struct {
//	    Name   *patchfields.PatchFieldString
//	    Active *patchfields.PatchFieldBool
//	}
//
//	service := services.NewPatchFieldService()
//	result := service.SetValues(user, patch)
package gopatchfields
