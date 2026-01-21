package go_objectutils

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	props := map[string]interface{}{
		"valid":   "test",
		"invalid": 123,
	}

	// GetString
	val, err := GetString(props, "valid")
	assert.NoError(t, err)
	assert.Equal(t, "test", val)

	_, err = GetString(props, "missing")
	assert.Error(t, err)
	assert.IsType(t, &MissingFieldError{}, err)

	_, err = GetString(props, "invalid")
	assert.Error(t, err)
	assert.IsType(t, &InvalidTypeError{}, err)

	// MustGetString
	assert.Equal(t, "test", MustGetString(props, "valid"))
	assert.Panics(t, func() { MustGetString(props, "missing") })
	assert.Panics(t, func() { MustGetString(props, "invalid") })

	// GetStringOrDefault
	assert.Equal(t, "test", GetStringOrDefault(props, "valid", "default"))
	assert.Equal(t, "default", GetStringOrDefault(props, "missing", "default"))
	assert.Equal(t, "default", GetStringOrDefault(props, "invalid", "default"))

	// Pointer variations
	ptrVal, err := GetStringPtr(props, "valid")
	assert.NoError(t, err)
	assert.Equal(t, "test", *ptrVal)

	ptrVal = GetStringPtrOrDefault(props, "missing", nil)
	assert.Nil(t, ptrVal)
}

func TestStringRegex(t *testing.T) {
	props := map[string]interface{}{
		"email":   "test@example.com",
		"invalid": "not-an-email",
		"badType": 123,
	}
	emailRegex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

	// GetStringRegex
	val, err := GetStringRegex(props, "email", emailRegex)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", val)

	_, err = GetStringRegex(props, "invalid", emailRegex)
	assert.Error(t, err)
	assert.IsType(t, &RegexMismatchError{}, err)

	_, err = GetStringRegex(props, "badType", emailRegex)
	assert.Error(t, err)
	assert.IsType(t, &InvalidTypeError{}, err)

	_, err = GetStringRegex(props, "missing", emailRegex)
	assert.Error(t, err)
	assert.IsType(t, &MissingFieldError{}, err)

	// Invalid regex
	_, err = GetStringRegex(props, "email", "[")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid regex")

	// MustGetStringRegex
	assert.Equal(t, "test@example.com", MustGetStringRegex(props, "email", emailRegex))
	assert.Panics(t, func() { MustGetStringRegex(props, "invalid", emailRegex) })

	// GetStringRegexOrDefault
	assert.Equal(t, "test@example.com", GetStringRegexOrDefault(props, "email", emailRegex, "def"))
	assert.Equal(t, "def", GetStringRegexOrDefault(props, "invalid", emailRegex, "def"))

	// Pointer variations
	ptrVal, err := GetStringRegexPtr(props, "email", emailRegex)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", *ptrVal)

	ptrVal = GetStringRegexPtrOrDefault(props, "invalid", emailRegex, nil)
	assert.Nil(t, ptrVal)

	assert.Equal(t, "test@example.com", *MustGetStringRegexPtr(props, "email", emailRegex))
	assert.Panics(t, func() { MustGetStringRegexPtr(props, "invalid", emailRegex) })
}

func TestNumber(t *testing.T) {
	props := map[string]interface{}{
		"int":     10,
		"float":   10.5,
		"string":  "20",
		"invalid": "abc",
	}

	// GetNumber
	i, err := GetNumber[int](props, "int")
	assert.NoError(t, err)
	assert.Equal(t, 10, i)

	f, err := GetNumber[float64](props, "float")
	assert.NoError(t, err)
	assert.Equal(t, 10.5, f)

	// String conversion
	i2, err := GetNumber[int](props, "string")
	assert.NoError(t, err)
	assert.Equal(t, 20, i2)

	// MustGetNumber
	assert.Equal(t, 10, MustGetNumber[int](props, "int"))
	assert.Panics(t, func() { MustGetNumber[int](props, "missing") })

	// GetNumberOrDefault
	assert.Equal(t, 100, GetNumberOrDefault(props, "missing", 100))

	// Pointers
	iPtr, err := GetNumberPtr[int](props, "int")
	assert.NoError(t, err)
	assert.Equal(t, 10, *iPtr)
}

