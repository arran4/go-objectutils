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

// GetStringArrayPtr retrieves a string array property as a pointer.
func GetStringArrayPtr(props map[string]interface{}, prop string) (*[]string, error) {
	val, err := GetStringArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetStringArrayPtr retrieves a string array property as a pointer or panics.
func MustGetStringArrayPtr(props map[string]interface{}, prop string) *[]string {
	val, err := GetStringArrayPtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringArrayPtrOrDefault retrieves a string array property as a pointer or returns a default value.
func GetStringArrayPtrOrDefault(props map[string]interface{}, prop string, defaultValue *[]string) *[]string {
	val, err := GetStringArrayPtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetStringPointerArray retrieves a property as a slice of string pointers.
func GetStringPointerArray(props map[string]interface{}, prop string) ([]*string, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]*string); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]*string, len(arr))
		for i, v := range arr {
			if v == nil {
				res[i] = nil
				continue
			}
			if s, ok := v.(string); ok {
				res[i] = &s
			} else if sp, ok := v.(*string); ok {
				res[i] = sp
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "string pointer element", Actual: v}
			}
		}
		return res, nil
	}
	// Also handle if the value is already []string
	if arr, ok := val.([]string); ok {
		res := make([]*string, len(arr))
		for i := range arr {
			res[i] = &arr[i]
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// GetStringPointerArrayPtr retrieves a property as a pointer to a slice of string pointers.
func GetStringPointerArrayPtr(props map[string]interface{}, prop string) (*[]*string, error) {
	val, err := GetStringPointerArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
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

// GetObjectArrayPtr retrieves an object array property as a pointer.
func GetObjectArrayPtr[T any](props map[string]interface{}, prop string) (*[]T, error) {
	val, err := GetObjectArray[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetObjectArrayPtr retrieves an object array property as a pointer or panics.
func MustGetObjectArrayPtr[T any](props map[string]interface{}, prop string) *[]T {
	val, err := GetObjectArrayPtr[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetObjectArrayPtrOrDefault retrieves an object array property as a pointer or returns a default value.
func GetObjectArrayPtrOrDefault[T any](props map[string]interface{}, prop string, defaultValue *[]T) *[]T {
	val, err := GetObjectArrayPtr[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetObjectPointerArray retrieves a property as a slice of object pointers.
func GetObjectPointerArray[T any](props map[string]interface{}, prop string) ([]*T, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]*T); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]*T, len(arr))
		for i, v := range arr {
			if v == nil {
				res[i] = nil
				continue
			}
			if castVal, ok := v.(T); ok {
				res[i] = &castVal
				continue
			}
			if castVal, ok := v.(*T); ok {
				res[i] = castVal
				continue
			}
			var zero T
			return nil, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("*%T element", zero), Actual: v}
		}
		return res, nil
	}
	// Also handle if the value is already []T
	if arr, ok := val.([]T); ok {
		res := make([]*T, len(arr))
		for i := range arr {
			res[i] = &arr[i]
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// GetObjectPointerArrayPtr retrieves a property as a pointer to a slice of object pointers.
func GetObjectPointerArrayPtr[T any](props map[string]interface{}, prop string) (*[]*T, error) {
	val, err := GetObjectPointerArray[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
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
	if arr, ok := val.([]time.Time); ok {
		return arr, nil
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

// GetDateArrayPtr retrieves a date array property as a pointer.
func GetDateArrayPtr(props map[string]interface{}, prop string) (*[]time.Time, error) {
	val, err := GetDateArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetDateArrayPtr retrieves a date array property as a pointer or panics.
func MustGetDateArrayPtr(props map[string]interface{}, prop string) *[]time.Time {
	val, err := GetDateArrayPtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetDateArrayPtrOrDefault retrieves a date array property as a pointer or returns a default value.
func GetDateArrayPtrOrDefault(props map[string]interface{}, prop string, defaultValue *[]time.Time) *[]time.Time {
	val, err := GetDateArrayPtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetDatePointerArray retrieves a property as a slice of date pointers.
func GetDatePointerArray(props map[string]interface{}, prop string) ([]*time.Time, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]*time.Time); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]*time.Time, len(arr))
		for i, v := range arr {
			if v == nil {
				res[i] = nil
				continue
			}
			if t, err := parseDate(v); err == nil {
				res[i] = &t
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "date pointer element", Actual: v, Cause: err}
			}
		}
		return res, nil
	}
	// Also handle if the value is already []time.Time
	if arr, ok := val.([]time.Time); ok {
		res := make([]*time.Time, len(arr))
		for i := range arr {
			res[i] = &arr[i]
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// GetDatePointerArrayPtr retrieves a property as a pointer to a slice of date pointers.
func GetDatePointerArrayPtr(props map[string]interface{}, prop string) (*[]*time.Time, error) {
	val, err := GetDatePointerArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// GetNumberArray retrieves a number array property.
func GetNumberArray[T NumberConstraint](props map[string]interface{}, prop string) ([]T, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]T, len(arr))
		for i, v := range arr {
			if num, err := convertToNumber[T](v); err == nil {
				res[i] = num
			} else {
				var zero T
				return nil, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("%T element", zero), Actual: v, Cause: err}
			}
		}
		return res, nil
	}
	// Also handle if the value is already []T (though unlikely from JSON unmarshal into map[string]interface{})
	if arr, ok := val.([]T); ok {
		return arr, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// MustGetNumberArray retrieves a number array property or panics.
func MustGetNumberArray[T NumberConstraint](props map[string]interface{}, prop string) []T {
	val, err := GetNumberArray[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetNumberArrayOrDefault retrieves a number array property or returns a default value.
func GetNumberArrayOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue []T) []T {
	val, err := GetNumberArray[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetNumberArrayPtr retrieves a number array property as a pointer.
func GetNumberArrayPtr[T NumberConstraint](props map[string]interface{}, prop string) (*[]T, error) {
	val, err := GetNumberArray[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetNumberArrayPtr retrieves a number array property as a pointer or panics.
func MustGetNumberArrayPtr[T NumberConstraint](props map[string]interface{}, prop string) *[]T {
	val, err := GetNumberArrayPtr[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetNumberArrayPtrOrDefault retrieves a number array property as a pointer or returns a default value.
func GetNumberArrayPtrOrDefault[T NumberConstraint](props map[string]interface{}, prop string, defaultValue *[]T) *[]T {
	val, err := GetNumberArrayPtr[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetNumberPointerArray retrieves a property as a slice of number pointers.
func GetNumberPointerArray[T NumberConstraint](props map[string]interface{}, prop string) ([]*T, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]*T); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]*T, len(arr))
		for i, v := range arr {
			if v == nil {
				res[i] = nil
				continue
			}
			if num, err := convertToNumber[T](v); err == nil {
				res[i] = &num
			} else {
				var zero T
				return nil, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("*%T element", zero), Actual: v, Cause: err}
			}
		}
		return res, nil
	}
	// Also handle if the value is already []T
	if arr, ok := val.([]T); ok {
		res := make([]*T, len(arr))
		for i := range arr {
			res[i] = &arr[i]
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// GetNumberPointerArrayPtr retrieves a property as a pointer to a slice of number pointers.
func GetNumberPointerArrayPtr[T NumberConstraint](props map[string]interface{}, prop string) (*[]*T, error) {
	val, err := GetNumberPointerArray[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// GetBooleanArray retrieves a boolean array property.
func GetBooleanArray(props map[string]interface{}, prop string) ([]bool, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]bool, len(arr))
		for i, v := range arr {
			if b, ok := v.(bool); ok {
				res[i] = b
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "bool element", Actual: v}
			}
		}
		return res, nil
	}
	if arr, ok := val.([]bool); ok {
		return arr, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// MustGetBooleanArray retrieves a boolean array property or panics.
func MustGetBooleanArray(props map[string]interface{}, prop string) []bool {
	val, err := GetBooleanArray(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetBooleanArrayOrDefault retrieves a boolean array property or returns a default value.
func GetBooleanArrayOrDefault(props map[string]interface{}, prop string, defaultValue []bool) []bool {
	val, err := GetBooleanArray(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetBooleanArrayPtr retrieves a boolean array property as a pointer.
func GetBooleanArrayPtr(props map[string]interface{}, prop string) (*[]bool, error) {
	val, err := GetBooleanArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetBooleanArrayPtr retrieves a boolean array property as a pointer or panics.
func MustGetBooleanArrayPtr(props map[string]interface{}, prop string) *[]bool {
	val, err := GetBooleanArrayPtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetBooleanArrayPtrOrDefault retrieves a boolean array property as a pointer or returns a default value.
func GetBooleanArrayPtrOrDefault(props map[string]interface{}, prop string, defaultValue *[]bool) *[]bool {
	val, err := GetBooleanArrayPtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetBooleanPointerArray retrieves a property as a slice of boolean pointers.
func GetBooleanPointerArray(props map[string]interface{}, prop string) ([]*bool, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if arr, ok := val.([]*bool); ok {
		return arr, nil
	}
	if arr, ok := val.([]interface{}); ok {
		res := make([]*bool, len(arr))
		for i, v := range arr {
			if v == nil {
				res[i] = nil
				continue
			}
			if b, ok := v.(bool); ok {
				res[i] = &b
			} else if bp, ok := v.(*bool); ok {
				res[i] = bp
			} else {
				return nil, &InvalidTypeError{Prop: prop, Expected: "bool pointer element", Actual: v}
			}
		}
		return res, nil
	}
	// Also handle if the value is already []bool
	if arr, ok := val.([]bool); ok {
		res := make([]*bool, len(arr))
		for i := range arr {
			res[i] = &arr[i]
		}
		return res, nil
	}
	return nil, &InvalidTypeError{Prop: prop, Expected: "array", Actual: val}
}

// GetBooleanPointerArrayPtr retrieves a property as a pointer to a slice of boolean pointers.
func GetBooleanPointerArrayPtr(props map[string]interface{}, prop string) (*[]*bool, error) {
	val, err := GetBooleanPointerArray(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
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
