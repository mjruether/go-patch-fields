package tests

import "github.com/mjruether/patchpropertiesgo/propertytypes"

// IntPatchModel represents a patch model for testing int properties
type IntPatchModel struct {
	SomeInt         *propertytypes.PatchPropertyInt
	SomeNullableInt *propertytypes.PatchPropertyInt
	MisTypedInt     *propertytypes.PatchPropertyString
}

// Int64PatchModel represents a patch model for testing int64 properties
type Int64PatchModel struct {
	SomeInt64 *propertytypes.PatchPropertyInt64
}

// StringPatchModel represents a patch model for testing string properties
type StringPatchModel struct {
	SomeString         *propertytypes.PatchPropertyString
	SomeNullableString *propertytypes.PatchPropertyString
}

// BoolPatchModel represents a patch model for testing bool properties
type BoolPatchModel struct {
	SomeBool         *propertytypes.PatchPropertyBool
	SomeNullableBool *propertytypes.PatchPropertyBool
}

// Float64PatchModel represents a patch model for testing float64 properties
type Float64PatchModel struct {
	SomeFloat64         *propertytypes.PatchPropertyFloat64
	SomeNullableFloat64 *propertytypes.PatchPropertyFloat64
}

// TimePatchModel represents a patch model for testing time properties
type TimePatchModel struct {
	SomeTime         *propertytypes.PatchPropertyTime
	SomeNullableTime *propertytypes.PatchPropertyTime
}

// MisNamedPatchModel represents a patch model with mismatched property names
type MisNamedPatchModel struct {
	MisNamedInt *propertytypes.PatchPropertyInt
}
