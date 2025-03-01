package patchfields

// PatchPropertyNullableInt represents a nullable integer property
type PatchFieldNullableInt struct {
	value *int
}

func (p *PatchFieldNullableInt) GetValue() *int {
	return p.value
}

func (p *PatchFieldNullableInt) SetValue(value *int) {
	p.value = value
}
