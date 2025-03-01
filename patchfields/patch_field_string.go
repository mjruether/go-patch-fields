package patchfields

// PatchPropertyString represents a string property
type PatchFieldString struct {
	Value string
}

func (p *PatchFieldString) GetValue() interface{} {
	return p.Value
}

func (p *PatchFieldString) SetValue(value string) {
	p.Value = value
}
