package patchfields

// PatchPropertyInt64 represents a 64-bit integer property (equivalent to C# long)
type PatchFieldInt64 struct {
	value int64
}

func (p *PatchFieldInt64) GetValue() interface{} {
	return p.value
}

func (p *PatchFieldInt64) SetValue(value int64) {
	p.value = value
}
