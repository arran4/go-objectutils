package go_objectutils

import (
	"fmt"
	"regexp"
)

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

// GetStringRegex retrieves a string property and validates it against a regular expression.
// It returns an error if the property is missing, not a string, the regex is invalid, or the value doesn't match.
func GetStringRegex(props map[string]interface{}, prop string, expression string) (string, error) {
	val, err := GetString(props, prop)
	if err != nil {
		return "", err
	}
	matched, err := regexp.MatchString(expression, val)
	if err != nil {
		return "", fmt.Errorf("invalid regex '%s': %w", expression, err)
	}
	if !matched {
		return "", &RegexMismatchError{Prop: prop, Value: val, Expression: expression}
	}
	return val, nil
}

// MustGetStringRegex retrieves a string property validated against a regex or panics.
func MustGetStringRegex(props map[string]interface{}, prop string, expression string) string {
	val, err := GetStringRegex(props, prop, expression)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringRegexOrDefault retrieves a string property validated against a regex or returns a default value.
func GetStringRegexOrDefault(props map[string]interface{}, prop string, expression string, defaultValue string) string {
	val, err := GetStringRegex(props, prop, expression)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetStringRegexPtr retrieves a string property as a pointer and validates it against a regular expression.
func GetStringRegexPtr(props map[string]interface{}, prop string, expression string) (*string, error) {
	val, err := GetStringRegex(props, prop, expression)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

// MustGetStringRegexPtr retrieves a string property as a pointer validated against a regex or panics.
func MustGetStringRegexPtr(props map[string]interface{}, prop string, expression string) *string {
	val, err := GetStringRegexPtr(props, prop, expression)
	if err != nil {
		panic(err)
	}
	return val
}

// GetStringRegexPtrOrDefault retrieves a string property as a pointer validated against a regex or returns a default value.
func GetStringRegexPtrOrDefault(props map[string]interface{}, prop string, expression string, defaultValue *string) *string {
	val, err := GetStringRegexPtr(props, prop, expression)
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
