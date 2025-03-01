package gopatchfields

// PatchFieldBool represents a boolean field
type PatchFieldBool struct {
	Value bool
}

func (p *PatchFieldBool) GetValue() interface{} {
	return p.Value
}

func (p *PatchFieldBool) SetValue(value bool) {
	p.Value = value
}
