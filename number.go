package go_objectutils

import (
	"strconv"
)

// NumberConstraint defines numeric types that we can return.
// We include float64 as the default for JSON numbers, and standard int types.
type NumberConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func convertToNumber[T NumberConstraint](val interface{}) (T, bool) {
	var zero T
	switch v := val.(type) {
	case float64:
		return T(v), true
	case int:
		return T(v), true
	case int64:
		return T(v), true
	case float32:
		return T(v), true
	case string:
		// Try parsing string as float
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return T(f), true
		}
	}
	// Fallback using reflection/fmt if needed, or strict.
	// JSON decoder usually gives float64 for numbers.
	// We handle the common cases above.
	return zero, false
}

// GetNumberPropOrDefault retrieves a numeric property or returns a default value.
func GetNumberPropOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue T) T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if numVal, ok := convertToNumber[T](val); ok {
		return numVal
	}
	return defaultValue
}

// GetNumberPropOrDefaultFunction retrieves a numeric property or returns a value from a default function.
func GetNumberPropOrDefaultFunction[T NumberConstraint](props map[string]interface{}, prop string, defaultFunction func() T) T {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	if numVal, ok := convertToNumber[T](val); ok {
		return numVal
	}
	return defaultFunction()
}

// GetNumberPropOrThrow retrieves a numeric property or panics if missing/invalid.
func GetNumberPropOrThrow[T NumberConstraint](props map[string]interface{}, prop string, message ...string) T {
	msg := "Property " + prop + " is missing or not a number"
	if len(message) > 0 {
		msg = message[0]
	}
	if props == nil {
		panic(msg)
	}
	val, ok := props[prop]
	if !ok {
		panic(msg)
	}
	if numVal, ok := convertToNumber[T](val); ok {
		return numVal
	}
	panic(msg)
}
