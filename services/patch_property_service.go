package services

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mjruether/go-patch-fields/services/handlers"
)

// PatchFieldService handles property patching operations
type PatchFieldService struct {
	ignoreUnmatchedProperties bool
	typeRegistry              *TypeRegistry
}

// NewPatchFieldService creates a new PatchFieldService
func NewPatchFieldService() *PatchFieldService {
	service := &PatchFieldService{
		ignoreUnmatchedProperties: false,
		typeRegistry:              NewTypeRegistry(),
	}

	// Register handlers for basic types
	intHandler := &handlers.IntHandler{}
	// Register both the value and pointer types for int
	service.typeRegistry.RegisterHandler(reflect.TypeOf(int(0)), intHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*int)(nil)).Elem(), intHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*int)(nil)), intHandler)

	int64Handler := &handlers.Int64Handler{}
	service.typeRegistry.RegisterHandler(reflect.TypeOf(int64(0)), int64Handler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*int64)(nil)).Elem(), int64Handler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*int64)(nil)), int64Handler)

	stringHandler := &handlers.StringHandler{}
	service.typeRegistry.RegisterHandler(reflect.TypeOf(""), stringHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*string)(nil)).Elem(), stringHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*string)(nil)), stringHandler)

	float64Handler := &handlers.Float64Handler{}
	service.typeRegistry.RegisterHandler(reflect.TypeOf(float64(0)), float64Handler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*float64)(nil)).Elem(), float64Handler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*float64)(nil)), float64Handler)

	boolHandler := &handlers.BoolHandler{}
	service.typeRegistry.RegisterHandler(reflect.TypeOf(false), boolHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*bool)(nil)).Elem(), boolHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*bool)(nil)), boolHandler)

	timeHandler := &handlers.TimeHandler{}
	service.typeRegistry.RegisterHandler(reflect.TypeOf(time.Time{}), timeHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*time.Time)(nil)).Elem(), timeHandler)
	service.typeRegistry.RegisterHandler(reflect.TypeOf((*time.Time)(nil)), timeHandler)

	return service
}

// GetTypeRegistry returns the service's type registry
func (s *PatchFieldService) GetTypeRegistry() *TypeRegistry {
	return s.typeRegistry
}

// SetIgnoreUnmatchedProperties configures whether to ignore unmatched properties
func (s *PatchFieldService) SetIgnoreUnmatchedProperties(ignore bool) {
	s.ignoreUnmatchedProperties = ignore
}

// SetValues applies patches from a patch model to an entity
func (s *PatchFieldService) SetValues(entity interface{}, patchModel interface{}) *SetValuesResponse {
	if patchModel == nil {
		return NewSetValuesResponse(entity)
	}

	response := NewSetValuesResponse(entity)
	entityValue := reflect.ValueOf(entity)
	if entityValue.Kind() == reflect.Ptr {
		entityValue = entityValue.Elem()
	}
	if entityValue.Kind() != reflect.Struct {
		response.Warnings["error"] = "entity must be a struct or pointer to struct"
		return response
	}

	patchValue := reflect.ValueOf(patchModel)
	if patchValue.Kind() == reflect.Ptr {
		patchValue = patchValue.Elem()
	}
	if patchValue.Kind() != reflect.Struct {
		response.Warnings["error"] = "patchModel must be a struct or pointer to struct"
		return response
	}

	patchType := patchValue.Type()
	for i := 0; i < patchValue.NumField(); i++ {
		field := patchValue.Field(i)
		fieldType := patchType.Field(i)

		if !field.IsValid() || field.IsNil() {
			// Handle case where field is nil
			entityField := entityValue.FieldByName(fieldType.Name)
			if entityField.IsValid() && entityField.CanSet() && entityField.Kind() == reflect.Ptr {
				entityField.Set(reflect.Zero(entityField.Type()))
			}
			continue
		}

		// Check if field implements PatchProperty interface
		if !field.Type().Implements(reflect.TypeOf((*interface{ GetValue() interface{} })(nil)).Elem()) {
			continue
		}

		// Get the patch property value using the GetValue method
		getValue := field.MethodByName("GetValue")
		if !getValue.IsValid() {
			continue
		}

		value := getValue.Call(nil)[0].Interface()
		if value == nil {
			// Handle case where we want to set a nil value
			entityField := entityValue.FieldByName(fieldType.Name)
			if entityField.IsValid() && entityField.CanSet() && entityField.Kind() == reflect.Ptr {
				entityField.Set(reflect.Zero(entityField.Type()))
			}
			continue
		}

		// Find matching field in entity
		entityField := entityValue.FieldByName(fieldType.Name)
		if !entityField.IsValid() {
			msg := fmt.Sprintf("%s did not match any properties of the entity to update", fieldType.Name)
			if s.ignoreUnmatchedProperties {
				continue
			}
			response.Warnings[fieldType.Name] = msg
			continue
		}

		if !entityField.CanSet() {
			response.Warnings[fieldType.Name] = fmt.Sprintf("%s cannot be set", fieldType.Name)
			continue
		}

		// Try to convert and set the value
		if handler, exists := s.typeRegistry.GetHandler(entityField.Type()); exists {
			convertedValue, err := handler.ConvertValue(value)
			if err != nil {
				response.Warnings[fieldType.Name] = fmt.Sprintf("failed to convert value: %v", err)
				continue
			}

			if entityField.Kind() == reflect.Ptr {
				// For pointer fields, create a new pointer to the converted value
				ptr := reflect.New(entityField.Type().Elem())
				ptr.Elem().Set(reflect.ValueOf(convertedValue))
				entityField.Set(ptr)
			} else {
				entityField.Set(reflect.ValueOf(convertedValue))
			}
		} else {
			// Fallback to basic type conversion
			if entityFieldType := entityField.Type(); entityFieldType != reflect.TypeOf(value) {
				response.Warnings[fieldType.Name] = fmt.Sprintf(
					"expected %v but received %v",
					entityFieldType,
					reflect.TypeOf(value),
				)
				continue
			}

			if entityField.Kind() == reflect.Ptr {
				// For pointer fields, create a new pointer to the value
				ptr := reflect.New(entityField.Type().Elem())
				ptr.Elem().Set(reflect.ValueOf(value))
				entityField.Set(ptr)
			} else {
				entityField.Set(reflect.ValueOf(value))
			}
		}
	}

	return response
}
