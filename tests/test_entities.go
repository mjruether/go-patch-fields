package tests

import "time"

// TestEntity represents a test entity with various properties
type TestEntity struct {
	// Int types
	SomeInt         int
	SomeInt64       int64
	SomeNullableInt *int
	MisTypedInt     int

	// String type
	SomeString         string
	SomeNullableString *string

	// Bool type
	SomeBool         bool
	SomeNullableBool *bool

	// Float type
	SomeFloat64         float64
	SomeNullableFloat64 *float64

	// Time type
	SomeTime         time.Time
	SomeNullableTime *time.Time
}
