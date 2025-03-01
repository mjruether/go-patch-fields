package tests

import "github.com/mjruether/go-patch-fields/patchfields"

// IntPatchModel represents a patch model for testing int properties
type IntPatchModel struct {
	SomeInt         *patchfields.PatchFieldInt
	SomeNullableInt *patchfields.PatchFieldInt
	MisTypedInt     *patchfields.PatchFieldString
}

// Int64PatchModel represents a patch model for testing int64 properties
type Int64PatchModel struct {
	SomeInt64 *patchfields.PatchFieldInt64
}

// StringPatchModel represents a patch model for testing string properties
type StringPatchModel struct {
	SomeString         *patchfields.PatchFieldString
	SomeNullableString *patchfields.PatchFieldString
}

// BoolPatchModel represents a patch model for testing bool properties
type BoolPatchModel struct {
	SomeBool         *patchfields.PatchFieldBool
	SomeNullableBool *patchfields.PatchFieldBool
}

// Float64PatchModel represents a patch model for testing float64 properties
type Float64PatchModel struct {
	SomeFloat64         *patchfields.PatchFieldFloat64
	SomeNullableFloat64 *patchfields.PatchFieldFloat64
}

// TimePatchModel represents a patch model for testing time properties
type TimePatchModel struct {
	SomeTime         *patchfields.PatchFieldTime
	SomeNullableTime *patchfields.PatchFieldTime
}

// MisNamedPatchModel represents a patch model with mismatched property names
type MisNamedPatchModel struct {
	MisNamedInt *patchfields.PatchFieldInt
}
