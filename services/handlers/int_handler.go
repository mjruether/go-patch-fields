package handlers

import (
	"fmt"
	"strconv"
)

// IntHandler handles conversions to int
type IntHandler struct{}

func (h *IntHandler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case float64:
		return int(v), nil
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int: %v", err)
		}
		return i, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to int", value)
	}
}

// Int64Handler handles conversions to int64
type Int64Handler struct{}

func (h *Int64Handler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int64: %v", err)
		}
		return i, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to int64", value)
	}
}
