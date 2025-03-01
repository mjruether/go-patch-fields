package patchfields

import "time"

// PatchPropertyTime represents a time property
type PatchFieldTime struct {
	Value time.Time
}

func (p *PatchFieldTime) GetValue() interface{} {
	return p.Value
}

func (p *PatchFieldTime) SetValue(value time.Time) {
	p.Value = value
}
