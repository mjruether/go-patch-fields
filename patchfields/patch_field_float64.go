package patchfields

// PatchPropertyFloat64 represents a 64-bit floating-point property (equivalent to C# double)
type PatchFieldFloat64 struct {
	value float64
}

func (p *PatchFieldFloat64) GetValue() interface{} {
	return p.value
}

func (p *PatchFieldFloat64) SetValue(value float64) {
	p.value = value
}
