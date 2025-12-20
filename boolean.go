package go_objectutils

// GetBooleanPropOrDefault retrieves a boolean property or returns a default value.
func GetBooleanPropOrDefault(props map[string]interface{}, prop string, defaultValue bool) bool {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if boolVal, ok := val.(bool); ok {
		return boolVal
	}
	return defaultValue
}

// GetBooleanPropOrDefaultFunction retrieves a boolean property or returns a value from a default function.
func GetBooleanPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() bool) bool {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	if boolVal, ok := val.(bool); ok {
		return boolVal
	}
	return defaultFunction()
}

// GetBooleanPropOrThrow retrieves a boolean property or panics if missing/invalid.
func GetBooleanPropOrThrow(props map[string]interface{}, prop string, message ...string) bool {
	msg := "Property " + prop + " is missing or not a boolean"
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
	if boolVal, ok := val.(bool); ok {
		return boolVal
	}
	panic(msg)
}

// GetBooleanFunctionPropOrDefault constructs a boolean using a function or returns default.
// In TS: constructorFunc: (v: unknown) => boolean
func GetBooleanFunctionPropOrDefault(props map[string]interface{}, prop string, constructorFunc func(interface{}) bool, defaultValue bool) bool {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// The constructorFunc is responsible for validating/converting the value.
	// But in TS version, it seems if constructorFunc is provided, it is used on the value.
	// If the value is there, we pass it to constructorFunc.
	// Wait, does constructorFunc handle invalid types? The signature says (v: unknown) => boolean.
	// So we should just call it.
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
