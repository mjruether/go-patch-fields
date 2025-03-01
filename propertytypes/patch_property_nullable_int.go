package propertytypes

// PatchPropertyNullableInt represents a nullable integer property
type PatchPropertyNullableInt struct {
	value *int
}

func (p *PatchPropertyNullableInt) GetValue() *int {
	return p.value
}

func (p *PatchPropertyNullableInt) SetValue(value *int) {
	p.value = value
}
