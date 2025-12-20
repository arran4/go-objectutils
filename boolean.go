package go_objectutils

// GetBoolean retrieves a boolean property.
func GetBoolean(props map[string]interface{}, prop string) (bool, error) {
	if props == nil {
		return false, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return false, &MissingFieldError{Prop: prop}
	}
	if boolVal, ok := val.(bool); ok {
		return boolVal, nil
	}
	return false, &InvalidTypeError{Prop: prop, Expected: "bool", Actual: val}
}

// MustGetBoolean retrieves a boolean property or panics.
func MustGetBoolean(props map[string]interface{}, prop string) bool {
	val, err := GetBoolean(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetBooleanOrDefault retrieves a boolean property or returns a default value.
func GetBooleanOrDefault(props map[string]interface{}, prop string, defaultValue bool) bool {
	val, err := GetBoolean(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetBooleanPtr retrieves a boolean property as a pointer.
func GetBooleanPtr(props map[string]interface{}, prop string) (*bool, error) {
	val, err := GetBoolean(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetBooleanPtr retrieves a boolean property as a pointer or panics.
func MustGetBooleanPtr(props map[string]interface{}, prop string) *bool {
	val, err := GetBooleanPtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetBooleanPtrOrDefault retrieves a boolean property as a pointer or returns a default value.
func GetBooleanPtrOrDefault(props map[string]interface{}, prop string, defaultValue *bool) *bool {
	val, err := GetBooleanPtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy Aliases

// GetBooleanPropOrDefault is an alias for GetBooleanOrDefault.
func GetBooleanPropOrDefault(props map[string]interface{}, prop string, defaultValue bool) bool {
	return GetBooleanOrDefault(props, prop, defaultValue)
}

// GetBooleanPropOrDefaultFunction retrieves a boolean property or returns a value from a default function.
func GetBooleanPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() bool) bool {
	val, err := GetBoolean(props, prop)
	if err != nil {
		return defaultFunction()
	}
	return val
}

// GetBooleanPropOrThrow retrieves a boolean property or panics if missing/invalid.
func GetBooleanPropOrThrow(props map[string]interface{}, prop string, message ...string) bool {
	val, err := GetBoolean(props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}

// GetBooleanFunctionPropOrDefault constructs a boolean using a function or returns default.
func GetBooleanFunctionPropOrDefault(props map[string]interface{}, prop string, constructorFunc func(interface{}) bool, defaultValue bool) bool {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	return constructorFunc(val)
}

// GetBooleanFunctionPropOrDefaultFunction
func GetBooleanFunctionPropOrDefaultFunction(props map[string]interface{}, prop string, constructorFunc func(interface{}) bool, defaultFunction func() bool) bool {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	return constructorFunc(val)
}
