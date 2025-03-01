package tests

import (
	"testing"
	"time"

	"github.com/mjruether/go-patch-fields/gopatchfields"
	"github.com/mjruether/go-patch-fields/gopatchservices"
)

func TestCanPatchInt64(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := int64(100)
	entityToUpdate := &TestEntity{}

	patchFields := &gopatchfields.PatchFieldInt64{}
	patchFields.SetValue(newValue)

	patchModel := &Int64PatchModel{
		SomeInt64: patchFields,
	}

	if got := entityToUpdate.SomeInt64; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeInt64; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchString(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := "test string"
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldString{}
	patchField.SetValue(newValue)

	patchModel := &StringPatchModel{
		SomeString: patchField,
	}

	if got := entityToUpdate.SomeString; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeString; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableString(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := "test string"
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldString{}
	patchField.SetValue(newValue)

	patchModel := &StringPatchModel{
		SomeNullableString: patchField,
	}

	if entityToUpdate.SomeNullableString != nil {
		t.Error("Before patch: expected nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableString == nil {
		t.Error("After patch: expected non-nil value")
	} else if got := *entityToUpdate.SomeNullableString; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchBool(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := true
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldBool{}
	patchField.SetValue(newValue)

	patchModel := &BoolPatchModel{
		SomeBool: patchField,
	}

	if got := entityToUpdate.SomeBool; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeBool; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableBool(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := true
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldBool{}
	patchField.SetValue(newValue)

	patchModel := &BoolPatchModel{
		SomeNullableBool: patchField,
	}

	if entityToUpdate.SomeNullableBool != nil {
		t.Error("Before patch: expected nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableBool == nil {
		t.Error("After patch: expected non-nil value")
	} else if got := *entityToUpdate.SomeNullableBool; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchFloat64(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := 123.456
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldFloat64{}
	patchField.SetValue(newValue)

	patchModel := &Float64PatchModel{
		SomeFloat64: patchField,
	}

	if got := entityToUpdate.SomeFloat64; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeFloat64; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableFloat64(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := 123.456
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldFloat64{}
	patchField.SetValue(newValue)

	patchModel := &Float64PatchModel{
		SomeNullableFloat64: patchField,
	}

	if entityToUpdate.SomeNullableFloat64 != nil {
		t.Error("Before patch: expected nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableFloat64 == nil {
		t.Error("After patch: expected non-nil value")
	} else if got := *entityToUpdate.SomeNullableFloat64; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchTime(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldTime{}
	patchField.SetValue(newValue)

	patchModel := &TimePatchModel{
		SomeTime: patchField,
	}

	if got := entityToUpdate.SomeTime; got.Equal(newValue) {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeTime; !got.Equal(newValue) {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableTime(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()
	newValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldTime{}
	patchField.SetValue(newValue)

	patchModel := &TimePatchModel{
		SomeNullableTime: patchField,
	}

	if entityToUpdate.SomeNullableTime != nil {
		t.Error("Before patch: expected nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableTime == nil {
		t.Error("After patch: expected non-nil value")
	} else if got := entityToUpdate.SomeNullableTime; !got.Equal(newValue) {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchInt(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldInt{}
	patchField.SetValue(newValue)

	patchModel := &IntPatchModel{
		SomeInt: patchField,
	}

	if got := entityToUpdate.SomeInt; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	service.SetValues(entityToUpdate, patchModel)

	if got := entityToUpdate.SomeInt; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableInt(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldInt{}
	patchField.SetValue(newValue)

	patchModel := &IntPatchModel{
		SomeNullableInt: patchField,
	}

	var currentValue int
	if entityToUpdate.SomeNullableInt != nil {
		currentValue = *entityToUpdate.SomeNullableInt
	}
	if currentValue == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, currentValue)
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableInt == nil {
		t.Error("After patch: expected non-nil value")
	} else if got := *entityToUpdate.SomeNullableInt; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestCanPatchNullableInt_ToNullValue(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	initialValue := 0
	entityToUpdate := &TestEntity{
		SomeNullableInt: &initialValue,
	}

	patchModel := &IntPatchModel{
		SomeNullableInt: &gopatchfields.PatchFieldInt{},
	}
	patchModel.SomeNullableInt = nil // Setting to nil to indicate null value

	if entityToUpdate.SomeNullableInt == nil {
		t.Error("Before patch: expected non-nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableInt != nil {
		t.Errorf("After patch: expected nil, got %v", *entityToUpdate.SomeNullableInt)
	}
}

func TestCannotPatchMisMatchedName(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldInt{}
	patchField.SetValue(newValue)

	patchModel := &MisNamedPatchModel{
		MisNamedInt: patchField,
	}

	response := service.SetValues(entityToUpdate, patchModel)

	if len(response.Warnings) != 1 {
		t.Errorf("Expected 1 warning, got %d", len(response.Warnings))
	}

	if _, exists := response.Warnings["MisNamedInt"]; !exists {
		t.Error("Expected warning for 'MisNamedInt'")
	}
}

func TestCanPatchNullableString_ToNullValue(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	initialValue := "test"
	entityToUpdate := &TestEntity{
		SomeNullableString: &initialValue,
	}

	patchModel := &StringPatchModel{
		SomeNullableString: &gopatchfields.PatchFieldString{},
	}
	patchModel.SomeNullableString = nil // Setting to nil to indicate null value

	if entityToUpdate.SomeNullableString == nil {
		t.Error("Before patch: expected non-nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableString != nil {
		t.Errorf("After patch: expected nil, got %v", *entityToUpdate.SomeNullableString)
	}
}

func TestCanPatchNullableBool_ToNullValue(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	initialValue := true
	entityToUpdate := &TestEntity{
		SomeNullableBool: &initialValue,
	}

	patchModel := &BoolPatchModel{
		SomeNullableBool: &gopatchfields.PatchFieldBool{},
	}
	patchModel.SomeNullableBool = nil // Setting to nil to indicate null value

	if entityToUpdate.SomeNullableBool == nil {
		t.Error("Before patch: expected non-nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableBool != nil {
		t.Errorf("After patch: expected nil, got %v", *entityToUpdate.SomeNullableBool)
	}
}

func TestCanPatchNullableFloat64_ToNullValue(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	initialValue := 123.456
	entityToUpdate := &TestEntity{
		SomeNullableFloat64: &initialValue,
	}

	patchModel := &Float64PatchModel{
		SomeNullableFloat64: &gopatchfields.PatchFieldFloat64{},
	}
	patchModel.SomeNullableFloat64 = nil // Setting to nil to indicate null value

	if entityToUpdate.SomeNullableFloat64 == nil {
		t.Error("Before patch: expected non-nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableFloat64 != nil {
		t.Errorf("After patch: expected nil, got %v", *entityToUpdate.SomeNullableFloat64)
	}
}

func TestCanPatchNullableTime_ToNullValue(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	initialValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{
		SomeNullableTime: &initialValue,
	}

	patchModel := &TimePatchModel{
		SomeNullableTime: &gopatchfields.PatchFieldTime{},
	}
	patchModel.SomeNullableTime = nil // Setting to nil to indicate null value

	if entityToUpdate.SomeNullableTime == nil {
		t.Error("Before patch: expected non-nil value")
	}

	service.SetValues(entityToUpdate, patchModel)

	if entityToUpdate.SomeNullableTime != nil {
		t.Errorf("After patch: expected nil, got %v", *entityToUpdate.SomeNullableTime)
	}
}

func TestCannotPatchMisMatchedType(t *testing.T) {
	service := gopatchservices.NewPatchFieldService()

	newValue := "asdf"

	entityToUpdate := &TestEntity{}

	patchField := &gopatchfields.PatchFieldString{}
	patchField.SetValue(newValue)

	patchModel := &IntPatchModel{
		MisTypedInt: patchField,
	}

	response := service.SetValues(entityToUpdate, patchModel)

	if len(response.Warnings) != 1 {
		t.Errorf("Expected 1 warning, got %d", len(response.Warnings))
	}

	if _, exists := response.Warnings["MisTypedInt"]; !exists {
		t.Error("Expected warning for 'MisTypedInt'")
	}
}
