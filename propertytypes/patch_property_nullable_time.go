package propertytypes

import "time"

// PatchPropertyNullableTime represents a nullable time property (equivalent to C# nullable DateTime)
type PatchPropertyNullableTime struct {
	value *time.Time
}

func (p *PatchPropertyNullableTime) GetValue() *time.Time {
	return p.value
}

func (p *PatchPropertyNullableTime) SetValue(value *time.Time) {
	p.value = value
}
