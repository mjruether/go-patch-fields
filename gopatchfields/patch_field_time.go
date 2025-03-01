package gopatchfields

import "time"

// PatchFieldTime represents a time field
type PatchFieldTime struct {
	Value time.Time
}

func (p *PatchFieldTime) GetValue() interface{} {
	return p.Value
}

func (p *PatchFieldTime) SetValue(value time.Time) {
	p.Value = value
}
