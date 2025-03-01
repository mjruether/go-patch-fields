package gopatchfields

import "time"

// PatchFieldNullableTime represents a nullable time field (equivalent to C# nullable DateTime)
type PatchFieldNullableTime struct {
	value *time.Time
}

func (p *PatchFieldNullableTime) GetValue() *time.Time {
	return p.value
}

func (p *PatchFieldNullableTime) SetValue(value *time.Time) {
	p.value = value
}
