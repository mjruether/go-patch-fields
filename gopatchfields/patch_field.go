package gopatchfields

// PatchFieldBase defines the non-generic base interface for all patch fields
type PatchFieldBase interface {
	GetValue() interface{}
}

// PatchField defines the generic interface for patch field types
type PatchField[T any] interface {
	PatchFieldBase
	SetValue(value T)
	GetTypedValue() T
}