func TestBoolean(t *testing.T) {
	props := map[string]interface{}{
		"true":    true,
		"false":   false,
		"invalid": "true",
	}

	// GetBoolean
	b, err := GetBoolean(props, "true")
	assert.NoError(t, err)
	assert.True(t, b)

	b, err = GetBoolean(props, "false")
	assert.NoError(t, err)
	assert.False(t, b)

	_, err = GetBoolean(props, "invalid")
	assert.Error(t, err)

	// MustGetBoolean
	assert.True(t, MustGetBoolean(props, "true"))
	assert.Panics(t, func() { MustGetBoolean(props, "invalid") })

	// GetBooleanOrDefault
	assert.True(t, GetBooleanOrDefault(props, "missing", true))
}

func TestDate(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond) // Truncate because float conversion might lose precision if we are not careful, but millis is safeish
	props := map[string]interface{}{
		"time":      now,
		"string":    now.Format(time.RFC3339),
		"timestamp": now.UnixMilli(),
		"invalid":   "not a date",
	}

	// GetDate
	d, err := GetDate(props, "time")
	assert.NoError(t, err)
	assert.True(t, d.Equal(now))

	d, err = GetDate(props, "string")
	assert.NoError(t, err)
	assert.True(t, d.Equal(now.Truncate(time.Second)), "RFC3339 usually seconds precision unless fractional specified")

	d, err = GetDate(props, "timestamp")
	assert.NoError(t, err)
	// UnixMilli loses sub-millisecond precision
	assert.Equal(t, now.UnixMilli(), d.UnixMilli())

	// MustGetDate
	assert.NotPanics(t, func() { MustGetDate(props, "time") })
	assert.Panics(t, func() { MustGetDate(props, "invalid") })
}

func TestArray(t *testing.T) {
	props := map[string]interface{}{
		"strings": []string{"a", "b"},
		"ints":    []int{1, 2},
		"mixed":   []interface{}{"a", "b"},
		"objects": []interface{}{
			map[string]interface{}{"name": "obj1"},
			map[string]interface{}{"name": "obj2"},
		},
	}

	// GetStringArray
	strs, err := GetStringArray(props, "strings")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, strs)

	// Mixed interface array that contains strings
	strs, err = GetStringArray(props, "mixed")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, strs)

	// MustGetStringArray
	assert.Equal(t, []string{"a", "b"}, MustGetStringArray(props, "strings"))

	// GetObjectArray
	// Note: We use map[string]interface{} because generic casting from map[string]interface{} to a named type
	// is not directly possible without reflection in Go if the underlying values are not already of that named type.
	// Since JSON unmarshalling (or manual map creation) usually produces map[string]interface{},
	// we stick to that for this test unless we implement reflection-based conversion.
	objs, err := GetObjectArray[map[string]interface{}](props, "objects")
	assert.NoError(t, err)
	assert.Len(t, objs, 2)
}

func TestObject(t *testing.T) {
	sub := map[string]interface{}{"subKey": "subVal"}
	props := map[string]interface{}{
		"obj": sub,
	}

	// GetObject
	o, err := GetObject[map[string]interface{}](props, "obj")
	assert.NoError(t, err)
	assert.Equal(t, sub, o)

	// MustGetObject
	assert.Equal(t, sub, MustGetObject[map[string]interface{}](props, "obj"))

	// Ptr
	oPtr, err := GetObjectPtr[map[string]interface{}](props, "obj")
	assert.NoError(t, err)
	assert.Equal(t, sub, *oPtr)
}

func TestLegacy(t *testing.T) {
	props := map[string]interface{}{
		"key": "value",
	}

	assert.Equal(t, "value", GetStringPropOrDefault(props, "key", "def"))
	assert.Panics(t, func() { GetStringPropOrThrow(props, "missing") })
}

func TestErrorMessages(t *testing.T) {
	props := map[string]interface{}{
		"wrong": 123,
	}
	_, err := GetString(props, "missing")
	assert.Equal(t, "property 'missing' is missing", err.Error())

	_, err = GetString(props, "wrong")
	assert.Contains(t, err.Error(), "property 'wrong' is not of type string")
}

