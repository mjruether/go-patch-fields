package patchfields

// PatchPropertyBase defines the non-generic base interface for all patch properties
type PatchFieldBase interface {
	GetValue() interface{}
}

// PatchProperty defines the generic interface for patch property types
type PatchField[T any] interface {
	PatchFieldBase
	SetValue(value T)
	GetTypedValue() T
}
