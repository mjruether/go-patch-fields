package handlers

import (
	"fmt"
	"strconv"
)

// StringHandler handles conversions to string
type StringHandler struct{}

func (h *StringHandler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case fmt.Stringer:
		return v.String(), nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

// BoolHandler handles conversions to bool
type BoolHandler struct{}

func (h *BoolHandler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to bool: %v", err)
		}
		return b, nil
	case int:
		return v != 0, nil
	default:
		return nil, fmt.Errorf("cannot convert %T to bool", value)
	}
}
