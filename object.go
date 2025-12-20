package go_objectutils

import "fmt"

// GetObject retrieves an object property (as T).
func GetObject[T any](props map[string]interface{}, prop string) (T, error) {
	var zero T
	if props == nil {
		return zero, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return zero, &MissingFieldError{Prop: prop}
	}
	if castVal, ok := val.(T); ok {
		return castVal, nil
	}
	return zero, &InvalidTypeError{Prop: prop, Expected: fmt.Sprintf("%T", zero), Actual: val}
}

// MustGetObject retrieves an object property or panics.
func MustGetObject[T any](props map[string]interface{}, prop string) T {
	val, err := GetObject[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetObjectOrDefault retrieves an object property or returns a default value.
func GetObjectOrDefault[T any](props map[string]interface{}, prop string, defaultValue T) T {
	val, err := GetObject[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetObjectPtr retrieves an object property as a pointer.
func GetObjectPtr[T any](props map[string]interface{}, prop string) (*T, error) {
	val, err := GetObject[T](props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetObjectPtr retrieves an object property as a pointer or panics.
func MustGetObjectPtr[T any](props map[string]interface{}, prop string) *T {
	val, err := GetObjectPtr[T](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetObjectPtrOrDefault retrieves an object property as a pointer or returns a default value.
func GetObjectPtrOrDefault[T any](props map[string]interface{}, prop string, defaultValue *T) *T {
	val, err := GetObjectPtr[T](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetMap retrieves a map property.
func GetMap[K comparable, V any](props map[string]interface{}, prop string) (map[K]V, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}
	if castVal, ok := val.(map[K]V); ok {
		return castVal, nil
	}
	// Special case for map[string]interface{}
	// Since we can't easily iterate and cast generic K, V without reflection.
	return nil, &InvalidTypeError{Prop: prop, Expected: "map", Actual: val}
}

// MustGetMap retrieves a map property or panics.
func MustGetMap[K comparable, V any](props map[string]interface{}, prop string) map[K]V {
	val, err := GetMap[K, V](props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// Legacy Aliases

// GetObjectPropOrDefault
func GetObjectPropOrDefault[T any](props map[string]interface{}, prop string, defaultValue T) T {
	return GetObjectOrDefault(props, prop, defaultValue)
}

// GetMapPropOrDefault
func GetMapPropOrDefault[K comparable, V any](props map[string]interface{}, prop string, defaultValue map[K]V) map[K]V {
	val, err := GetMap[K, V](props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetObjectPropOrThrow
func GetObjectPropOrThrow[T any](props map[string]interface{}, prop string, message ...string) T {
	val, err := GetObject[T](props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}

// GetObjectFunctionPropOrDefault uses a constructor function.
func GetObjectFunctionPropOrDefault[T any](props map[string]interface{}, prop string, constructorFunc func(map[string]interface{}) T, defaultValue T) T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if m, ok := val.(map[string]interface{}); ok {
		return constructorFunc(m)
	}
	return defaultValue
}

// GetObjectFunctionPropOrThrow
func GetObjectFunctionPropOrThrow[T any](props map[string]interface{}, prop string, constructorFunc func(map[string]interface{}) T, message ...string) T {
	msg := "Property " + prop + " is missing or not an object"
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
	if m, ok := val.(map[string]interface{}); ok {
		return constructorFunc(m)
	}
	panic(msg)
}

// GetObjectPropOrDefaultAllowNull
func GetObjectPropOrDefaultAllowNull[T any](props map[string]interface{}, prop string, defaultValue T) *T {
	if props == nil {
		return &defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return &defaultValue
	}
	if val == nil {
		return nil
	}
	if castVal, ok := val.(T); ok {
		return &castVal
	}
	return &defaultValue
}
