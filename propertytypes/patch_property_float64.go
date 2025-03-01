package propertytypes

// PatchPropertyFloat64 represents a 64-bit floating-point property (equivalent to C# double)
type PatchPropertyFloat64 struct {
	value float64
}

func (p *PatchPropertyFloat64) GetValue() float64 {
	return p.value
}

func (p *PatchPropertyFloat64) SetValue(value float64) {
	p.value = value
}
