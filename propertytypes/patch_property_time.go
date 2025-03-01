package propertytypes

import "time"

// PatchPropertyTime represents a time property
type PatchPropertyTime struct {
	Value time.Time
}

func (p *PatchPropertyTime) GetValue() interface{} {
	return p.Value
}

func (p *PatchPropertyTime) SetValue(value time.Time) {
	p.Value = value
}
