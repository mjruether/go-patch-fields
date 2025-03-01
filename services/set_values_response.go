package services

// SetValuesResponse represents the result of a patch operation
type SetValuesResponse struct {
	Entity   interface{}       // The updated entity
	Warnings map[string]string // Warnings encountered during patching
}

// NewSetValuesResponse creates a new SetValuesResponse with initialized maps
func NewSetValuesResponse(entity interface{}) *SetValuesResponse {
	return &SetValuesResponse{
		Entity:   entity,
		Warnings: make(map[string]string),
	}
}
