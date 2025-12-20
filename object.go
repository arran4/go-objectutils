package go_objectutils

// GetObjectPropOrDefault retrieves a nested object (as a map) or returns a default value.
// In Go, since we can't easily return a generic struct without reflection or Unmarshal,
// we return map[string]interface{} by default if T is not specified or if T is map.
// However, to match the TS generics, we can make this generic.
func GetObjectPropOrDefault[T any](props map[string]interface{}, prop string, defaultValue T) T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// If T is map[string]interface{}, we can cast.
	// If val is map[string]interface{}, we can check.
	// We need runtime check.
	if castVal, ok := val.(T); ok {
		return castVal
	}
	// Also handle if val is map[string]interface{} but T is something else?
	// TS "GetObjectProp" usually just gets the object.
	// If the user wants to cast it to a struct, direct casting won't work if val is a map.
	// But without a constructor function, we can't automagically convert map to struct here easily.
	// So we assume the user asks for map[string]interface{}.
	return defaultValue
}

// GetMapPropOrDefault is an alias for GetObjectPropOrDefault specialized for maps?
// TS has GetMapPropOrDefault. In JS Map is different from Object.
// But in JSON, they are both objects.
// If the source is a JS Map, it's not JSON serializable directly usually.
// But if we are processing JSON, it's an object.
// We'll implement GetMapPropOrDefault to return map[string]interface{} or map[K]V?
// TS: GetMapPropOrDefault<K, V, R> returns Map<K,V> | R.
// In Go: map[K]V.
func GetMapPropOrDefault[K comparable, V any](props map[string]interface{}, prop string, defaultValue map[K]V) map[K]V {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	if castVal, ok := val.(map[K]V); ok {
		return castVal
	}
	// If V is interface{}, and K is string, it might match map[string]interface{}
	if m, ok := val.(map[string]interface{}); ok {
		// If we want map[string]interface{}, fine.
		// If we want map[string]int, we might need conversion?
		// Given TS "Type Safety", we should probably only return if it matches or is convertible.
		// For now, let's try to simple cast.
		// We can't cast map[string]interface{} to map[K]V directly in Go.
		// We would need to iterate and copy.
		// Since we don't know K and V types at runtime easily without reflection (except K=string),
		// we'll stick to direct cast check first.
		_ = m
	}
	return defaultValue
}

// GetObjectPropOrThrow
func GetObjectPropOrThrow[T any](props map[string]interface{}, prop string, message ...string) T {
	msg := "Property " + prop + " is missing or not the expected type"
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
	if castVal, ok := val.(T); ok {
		return castVal
	}
	panic(msg)
}

// GetObjectFunctionPropOrDefault uses a constructor function.
// constructorFunc: (params: object) => Y
// In Go: func(map[string]interface{}) T
func GetObjectFunctionPropOrDefault[T any](props map[string]interface{}, prop string, constructorFunc func(map[string]interface{}) T, defaultValue T) T {
	if props == nil {
		return defaultValue
	}
	val, ok := props[prop]
	if !ok {
		return defaultValue
	}
	// val should be a map (object)
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
		return nil // Explicit null
	}
	if castVal, ok := val.(T); ok {
		return &castVal
	}
	return &defaultValue
}

// Note: In TS "AllowNull" means return type includes null.
// In Go, usually implies returning pointer or interface.
// If T is a pointer type, we can return nil.
// If T is a struct, we return *T.
// But generics `T any` doesn't guarantee we can return nil unless we use *T or interface.
// The TS signature: GetObjectPropOrDefaultAllowNull<Y>(...): Y.
// Wait, if Y includes null (Y | null), then it returns Y.
// If Y doesn't include null, but "AllowNull" is in name...
// The TS signature: `GetObjectPropOrDefaultAllowNull<Y>(..., defaultValue: Y): Y`.
// Usually defaultValue handles the null case?
// No, `AllowNull` usually means if the property is explicitly null in JSON, we might want to return null (if valid) or handle it.
// Looking at TS definitions:
// `GetObjectPropOrDefaultAllowNull<Y>(..., defaultValue: Y): Y`
// This looks exactly the same as GetObjectPropOrDefault.
// Maybe it accepts `null` as a valid value for the property?
// In `GetObjectPropOrDefault`, if `props[prop]` is `null`, does it return default?
// TS `unknown` includes null.
// If `val` is null, `ok` is true in Go map lookup?
// In Go, if key exists and value is nil (interface{}(nil)), ok is true.
// `val.(T)` where T is e.g. `map[string]interface{}` will fail if val is nil.
// So `GetObjectPropOrDefault` returns defaultValue if val is nil.
// `GetObjectPropOrDefaultAllowNull` might imply that `null` is a valid value to return (if T allows it).
// But `defaultValue` is type Y.
// If we want to return nil, T must be a pointer or interface.
// For now, I'll skip implementing every single variation if they are redundant in Go, or implement them similarly.
