package propertytypes

// PatchPropertyInt represents an integer property and implements PatchProperty[int]
type PatchPropertyInt struct {
	value int
}

func (p *PatchPropertyInt) GetValue() interface{} {
	return p.value
}

func (p *PatchPropertyInt) GetTypedValue() int {
	return p.value
}

func (p *PatchPropertyInt) SetValue(value int) {
	p.value = value
}