func TestBigInt(t *testing.T) {
	props := map[string]interface{}{
		"str":     "12345678901234567890",
		"int64":   int64(123),
		"int":     123,
		"float":   123.0,
		"invalid": "abc",
	}

	// GetBigInt
	bi, err := GetBigInt(props, "str")
	assert.NoError(t, err)
	expected, _ := new(big.Int).SetString("12345678901234567890", 10)
	assert.Equal(t, expected, bi)

	bi, err = GetBigInt(props, "int64")
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(123), bi)

	bi, err = GetBigInt(props, "int")
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(123), bi)

	bi, err = GetBigInt(props, "float")
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(123), bi)

	_, err = GetBigInt(props, "invalid")
	assert.Error(t, err)

	// MustGetBigInt
	assert.Equal(t, expected, MustGetBigInt(props, "str"))
	assert.Panics(t, func() { MustGetBigInt(props, "missing") })
	assert.Panics(t, func() { MustGetBigInt(props, "invalid") })

	// GetBigIntOrDefault
	def := big.NewInt(0)
	assert.Equal(t, expected, GetBigIntOrDefault(props, "str", def))
	assert.Equal(t, def, GetBigIntOrDefault(props, "missing", def))
}

func TestLegacyBigInt(t *testing.T) {
	props := map[string]interface{}{
		"val": 100,
	}
	assert.Equal(t, int64(100), GetBigIntPropOrDefault(props, "val", 0))
	assert.Equal(t, int64(0), GetBigIntPropOrDefault(props, "missing", 0))
	assert.Equal(t, int64(100), GetBigIntPropOrThrow(props, "val"))
	assert.Equal(t, int64(100), GetBigIntPropOrThrow(props, "val", "custom msg"))
	assert.Panics(t, func() { GetBigIntPropOrThrow(props, "missing") })
	assert.Panics(t, func() { GetBigIntPropOrThrow(props, "missing", "msg") })

	assert.Equal(t, int64(100), GetBigIntPropOrDefaultFunction(props, "val", func() int64 { return 1 }))
	assert.Equal(t, int64(1), GetBigIntPropOrDefaultFunction(props, "missing", func() int64 { return 1 }))
}

func TestPointers(t *testing.T) {
	props := map[string]interface{}{
		"str":  "s",
		"num":  1,
		"bool": true,
		"date": time.Now(),
		"obj":  map[string]interface{}{},
	}

	// String
	assert.Equal(t, "s", *MustGetStringPtr(props, "str"))
	assert.Nil(t, GetStringPtrOrDefault(props, "missing", nil))
	assert.Panics(t, func() { MustGetStringPtr(props, "missing") })

	// Number
	assert.Equal(t, 1, *MustGetNumberPtr[int](props, "num"))
	assert.Nil(t, GetNumberPtrOrDefault[int](props, "missing", nil))
	assert.Panics(t, func() { MustGetNumberPtr[int](props, "missing") })

	// Boolean
	assert.Equal(t, true, *MustGetBooleanPtr(props, "bool"))
	assert.Nil(t, GetBooleanPtrOrDefault(props, "missing", nil))
	assert.Panics(t, func() { MustGetBooleanPtr(props, "missing") })

	// Date
	assert.NotNil(t, MustGetDatePtr(props, "date"))
	assert.Nil(t, GetDatePtrOrDefault(props, "missing", nil))
	assert.Panics(t, func() { MustGetDatePtr(props, "missing") })

	// Object
	assert.NotNil(t, MustGetObjectPtr[map[string]interface{}](props, "obj"))
	assert.Nil(t, GetObjectPtrOrDefault[map[string]interface{}](props, "missing", nil))
	assert.Panics(t, func() { MustGetObjectPtr[map[string]interface{}](props, "missing") })
}

func TestMap(t *testing.T) {
	sub := map[string]interface{}{"k": "v"}
	props := map[string]interface{}{
		"map": sub,
	}

	// GetMap
	m, err := GetMap[string, interface{}](props, "map")
	assert.NoError(t, err)
	assert.Equal(t, sub, m)

	_, err = GetMap[string, int](props, "map")
	// Since casting map[string]interface{} to map[string]int fails without reflection copy
	assert.Error(t, err)

	// MustGetMap
	assert.Equal(t, sub, MustGetMap[string, interface{}](props, "map"))
	assert.Panics(t, func() { MustGetMap[string, interface{}](props, "missing") })

	// Legacy
	assert.Equal(t, sub, GetMapPropOrDefault[string, interface{}](props, "map", nil))
	assert.Nil(t, GetMapPropOrDefault[string, interface{}](props, "missing", nil))
}

