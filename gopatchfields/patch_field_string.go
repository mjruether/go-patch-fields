package gopatchfields

// PatchFieldString represents a string field
type PatchFieldString struct {
	Value string
}

func (p *PatchFieldString) GetValue() interface{} {
	return p.Value
}

func (p *PatchFieldString) SetValue(value string) {
	p.Value = value
}
