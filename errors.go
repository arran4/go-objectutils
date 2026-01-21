package go_objectutils

import "fmt"

// MissingFieldError indicates that a required field is missing from the map.
type MissingFieldError struct {
	Prop string
}

func (e *MissingFieldError) Error() string {
	return fmt.Sprintf("property '%s' is missing", e.Prop)
}

// InvalidTypeError indicates that a field exists but is not of the expected type.
type InvalidTypeError struct {
	Prop     string
	Expected string
	Actual   interface{}
	Cause    error
}

func (e *InvalidTypeError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("property '%s' is not of type %s, got %T: %v", e.Prop, e.Expected, e.Actual, e.Cause)
	}
	return fmt.Sprintf("property '%s' is not of type %s, got %T", e.Prop, e.Expected, e.Actual)
}

func (e *InvalidTypeError) Unwrap() error {
	return e.Cause
}

// RegexMismatchError indicates that a string field does not match the expected regular expression.
type RegexMismatchError struct {
	Prop       string
	Value      string
	Expression string
}

func (e *RegexMismatchError) Error() string {
	return fmt.Sprintf("property '%s' value '%s' does not match regex '%s'", e.Prop, e.Value, e.Expression)
}