func TestArrayMore(t *testing.T) {
	props := map[string]interface{}{
		"dates": []interface{}{
			time.Now().Format(time.RFC3339),
			time.Now().UnixMilli(),
		},
		"badDates": []interface{}{"not a date"},
		"strs":     []string{"a", "b"},
		"objs":     []interface{}{map[string]interface{}{"a": 1}},
	}

	// Date Array
	dates, err := GetDateArray(props, "dates")
	assert.NoError(t, err)
	assert.Len(t, dates, 2)

	_, err = GetDateArray(props, "badDates")
	assert.Error(t, err)

	assert.Len(t, MustGetDateArray(props, "dates"), 2)
	assert.Panics(t, func() { MustGetDateArray(props, "badDates") })

	assert.Len(t, GetDateArrayOrDefault(props, "dates", nil), 2)
	assert.Nil(t, GetDateArrayOrDefault(props, "badDates", nil))

	// Legacy
	assert.Len(t, GetDateArrayPropOrDefault(props, "dates", nil), 2)

	// String Array Extra
	assert.Equal(t, []string{"a", "b"}, GetStringArrayPropOrDefault(props, "strs", nil))
	assert.Equal(t, []string{"a", "b"}, GetStringArrayPropOrThrow(props, "strs"))
	assert.Panics(t, func() { GetStringArrayPropOrThrow(props, "missing") })

	// Object Array Extra
	assert.Len(t, GetObjectArrayPropOrDefault[map[string]interface{}](props, "objs", nil), 1)
}

func TestFunctionProps(t *testing.T) {
	props := map[string]interface{}{
		"str":  "val",
		"bool": true,
		"date": time.Now().Format(time.RFC3339),
		"obj":  map[string]interface{}{"a": 1},
		"objArr": []interface{}{
			map[string]interface{}{"a": 1},
		},
	}

	// String
	assert.Equal(t, "val", GetStringPropOrDefaultFunction(props, "str", func() string { return "def" }))
	assert.Equal(t, "def", GetStringPropOrDefaultFunction(props, "missing", func() string { return "def" }))

	// Bool
	assert.True(t, GetBooleanPropOrDefaultFunction(props, "bool", func() bool { return false }))
	assert.False(t, GetBooleanPropOrDefaultFunction(props, "missing", func() bool { return false }))

	// Constructor Bool
	assert.True(t, GetBooleanFunctionPropOrDefault(props, "bool", func(v interface{}) bool { return v.(bool) }, false))
	assert.False(t, GetBooleanFunctionPropOrDefault(props, "missing", func(v interface{}) bool { return true }, false))

	assert.True(t, GetBooleanFunctionPropOrDefaultFunction(props, "bool", func(v interface{}) bool { return true }, func() bool { return false }))
	assert.False(t, GetBooleanFunctionPropOrDefaultFunction(props, "missing", func(v interface{}) bool { return true }, func() bool { return false }))

	// Number
	assert.Equal(t, 1, GetNumberPropOrDefaultFunction(props, "missing", func() int { return 1 }))

	// Date
	assert.Equal(t, time.Time{}, GetDatePropOrDefaultFunction(props, "missing", func() time.Time { return time.Time{} }))

	// Object Constructor
	res := GetObjectFunctionPropOrDefault(props, "obj", func(m map[string]interface{}) int {
		return 10
	}, 0)
	assert.Equal(t, 10, res)
	// Missing
	assert.Equal(t, 0, GetObjectFunctionPropOrDefault(props, "missing", func(m map[string]interface{}) int { return 10 }, 0))

	res = GetObjectFunctionPropOrThrow(props, "obj", func(m map[string]interface{}) int {
		return 20
	})
	assert.Equal(t, 20, res)
	assert.Panics(t, func() { GetObjectFunctionPropOrThrow[int](props, "missing", nil) })
	assert.Panics(t, func() { GetObjectFunctionPropOrThrow[int](props, "missing", nil, "msg") })

	// Object Array Constructor
	arrRes := GetObjectArrayFunctionPropOrDefault(props, "objArr", func(m map[string]interface{}) int {
		return 1
	}, nil)
	assert.Equal(t, []int{1}, arrRes)
	assert.Nil(t, GetObjectArrayFunctionPropOrDefault(props, "missing", func(m map[string]interface{}) int { return 1 }, nil))
}

func TestNullAllow(t *testing.T) {
	props := map[string]interface{}{
		"null": nil,
		"str":  "s",
	}

	val := GetObjectPropOrDefaultAllowNull(props, "null", "default")
	assert.Nil(t, val) // Should be nil because explicit null?

	val = GetObjectPropOrDefaultAllowNull(props, "str", "default")
	assert.Equal(t, "s", *val)

	val = GetObjectPropOrDefaultAllowNull(props, "missing", "default")
	assert.Equal(t, "default", *val)
}

