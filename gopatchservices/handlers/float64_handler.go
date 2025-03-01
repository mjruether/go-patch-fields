package handlers

import (
	"fmt"
	"strconv"
)

// Float64Handler handles conversions to float64
type Float64Handler struct{}

func (h *Float64Handler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to float64: %v", err)
		}
		return f, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to float64", value)
	}
}
