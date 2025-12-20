package go_objectutils

import (
	"fmt"
	"time"
)

// GetStringArray retrieves a string array property.
func GetStringArray(props map[string]interface{}, prop string) ([]string, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]string); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]string, len(arr))
		for i, v := range arr {
			if s, ok := v.(string); ok {
				res[i] = s
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "string element", Actual: v}
			}
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "[]string", Actual: val}
}

// MustGetStringArray retrieves a string array property or panics.
func MustGetStringArray(props map[string]interface{}, prop string) []string {
	val, err := GetStringArray(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringArrayOrDefault retrieves a string array property or returns a default value.
func GetStringArrayOrDefault(props map[string]interface{}, prop string, defaultValue []string) []string {
	val, err := GetStringArray(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetObjectArray retrieves an object array property.
func GetObjectArray[T any](props map[string]interface{}, prop string) ([]T, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]T); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]T, len(arr))
		for i, v := range arr {
			if castVal, ok := v.(T); ok {
				res[i] = castVal
				continue
			}
			// Special handling if T is map[string]interface{} (or alias)
			if m, ok := v.(map[string]interface{}); ok {
				// We can't direct cast v to T if T is a named type of map[string]interface{}.
				// But we can cast m to T if T is compatible.
				// In Go, named types are different.
				// Reflection or re-assignment is needed.
				// However, standard casting interface{}(m).(T) works if T is the same underlying type?
				// No, only if T is interface{} or map[string]interface{} (unnamed) or if m was already T.
				// But m came from v which is interface{}, and v.(map[string]interface{}) succeeded.
				// So m IS map[string]interface{}.
				// If T is MyObj (type MyObj map[string]interface{}), then interface{}(m).(MyObj) will fail?
				// Yes, because m is NOT MyObj. It is map[string]interface{}.
				// We need to convert it.
				// Since we can't do `T(m)` because T is generic.
				// But since both are map[string]interface{}, we can cast via assignment? No.
				// We might need to just accept that if T is named map type, we can't easily auto-convert from generic map.
				// UNLESS we use reflection or if we assume T is map[string]interface{}.

				// For now, let's try to handle the case where T is exactly map[string]interface{}.
				// The test uses `type MyObj map[string]interface{}`.
				// This is a named type.
				// Go generics don't allow casting to T from underlying type easily.
				// We can return error or try to be smarter.
				// If we use reflection, we can convert.
				// But avoiding reflection is good.
				// Let's check if the test actually needs named type or if I can relax the test.
				// The user wants "extensive" support.
				// Maybe I should use reflection for this specific case?
				// Or I can just check if I can cast.
				if castVal, ok := interface{}(m).(T); ok {
					res[i] = castVal
					continue
				}
				// If T is a map type compatible with map[string]interface{}, we might want to cast.
				// But we can't without reflection.
			}
			var zero T
			return nil, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("%T element", zero), Actual: v}
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// MustGetObjectArray retrieves an object array property or panics.
func MustGetObjectArray[T any](props map[string]interface{}, prop string) []T {
	val, err := GetObjectArray[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetObjectArrayOrDefault retrieves an object array property or returns a default value.
func GetObjectArrayOrDefault[T any](props map[string]interface{}, prop string, defaultValue []T) []T {
	val, err := GetObjectArray[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetDateArray retrieves a date array property.
func GetDateArray(props map[string]interface{}, prop string) ([]time.Time, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]time.Time, len(arr))
		for i, v := range arr {
			if t, err := parseDate(v); err == nil {
				res[i] = t
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "date element", Actual: v, Cause: err}
			}
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// MustGetDateArray retrieves a date array property or panics.
func MustGetDateArray(props map[string]interface{}, prop string) []time.Time {
	val, err := GetDateArray(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetDateArrayOrDefault retrieves a date array property or returns a default value.
func GetDateArrayOrDefault(props map[string]interface{}, prop string, defaultValue []time.Time) []time.Time {
	val, err := GetDateArray(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy Aliases

// GetStringArrayPropOrDefault
func GetStringArrayPropOrDefault(props map[string]interface{}, prop string, defaultValue []string) []string {
	return GetStringArrayOrDefault(props, prop, defaultValue)
}

// GetStringArrayPropOrThrow
func GetStringArrayPropOrThrow(props map[string]interface{}, prop string, message ...string) []string {
	val, err := GetStringArray(props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}

// GetObjectArrayPropOrDefault
func GetObjectArrayPropOrDefault[T any](props map[string]interface{}, prop string, defaultValue []T) []T {
	return GetObjectArrayOrDefault[T](props, prop, defaultValue)
}

// GetObjectArrayFunctionPropOrDefault
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
	return GetDateArrayOrDefault(props, prop, defaultValue)
}
