package go_objectutils

import (
	"fmt"
	"time"
)

// GetDate retrieves a date property.
func GetDate(props map[string]interface{}, prop string) (time.Time, error) {
	if props == nil {
		return time.Time{}, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return time.Time{}, &MissingFieldError{Prop: prop}
	}
	if t, err := parseDate(val); err == nil {
		return t, nil
	} else {
		return time.Time{}, &InvalidTypeError{Prop: prop, Expected: "time.Time", Actual: val, Cause: err}
	}
}

// MustGetDate retrieves a date property or panics.
func MustGetDate(props map[string]interface{}, prop string) time.Time {
	val, err := GetDate(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetDateOrDefault retrieves a date property or returns a default value.
func GetDateOrDefault(props map[string]interface{}, prop string, defaultValue time.Time) time.Time {
	val, err := GetDate(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetDatePtr retrieves a date property as a pointer.
func GetDatePtr(props map[string]interface{}, prop string) (*time.Time, error) {
	val, err := GetDate(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetDatePtr retrieves a date property as a pointer or panics.
func MustGetDatePtr(props map[string]interface{}, prop string) *time.Time {
	val, err := GetDatePtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetDatePtrOrDefault retrieves a date property as a pointer or returns a default value.
func GetDatePtrOrDefault(props map[string]interface{}, prop string, defaultValue *time.Time) *time.Time {
	val, err := GetDatePtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy Aliases

// GetDatePropOrDefault is an alias for GetDateOrDefault.
func GetDatePropOrDefault(props map[string]interface{}, prop string, defaultValue time.Time) time.Time {
	return GetDateOrDefault(props, prop, defaultValue)
}

// GetDatePropOrDefaultFunction
func GetDatePropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() time.Time) time.Time {
	val, err := GetDate(props, prop)
	if err != nil {
		return defaultFunction()
	}
	return val
}

// GetDatePropOrThrow
func GetDatePropOrThrow(props map[string]interface{}, prop string, message ...string) time.Time {
	val, err := GetDate(props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}

// parseDate tries to convert interface{} to time.Time
func parseDate(val interface{}) (time.Time, error) {
	switch v := val.(type) {
	case time.Time:
		return v, nil
	case string:
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			return t, nil
		} else {
			return time.Time{}, err
		}
	case float64:
		return time.UnixMilli(int64(v)), nil
	case int64:
		return time.UnixMilli(v), nil
	case int:
		return time.UnixMilli(int64(v)), nil
	}
	return time.Time{}, fmt.Errorf("cannot parse %T as date", val)
}
