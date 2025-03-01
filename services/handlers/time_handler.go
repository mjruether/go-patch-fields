package handlers

import (
	"fmt"
	"time"
)

// TimeHandler handles conversions to time.Time
type TimeHandler struct{}

func (h *TimeHandler) ConvertValue(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case time.Time:
		return v, nil
	case string:
		// Try multiple common date formats
		formats := []string{
			time.RFC3339,
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
			"2006-01-02",
		}

		for _, format := range formats {
			if t, err := time.Parse(format, v); err == nil {
				return t, nil
			}
		}
		return nil, fmt.Errorf("could not parse time from string: %s", v)

	case int64:
		// Assume Unix timestamp
		return time.Unix(v, 0), nil

	case float64:
		// Assume Unix timestamp with potential fractional seconds
		sec := int64(v)
		nsec := int64((v - float64(sec)) * 1e9)
		return time.Unix(sec, nsec), nil

	default:
		return nil, fmt.Errorf("cannot convert %T to time.Time", value)
	}
}

// NullableTimeHandler handles conversions to *time.Time
type NullableTimeHandler struct {
	timeHandler TimeHandler
}

func (h *NullableTimeHandler) ConvertValue(value interface{}) (interface{}, error) {
	if value == nil {
		return (*time.Time)(nil), nil
	}

	result, err := h.timeHandler.ConvertValue(value)
	if err != nil {
		return nil, err
	}

	t := result.(time.Time)
	return &t, nil
}
