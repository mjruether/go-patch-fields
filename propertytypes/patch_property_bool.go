package propertytypes

// PatchPropertyBool represents a boolean property
type PatchPropertyBool struct {
	Value bool
}

func (p *PatchPropertyBool) GetValue() interface{} {
	return p.Value
}

func (p *PatchPropertyBool) SetValue(value bool) {
	p.Value = value
}
