package go_objectutils

// GetBigIntPropOrDefault retrieves a bigint (int64) property or returns a default value.
// In Go, we use int64 as the closest mapping for JS bigint in typical usage.
func GetBigIntPropOrDefault(props map[string]interface{}, prop string, defaultValue int64) int64 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// Handle different numeric types and string
	if i, ok := convertToNumber[int64](val); ok {
		return i
	}
	return defaultValue
}

// GetBigIntPropOrDefaultFunction
func GetBigIntPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() int64) int64 {
	if props == nil {
		return defaultFunction()
	}
	val, ok := props[prop]
	if !ok {
		return defaultFunction()
	}
	if i, ok := convertToNumber[int64](val); ok {
		return i
	}
	return defaultFunction()
}

// GetBigIntPropOrThrow
func GetBigIntPropOrThrow(props map[string]interface{}, prop string, message ...string) int64 {
	msg := "Property " + prop + " is missing or not a bigint"
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
	if i, ok := convertToNumber[int64](val); ok {
		return i
	}
	panic(msg)
}
