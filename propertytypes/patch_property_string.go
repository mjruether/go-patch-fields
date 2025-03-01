package propertytypes

// PatchPropertyString represents a string property
type PatchPropertyString struct {
	Value string
}

func (p *PatchPropertyString) GetValue() interface{} {
	return p.Value
}

func (p *PatchPropertyString) SetValue(value string) {
	p.Value = value
}
