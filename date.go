package go_objectutils

import (
	"time"
)

// parseDate tries to convert interface{} to time.Time
func parseDate(val interface{}) (time.Time, bool) {
	switch v := val.(type) {
	case time.Time:
		return v, true
	case string:
		// Try parsing ISO8601 / RFC3339
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			return t, true
		}
		// Try other formats if needed, e.g. JS Date.toString()
		// For now, stick to RFC3339 as standard JSON date format
	case float64:
		// Timestamp in ms (JS Date default) or seconds?
		// JS Date.now() returns ms.
		// Go time.UnixMilli takes int64.
		return time.UnixMilli(int64(v)), true
	case int64:
		return time.UnixMilli(v), true
	case int:
		return time.UnixMilli(int64(v)), true
	}
	return time.Time{}, false
}

// GetDatePropOrDefault retrieves a date property or returns a default value.
func GetDatePropOrDefault(props map[string]interface{}, prop string, defaultValue time.Time) time.Time {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if t, ok := parseDate(val); ok {
		return t
	}
	return defaultValue
}

// GetDatePropOrDefaultFunction
func GetDatePropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() time.Time) time.Time {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	if t, ok := parseDate(val); ok {
		return t
	}
	return defaultFunction()
}

// GetDatePropOrThrow
func GetDatePropOrThrow(props map[string]interface{}, prop string, message ...string) time.Time {
	msg := "Property " + prop + " is missing or not a date"
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
	if t, ok := parseDate(val); ok {
		return t
	}
	panic(msg)
}
