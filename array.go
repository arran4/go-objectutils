package go_objectutils

import (
	"time"
)

// GetStringArrayPropOrDefault
func GetStringArrayPropOrDefault(props map[string]interface{}, prop string, defaultValue []string) []string {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// val should be []interface{} or []string
	if arr, ok := val.([]string); ok {
		return arr
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]string, len(arr))
		for i, v := range arr {
			if s, ok := v.(string); ok {
				res[i] = s
			} else {
				// If strictly string array, maybe fail or skip?
				// TS "Type safety" implies checks.
				// If one element is not string, is the whole array invalid?
				// Usually yes.
				return defaultValue
			}
		}
		return res
	}
	return defaultValue
}

// GetStringArrayPropOrThrow
func GetStringArrayPropOrThrow(props map[string]interface{}, prop string, message ...string) []string {
	msg := "Property " + prop + " is missing or not a string array"
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
	if arr, ok := val.([]string); ok {
		return arr
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]string, len(arr))
		for i, v := range arr {
			if s, ok := v.(string); ok {
				res[i] = s
			} else {
				panic(msg)
			}
		}
		return res
	}
	panic(msg)
}

// GetObjectArrayPropOrDefault
// Y is the item type. X is the array type (Y[]).
// In Go: []T
func GetObjectArrayPropOrDefault[T any](props map[string]interface{}, prop string, defaultValue []T) []T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// Check if it is []T
	if arr, ok := val.([]T); ok {
		return arr
	}
	// Check if it is []interface{} and cast items?
	// If T is map[string]interface{}, and val is []interface{} (where items are maps).
	if arr, ok := val.([]interface{}); ok {
		res := make([]T, len(arr))
		for i, v := range arr {
			if castVal, ok := v.(T); ok {
				res[i] = castVal
			} else {
				return defaultValue
			}
		}
		return res
	}
	return defaultValue
}

// GetObjectArrayFunctionPropOrDefault
// constructorFunc: (params: object) => Y
func GetObjectArrayFunctionPropOrDefault[T any](props map[string]interface{}, prop string, constructorFunc func(map[string]interface{}) T, defaultValue []T) []T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]T, len(arr))
		for i, v := range arr {
			if m, ok := v.(map[string]interface{}); ok {
				res[i] = constructorFunc(m)
			} else {
				return defaultValue
			}
		}
		return res
	}
	return defaultValue
}

// GetDateArrayPropOrDefault
func GetDateArrayPropOrDefault(props map[string]interface{}, prop string, defaultValue []time.Time) []time.Time {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]time.Time, len(arr))
		for i, v := range arr {
			if t, ok := parseDate(v); ok {
				res[i] = t
			} else {
				return defaultValue
			}
		}
		return res
	}
	return defaultValue
}
