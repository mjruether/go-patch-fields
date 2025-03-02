package main

import (
	"fmt"
	"time"

	"github.com/mjruether/go-patch-fields/gopatchfields"
	"github.com/mjruether/go-patch-fields/gopatchservices"
)

// Example entity we want to patch
type User struct {
	CreatedAt time.Time // 24 bytes
	Name      string    // 16 bytes
	ID        int       // 8 bytes
	Active    bool      // 1 byte + padding
}

// Example patch model
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
		Name:      &gopatchfields.PatchFieldString{Value: "Updated Name"},
		Active:    &gopatchfields.PatchFieldBool{Value: false},
		CreatedAt: nil, // This field will not be updated
	}

	fmt.Printf("Before patch:\n%+v\n\n", user)

	// Apply the patch
	result := service.SetValues(user, patch)

	if user, ok := result.Entity.(*User); ok {
		fmt.Printf("After patch:\n%+v\n\n", user)
	} else {
		fmt.Println("Error: Invalid type assertion")
	}

	if len(result.Warnings) > 0 {
		fmt.Println("Warnings:")
		for field, warning := range result.Warnings {
			fmt.Printf("%s: %s\n", field, warning)
		}
	}
}