func TestLegacyMissing(t *testing.T) {
	props := map[string]interface{}{
		"val": "s",
	}
	assert.Equal(t, "s", GetStringPropOrThrow(props, "val", "msg"))
	assert.Panics(t, func() { GetStringPropOrThrow(props, "missing", "msg") })

	assert.Panics(t, func() { GetBooleanPropOrThrow(props, "val", "msg") }) // This should fail as val is string
}

func TestAllLegacyOrThrow(t *testing.T) {
	props := map[string]interface{}{
		"str":  "s",
		"num":  1,
		"bool": true,
		"date": time.Now(),
		"obj":  map[string]interface{}{},
	}

	assert.Equal(t, "s", GetStringPropOrThrow(props, "str"))
	assert.Equal(t, 1, GetNumberPropOrThrow[int](props, "num"))
	assert.Equal(t, true, GetBooleanPropOrThrow(props, "bool"))
	assert.NotEqual(t, time.Time{}, GetDatePropOrThrow(props, "date"))
	assert.Equal(t, map[string]interface{}{}, GetObjectPropOrThrow[map[string]interface{}](props, "obj"))

	// Custom messages
	assert.Panics(t, func() { GetStringPropOrThrow(props, "missing", "custom") })
	assert.Panics(t, func() { GetNumberPropOrThrow[int](props, "missing", "custom") })
	assert.Panics(t, func() { GetBooleanPropOrThrow(props, "missing", "custom") })
	assert.Panics(t, func() { GetDatePropOrThrow(props, "missing", "custom") })
	assert.Panics(t, func() { GetObjectPropOrThrow[int](props, "missing", "custom") })
}

func TestAllLegacyOrDefault(t *testing.T) {
	props := map[string]interface{}{}
	assert.Equal(t, "d", GetStringPropOrDefault(props, "m", "d"))
	assert.Equal(t, 1, GetNumberPropOrDefault(props, "m", 1))
	assert.Equal(t, true, GetBooleanPropOrDefault(props, "m", true))
	assert.Equal(t, time.Time{}, GetDatePropOrDefault(props, "m", time.Time{}))
	assert.Equal(t, 1, GetObjectPropOrDefault(props, "m", 1))
}

