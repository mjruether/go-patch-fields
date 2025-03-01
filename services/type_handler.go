package services

import "reflect"

// TypeHandler defines the interface for converting values between types
type TypeHandler interface {
	ConvertValue(value interface{}) (interface{}, error)
}

// TypeRegistry maintains a mapping of types to their handlers
type TypeRegistry struct {
	handlers map[reflect.Type]TypeHandler
}

// NewTypeRegistry creates a new TypeRegistry with initialized maps
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		handlers: make(map[reflect.Type]TypeHandler),
	}
}

// RegisterHandler registers a TypeHandler for a specific type
func (r *TypeRegistry) RegisterHandler(t reflect.Type, handler TypeHandler) {
	r.handlers[t] = handler
}

// GetHandler retrieves a TypeHandler for a specific type
func (r *TypeRegistry) GetHandler(t reflect.Type) (TypeHandler, bool) {
	handler, exists := r.handlers[t]
	return handler, exists
}
