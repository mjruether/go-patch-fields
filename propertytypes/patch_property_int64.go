package propertytypes

// PatchPropertyInt64 represents a 64-bit integer property (equivalent to C# long)
type PatchPropertyInt64 struct {
	value int64
}

func (p *PatchPropertyInt64) GetValue() interface{} {
	return p.value
}

func (p *PatchPropertyInt64) SetValue(value int64) {
	p.value = value
}
