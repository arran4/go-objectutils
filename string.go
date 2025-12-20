package go_objectutils

// GetStringPropOrDefault retrieves a string property or returns a default value.
func GetStringPropOrDefault(props map[string]interface{}, prop string, defaultValue string) string {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if strVal, ok := val.(string); ok {
		return strVal
	}
	return defaultValue
}

// GetStringPropOrDefaultFunction retrieves a string property or returns a value from a default function.
func GetStringPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() string) string {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	if strVal, ok := val.(string); ok {
		return strVal
	}
	return defaultFunction()
}

// GetStringPropOrThrow retrieves a string property or panics if missing/invalid.
func GetStringPropOrThrow(props map[string]interface{}, prop string, message ...string) string {
	msg := "Property " + prop + " is missing or not a string"
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
	if strVal, ok := val.(string); ok {
		return strVal
	}
	panic(msg)
}
