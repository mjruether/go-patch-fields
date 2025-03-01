package propertytypes

// PatchPropertyBase defines the non-generic base interface for all patch properties
type PatchPropertyBase interface {
	GetValue() interface{}
}

// PatchProperty defines the generic interface for patch property types
type PatchProperty[T any] interface {
	PatchPropertyBase
	SetValue(value T)
	GetTypedValue() T
}
