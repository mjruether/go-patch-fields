package tests

import "time"

// TestEntity represents a test entity with various properties
type TestEntity struct {
	SomeTime            time.Time
	SomeNullableTime    *time.Time
	SomeNullableFloat64 *float64
	SomeNullableString  *string
	SomeNullableInt     *int
	SomeNullableBool    *bool
	SomeString          string
	SomeInt64           int64
	SomeFloat64         float64
	SomeInt             int
	MisTypedInt         int
	SomeBool            bool
}
