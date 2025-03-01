package tests

import (
	"testing"
	"time"

	"github.com/mjruether/patchpropertiesgo/patchfields"
	"github.com/mjruether/patchpropertiesgo/services"
)

func TestCanPatchInt64(t *testing.T) {
	service := services.NewPatchFieldService()
	newValue := int64(100)
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldInt64{}
	patchProperty.SetValue(newValue)

	patchModel := &Int64PatchModel{
		SomeInt64: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := "test string"
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldString{}
	patchProperty.SetValue(newValue)

	patchModel := &StringPatchModel{
		SomeString: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := "test string"
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldString{}
	patchProperty.SetValue(newValue)

	patchModel := &StringPatchModel{
		SomeNullableString: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := true
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldBool{}
	patchProperty.SetValue(newValue)

	patchModel := &BoolPatchModel{
		SomeBool: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := true
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldBool{}
	patchProperty.SetValue(newValue)

	patchModel := &BoolPatchModel{
		SomeNullableBool: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := 123.456
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldFloat64{}
	patchProperty.SetValue(newValue)

	patchModel := &Float64PatchModel{
		SomeFloat64: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := 123.456
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldFloat64{}
	patchProperty.SetValue(newValue)

	patchModel := &Float64PatchModel{
		SomeNullableFloat64: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldTime{}
	patchProperty.SetValue(newValue)

	patchModel := &TimePatchModel{
		SomeTime: patchProperty,
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
	service := services.NewPatchFieldService()
	newValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldTime{}
	patchProperty.SetValue(newValue)

	patchModel := &TimePatchModel{
		SomeNullableTime: patchProperty,
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
	service := services.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldInt{}
	patchProperty.SetValue(newValue)

	patchModel := &IntPatchModel{
		SomeInt: patchProperty,
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
	service := services.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldInt{}
	patchProperty.SetValue(newValue)

	patchModel := &IntPatchModel{
		SomeNullableInt: patchProperty,
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
	service := services.NewPatchFieldService()

	initialValue := 0
	entityToUpdate := &TestEntity{
		SomeNullableInt: &initialValue,
	}

	patchModel := &IntPatchModel{
		SomeNullableInt: &patchfields.PatchFieldInt{},
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
	service := services.NewPatchFieldService()

	newValue := 1

	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldInt{}
	patchProperty.SetValue(newValue)

	patchModel := &MisNamedPatchModel{
		MisNamedInt: patchProperty,
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
	service := services.NewPatchFieldService()

	initialValue := "test"
	entityToUpdate := &TestEntity{
		SomeNullableString: &initialValue,
	}

	patchModel := &StringPatchModel{
		SomeNullableString: &patchfields.PatchFieldString{},
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
	service := services.NewPatchFieldService()

	initialValue := true
	entityToUpdate := &TestEntity{
		SomeNullableBool: &initialValue,
	}

	patchModel := &BoolPatchModel{
		SomeNullableBool: &patchfields.PatchFieldBool{},
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
	service := services.NewPatchFieldService()

	initialValue := 123.456
	entityToUpdate := &TestEntity{
		SomeNullableFloat64: &initialValue,
	}

	patchModel := &Float64PatchModel{
		SomeNullableFloat64: &patchfields.PatchFieldFloat64{},
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
	service := services.NewPatchFieldService()

	initialValue := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	entityToUpdate := &TestEntity{
		SomeNullableTime: &initialValue,
	}

	patchModel := &TimePatchModel{
		SomeNullableTime: &patchfields.PatchFieldTime{},
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
	service := services.NewPatchFieldService()

	newValue := "asdf"

	entityToUpdate := &TestEntity{}

	patchProperty := &patchfields.PatchFieldString{}
	patchProperty.SetValue(newValue)

	patchModel := &IntPatchModel{
		MisTypedInt: patchProperty,
	}

	response := service.SetValues(entityToUpdate, patchModel)

	if len(response.Warnings) != 1 {
		t.Errorf("Expected 1 warning, got %d", len(response.Warnings))
	}

	if _, exists := response.Warnings["MisTypedInt"]; !exists {
		t.Error("Expected warning for 'MisTypedInt'")
	}
}
