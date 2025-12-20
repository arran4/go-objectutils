package go_objectutils

// GetString retrieves a string property.
// It returns an error if the property is missing or not a string.
func GetString(props map[string]interface{}, prop string) (string, error) {
	if props == nil {
		return "", &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return "", &MissingFieldError{Prop: prop}
	}
	if strVal, ok := val.(string); ok {
		return strVal, nil
	}
	return "", &InvalidTypeError{Prop: prop, Expected: "string", Actual: val}
}

// MustGetString retrieves a string property or panics.
func MustGetString(props map[string]interface{}, prop string) string {
	val, err := GetString(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringOrDefault retrieves a string property or returns a default value.
func GetStringOrDefault(props map[string]interface{}, prop string, defaultValue string) string {
	val, err := GetString(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetStringPtr retrieves a string property as a pointer.
func GetStringPtr(props map[string]interface{}, prop string) (*string, error) {
	val, err := GetString(props, prop)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetStringPtr retrieves a string property as a pointer or panics.
func MustGetStringPtr(props map[string]interface{}, prop string) *string {
	val, err := GetStringPtr(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringPtrOrDefault retrieves a string property as a pointer or returns a default value.
func GetStringPtrOrDefault(props map[string]interface{}, prop string, defaultValue *string) *string {
	val, err := GetStringPtr(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy aliases or extended functionality

// GetStringPropOrDefault is an alias for GetStringOrDefault
func GetStringPropOrDefault(props map[string]interface{}, prop string, defaultValue string) string {
	return GetStringOrDefault(props, prop, defaultValue)
}

// GetStringPropOrDefaultFunction retrieves a string property or returns a value from a default function.
func GetStringPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() string) string {
	val, err := GetString(props, prop)
	if err != nil {
		return defaultFunction()
	}
	return val
}

// GetStringPropOrThrow behaves like MustGetString but allows a custom message.
func GetStringPropOrThrow(props map[string]interface{}, prop string, message ...string) string {
	val, err := GetString(props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}
