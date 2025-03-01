package tests

import "github.com/mjruether/go-patch-fields/gopatchfields"

// IntPatchModel represents a patch model for testing int fields
type IntPatchModel struct {
	SomeInt         *gopatchfields.PatchFieldInt
	SomeNullableInt *gopatchfields.PatchFieldInt
	MisTypedInt     *gopatchfields.PatchFieldString
}

// Int64PatchModel represents a patch model for testing int64 fields
type Int64PatchModel struct {
	SomeInt64 *gopatchfields.PatchFieldInt64
}

// StringPatchModel represents a patch model for testing string fields
type StringPatchModel struct {
	SomeString         *gopatchfields.PatchFieldString
	SomeNullableString *gopatchfields.PatchFieldString
}

// BoolPatchModel represents a patch model for testing bool fields
type BoolPatchModel struct {
	SomeBool         *gopatchfields.PatchFieldBool
	SomeNullableBool *gopatchfields.PatchFieldBool
}

// Float64PatchModel represents a patch model for testing float64 fields
type Float64PatchModel struct {
	SomeFloat64         *gopatchfields.PatchFieldFloat64
	SomeNullableFloat64 *gopatchfields.PatchFieldFloat64
}

// TimePatchModel represents a patch model for testing time fields
type TimePatchModel struct {
	SomeTime         *gopatchfields.PatchFieldTime
	SomeNullableTime *gopatchfields.PatchFieldTime
}

// MisNamedPatchModel represents a patch model with mismatched field names
type MisNamedPatchModel struct {
	MisNamedInt *gopatchfields.PatchFieldInt
}