func TestNewArrayFeatures(t *testing.T) {
	// Setup test data
	now := time.Now()
	s1, s2 := "a", "b"
	i1, i2 := 1, 2
	b1, b2 := true, false
	d1, d2 := now, now.Add(time.Hour)
	o1, o2 := map[string]interface{}{"k": 1}, map[string]interface{}{"k": 2}

	props := map[string]interface{}{
		// Array of Pointers
		"strPtrArr":  []*string{&s1, &s2},
		"intPtrArr":  []*int{&i1, &i2},
		"boolPtrArr": []*bool{&b1, &b2},
		"datePtrArr": []*time.Time{&d1, &d2},
		"objPtrArr":  []*map[string]interface{}{&o1, &o2},

		// Mixed Array with Pointers (simulating interface array)
		"strPtrMix":  []interface{}{&s1, &s2},
		"intPtrMix":  []interface{}{&i1, &i2}, // Will fail if not converted properly, but GetNumberPointerArray handles conversion
		"boolPtrMix": []interface{}{&b1, &b2},
		"datePtrMix": []interface{}{&d1, &d2},
		"objPtrMix":  []interface{}{&o1, &o2},

		// Plain arrays to test conversion to pointer arrays
		"strArr":  []string{"a", "b"},
		"intArr":  []int{1, 2},
		"boolArr": []bool{true, false},
		"dateArr": []time.Time{now, now.Add(time.Hour)},
		"objArr":  []interface{}{o1, o2},
	}

	// 1. Number Array Support
	// GetNumberArray
	nums, err := GetNumberArray[int](props, "intArr")
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, nums)

	// MustGetNumberArray
	assert.Equal(t, []int{1, 2}, MustGetNumberArray[int](props, "intArr"))

	// GetNumberArrayOrDefault
	assert.Equal(t, []int{1, 2}, GetNumberArrayOrDefault[int](props, "intArr", nil))
	assert.Equal(t, []int{10}, GetNumberArrayOrDefault[int](props, "missing", []int{10}))

	// 2. Boolean Array Support
	// GetBooleanArray
	bools, err := GetBooleanArray(props, "boolArr")
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false}, bools)

	// MustGetBooleanArray
	assert.Equal(t, []bool{true, false}, MustGetBooleanArray(props, "boolArr"))

	// GetBooleanArrayOrDefault
	assert.Equal(t, []bool{true, false}, GetBooleanArrayOrDefault(props, "boolArr", nil))

	// 3. Pointer to Slice (*[]T) Support

	// String
	strArrPtr, err := GetStringArrayPtr(props, "strArr")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, *strArrPtr)
	assert.NotNil(t, MustGetStringArrayPtr(props, "strArr"))
	assert.Nil(t, GetStringArrayPtrOrDefault(props, "missing", nil))

	// Number
	numArrPtr, err := GetNumberArrayPtr[int](props, "intArr")
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, *numArrPtr)
	assert.NotNil(t, MustGetNumberArrayPtr[int](props, "intArr"))
	assert.Nil(t, GetNumberArrayPtrOrDefault[int](props, "missing", nil))

	// Boolean
	boolArrPtr, err := GetBooleanArrayPtr(props, "boolArr")
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false}, *boolArrPtr)
	assert.NotNil(t, MustGetBooleanArrayPtr(props, "boolArr"))
	assert.Nil(t, GetBooleanArrayPtrOrDefault(props, "missing", nil))

	// Date
	dateArrPtr, err := GetDateArrayPtr(props, "dateArr")
	assert.NoError(t, err)
	assert.Len(t, *dateArrPtr, 2)
	assert.NotNil(t, MustGetDateArrayPtr(props, "dateArr"))
	assert.Nil(t, GetDateArrayPtrOrDefault(props, "missing", nil))

	// Object
	objArrPtr, err := GetObjectArrayPtr[map[string]interface{}](props, "objArr")
	assert.NoError(t, err)
	assert.Len(t, *objArrPtr, 2)
	assert.NotNil(t, MustGetObjectArrayPtr[map[string]interface{}](props, "objArr"))
	assert.Nil(t, GetObjectArrayPtrOrDefault[map[string]interface{}](props, "missing", nil))

	// 4. Slice of Pointers ([]*T) Support

	// String
	strPtrs, err := GetStringPointerArray(props, "strArr")
	assert.NoError(t, err)
	assert.Equal(t, "a", *strPtrs[0])

	// Number
	numPtrs, err := GetNumberPointerArray[int](props, "intArr")
	assert.NoError(t, err)
	assert.Equal(t, 1, *numPtrs[0])

	// Boolean
	boolPtrs, err := GetBooleanPointerArray(props, "boolArr")
	assert.NoError(t, err)
	assert.Equal(t, true, *boolPtrs[0])

	// Date
	datePtrs, err := GetDatePointerArray(props, "dateArr")
	assert.NoError(t, err)
	assert.True(t, datePtrs[0].Equal(now))

	// Object
	objPtrs, err := GetObjectPointerArray[map[string]interface{}](props, "objArr")
	assert.NoError(t, err)
	assert.Equal(t, o1, *objPtrs[0])

	// Test with explicit pointer arrays
	strPtrs2, err := GetStringPointerArray(props, "strPtrArr")
	assert.NoError(t, err)
	assert.Equal(t, "a", *strPtrs2[0])

	// 5. Pointer to Slice of Pointers (*[]*T) Support

	// String
	strPtrArrPtr, err := GetStringPointerArrayPtr(props, "strArr")
	assert.NoError(t, err)
	assert.Equal(t, "a", *(*strPtrArrPtr)[0])

	// Number
	numPtrArrPtr, err := GetNumberPointerArrayPtr[int](props, "intArr")
	assert.NoError(t, err)
	assert.Equal(t, 1, *(*numPtrArrPtr)[0])

	// Boolean
	boolPtrArrPtr, err := GetBooleanPointerArrayPtr(props, "boolArr")
	assert.NoError(t, err)
	assert.Equal(t, true, *(*boolPtrArrPtr)[0])

	// Date
	datePtrArrPtr, err := GetDatePointerArrayPtr(props, "dateArr")
	assert.NoError(t, err)
	assert.True(t, (*datePtrArrPtr)[0].Equal(now))

	// Object
	objPtrArrPtr, err := GetObjectPointerArrayPtr[map[string]interface{}](props, "objArr")
	assert.NoError(t, err)
	assert.Equal(t, o1, *(*objPtrArrPtr)[0])
}
