package gopatchfields

// PatchFieldNullableInt represents a nullable integer field
type PatchFieldNullableInt struct {
	value *int
}

func (p *PatchFieldNullableInt) GetValue() *int {
	return p.value
}

func (p *PatchFieldNullableInt) SetValue(value *int) {
	p.value = value
}
