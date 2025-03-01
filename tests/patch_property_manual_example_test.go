package tests

import (
	"testing"

	"github.com/mjruether/go-patch-fields/gopatchfields"
)

// TestPatchInt_NonNullField_ManuallyChecked_SetValue demonstrates how to use the PatchFieldInt
// type to apply updates manually. This is more efficient than using reflection-based approaches.
func TestPatchInt_NonNullField_ManuallyChecked_SetValue(t *testing.T) {
	newValue := 1

	entityToUpdate := &TestEntity{} // Zero value for int is 0

	patchModel := &IntPatchModel{
		SomeInt: &gopatchfields.PatchFieldInt{},
	}
	patchModel.SomeInt.SetValue(newValue)

	if got := entityToUpdate.SomeInt; got == newValue {
		t.Errorf("Before patch: expected != %v, got %v", newValue, got)
	}

	if patchModel.SomeInt != nil {
		entityToUpdate.SomeInt = patchModel.SomeInt.GetTypedValue()
	}

	if got := entityToUpdate.SomeInt; got != newValue {
		t.Errorf("After patch: expected %v, got %v", newValue, got)
	}
}

func TestPatchInt_NullField_ManuallyChecked_DontSetValue(t *testing.T) {
	newValue := 1
	initialValue := 0

	entityToUpdate := &TestEntity{} // Zero value for int is 0

	patchModel := &IntPatchModel{
		// SomeInt field is not set, and therefore nil on the incoming request model
	}

	if got := entityToUpdate.SomeInt; got != initialValue {
		t.Errorf("Before patch: expected %v, got %v", initialValue, got)
	}

	if patchModel.SomeInt != nil {
		entityToUpdate.SomeInt = patchModel.SomeInt.GetTypedValue()
	}

	if got := entityToUpdate.SomeInt; got != initialValue {
		t.Errorf("After patch: expected %v, got %v", initialValue, got)
	}

	if got := entityToUpdate.SomeInt; got == newValue {
		t.Errorf("Value should not be %v", newValue)
	}
}
