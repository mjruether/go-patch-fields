package patchfields

// PatchPropertyInt represents an integer property and implements PatchProperty[int]
type PatchFieldInt struct {
	value int
}

func (p *PatchFieldInt) GetValue() interface{} {
	return p.value
}

func (p *PatchFieldInt) GetTypedValue() int {
	return p.value
}

func (p *PatchFieldInt) SetValue(value int) {
	p.value = value
}
