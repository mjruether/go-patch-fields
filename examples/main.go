package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mjruether/patchpropertiesgo/propertytypes"
	"github.com/mjruether/patchpropertiesgo/services"
	"github.com/mjruether/patchpropertiesgo/services/handlers"
)

// Example entity we want to patch
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Active    bool
}

// Example patch model
type UserPatch struct {
	Name      *propertytypes.PatchPropertyString
	CreatedAt *propertytypes.PatchPropertyTime
	Active    *propertytypes.PatchPropertyBool
}

func main() {
	// Create service and register type handlers
	service := services.NewPatchPropertyService()

	// Register type handlers
	registry := service.GetTypeRegistry()
	registry.RegisterHandler(reflect.TypeOf(""), &handlers.StringHandler{})
	registry.RegisterHandler(reflect.TypeOf(time.Time{}), &handlers.TimeHandler{})
	registry.RegisterHandler(reflect.TypeOf(true), &handlers.BoolHandler{})
	registry.RegisterHandler(reflect.TypeOf(int(0)), &handlers.IntHandler{})

	// Create an example user
	user := &User{
		ID:        1,
		Name:      "Original Name",
		CreatedAt: time.Now(),
		Active:    true,
	}

	// Create a patch with some changes
	patch := &UserPatch{
		Name:   &propertytypes.PatchPropertyString{Value: "Updated Name"},
		Active: &propertytypes.PatchPropertyBool{Value: false},
	}

	fmt.Printf("Before patch:\n%+v\n\n", user)

	// Apply the patch
	result := service.SetValues(user, patch)

	fmt.Printf("After patch:\n%+v\n\n", result.Entity.(*User))

	if len(result.Warnings) > 0 {
		fmt.Println("Warnings:")
		for field, warning := range result.Warnings {
			fmt.Printf("%s: %s\n", field, warning)
		}
	}
}
