package patchfields

import "time"

// PatchPropertyNullableTime represents a nullable time property (equivalent to C# nullable DateTime)
type PatchFieldNullableTime struct {
	value *time.Time
}

func (p *PatchFieldNullableTime) GetValue() *time.Time {
	return p.value
}

func (p *PatchFieldNullableTime) SetValue(value *time.Time) {
	p.value = value
}
