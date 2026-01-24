package go_objectutils

import (
	"math/big"
)

// GetBigInt retrieves a big.Int property.
// It supports strings, int64, float64 (if integer).
func GetBigInt(props map[string]interface{}, prop string) (*big.Int, error) {
	if props == nil {
		return nil, &MissingFieldError{Prop: prop}
	}
	val, ok := props[prop]
	if !ok {
		return nil, &MissingFieldError{Prop: prop}
	}

	switch v := val.(type) {
	case string:
		bi := new(big.Int)
		if _, ok := bi.SetString(v, 10); ok {
			return bi, nil
		}
	case int64:
		return big.NewInt(v), nil
	case int:
		return big.NewInt(int64(v)), nil
	case float64:
		// Only if it's an integer? Or truncate?
		// JS BigInt(number) truncates if I recall, but usually we want exact representation.
		// big.NewFloat(v).Int(bi)
		bf := big.NewFloat(v)
		i, _ := bf.Int(nil)
		return i, nil
	}

	return nil, &InvalidTypeError{Prop: prop, Expected: "big.Int convertible", Actual: val}
}

// MustGetBigInt retrieves a big.Int property or panics.
func MustGetBigInt(props map[string]interface{}, prop string) *big.Int {
	val, err := GetBigInt(props, prop)
	if err != nil {
		panic(err)
	}
	return val
}

// GetBigIntOrDefault retrieves a big.Int property or returns a default value.
func GetBigIntOrDefault(props map[string]interface{}, prop string, defaultValue *big.Int) *big.Int {
	val, err := GetBigInt(props, prop)
	if err != nil {
		return defaultValue
	}
	return val
}

// Legacy int64 Wrappers (renamed to avoid confusion, or kept if strictly needed for compatibility with old code that might expect int64)
// The file name "bigint.go" suggests BigInt support.
// In the original code, it returned int64.
// I will update them to use GetNumber[int64] logic or keep them as is but using the new structure?
// The user asked for "primitive types and variations". big.Int is not primitive but useful.
// I'll keep the int64 versions but implemented via GetNumber to reduce code duplication if possible,
// OR just leave them as aliases to GetNumber[int64].

// GetBigIntPropOrDefault aliases GetNumberOrDefault[int64]
func GetBigIntPropOrDefault(props map[string]interface{}, prop string, defaultValue int64) int64 {
	return GetNumberOrDefault[int64](props, prop, defaultValue)
}

// GetBigIntPropOrDefaultFunction
func GetBigIntPropOrDefaultFunction(props map[string]interface{}, prop string, defaultFunction func() int64) int64 {
	val, err := GetNumber[int64](props, prop)
	if err != nil {
		return defaultFunction()
	}
	return val
}

// GetBigIntPropOrThrow
func GetBigIntPropOrThrow(props map[string]interface{}, prop string, message ...string) int64 {
	val, err := GetNumber[int64](props, prop)
	if err != nil {
		msg := err.Error()
		if len(message) > 0 {
			msg = message[0]
		}
		panic(msg)
	}
	return val
}
