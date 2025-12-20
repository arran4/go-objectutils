package go_objectutils

import (
	"fmt"
	"strconv"
)

// NumberConstraint defines numeric types that we can return.
type NumberConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func convertToNumber[T NumberConstraint](val interface{}) (T, error) {
	var zero T
	switch v := val.(type) {
	case float64:
		return T(v), nil
	case int:
		return T(v), nil
	case int64:
		return T(v), nil
	case float32:
		return T(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return zero, err
		}
		return T(f), nil
	}
	return zero, fmt.Errorf("cannot convert %T to number", val)
}

// GetNumber retrieves a numeric property.
func GetNumber[T NumberConstraint](props map[string]interface{}, prop string) (T, error) {
	var zero T
	if props == nil {
		return zero, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return zero, &MissingFieldError{Prop: prop}
	}
	if numVal, err := convertToNumber[T](val); err == nil {
		return numVal, nil
	} else {
		return zero, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("%T", zero), Actual: val, Cause: err}
	}
}

// MustGetNumber retrieves a numeric property or panics.
func MustGetNumber[T NumberConstraint](props map[string]interface{}, prop string) T {
	val, err := GetNumber[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetNumberOrDefault retrieves a numeric property or returns a default value.
func GetNumberOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue T) T {
	val, err := GetNumber[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetNumberPtr retrieves a numeric property as a pointer.
func GetNumberPtr[T NumberConstraint](props map[string]interface{}, prop string) (*T, error) {
	val, err := GetNumber[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetNumberPtr retrieves a numeric property as a pointer or panics.
func MustGetNumberPtr[T NumberConstraint](props map[string]interface{}, prop string) *T {
	val, err := GetNumberPtr[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetNumberPtrOrDefault retrieves a numeric property as a pointer or returns a default value.
func GetNumberPtrOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue *T) *T {
	val, err := GetNumberPtr[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy Aliases

// GetNumberPropOrDefault is an alias for GetNumberOrDefault.
func GetNumberPropOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue T) T {
	return GetNumberOrDefault(props, prop, defaultValue)
}

// GetNumberPropOrDefaultFunction retrieves a numeric property or returns a value from a default function.
func GetNumberPropOrDefaultFunction[T NumberConstraint](props map[string]interface{}, prop string, defaultFunction func() T) T {
	val, err := GetNumber[T](props, prop)
	if err != nil {
		return defaultFunction()
	}
	return val
}

// GetNumberPropOrThrow retrieves a numeric property or panics if missing/invalid.
func GetNumberPropOrThrow[T NumberConstraint](props map[string]interface{}, prop string, message ...string) T {
	val, err := GetNumber[T](props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}
