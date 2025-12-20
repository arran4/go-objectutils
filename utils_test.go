package go_objectutils

import (
	"testing"
	"time"
)

func TestGetStringProp(t *testing.T) {
	props := map[string]interface{}{
		"name": "Alice",
		"age":  30,
	}

	if got := GetStringPropOrDefault(props, "name", "Bob"); got != "Alice" {
		t.Errorf("GetStringPropOrDefault = %v, want %v", got, "Alice")
	}
	if got := GetStringPropOrDefault(props, "missing", "Bob"); got != "Bob" {
		t.Errorf("GetStringPropOrDefault = %v, want %v", got, "Bob")
	}
	if got := GetStringPropOrDefault(props, "age", "Bob"); got != "Bob" {
		t.Errorf("GetStringPropOrDefault(wrong type) = %v, want %v", got, "Bob")
	}

	if got := GetStringPropOrDefaultFunction(props, "missing", func() string { return "Bob" }); got != "Bob" {
		t.Errorf("GetStringPropOrDefaultFunction = %v, want %v", got, "Bob")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("GetStringPropOrThrow did not panic on missing")
			}
		}()
		GetStringPropOrThrow(props, "missing")
	}()

	if got := GetStringPropOrThrow(props, "name"); got != "Alice" {
		t.Errorf("GetStringPropOrThrow = %v, want %v", got, "Alice")
	}
}

func TestGetNumberProp(t *testing.T) {
	props := map[string]interface{}{
		"count": 10,
		"pi":    3.14,
		"str":   "123",
	}

	// Int
	if got := GetNumberPropOrDefault[int](props, "count", 0); got != 10 {
		t.Errorf("GetNumberPropOrDefault[int] = %v, want %v", got, 10)
	}
	// Float -> Int conversion
	if got := GetNumberPropOrDefault[int](props, "pi", 0); got != 3 {
		t.Errorf("GetNumberPropOrDefault[int](float) = %v, want %v", got, 3)
	}
	// String parsing
	if got := GetNumberPropOrDefault[int](props, "str", 0); got != 123 {
		t.Errorf("GetNumberPropOrDefault[int](string) = %v, want %v", got, 123)
	}

	// Float64
	if got := GetNumberPropOrDefault[float64](props, "pi", 0.0); got != 3.14 {
		t.Errorf("GetNumberPropOrDefault[float64] = %v, want %v", got, 3.14)
	}
}

func TestGetBooleanProp(t *testing.T) {
	props := map[string]interface{}{
		"active": true,
		"f":      false,
	}

	if got := GetBooleanPropOrDefault(props, "active", false); got != true {
		t.Errorf("GetBooleanPropOrDefault = %v, want %v", got, true)
	}
	if got := GetBooleanPropOrDefault(props, "missing", true); got != true {
		t.Errorf("GetBooleanPropOrDefault = %v, want %v", got, true)
	}
}

func TestGetDateProp(t *testing.T) {
	now := time.Now()
	nowISO := now.Format(time.RFC3339)
	nowMilli := now.UnixMilli()

	props := map[string]interface{}{
		"iso":   nowISO,
		"milli": nowMilli,
		"bad":   "notadate",
	}

	d1 := GetDatePropOrDefault(props, "iso", time.Time{})
	if d1.Unix() != now.Unix() {
		t.Errorf("GetDatePropOrDefault(iso) = %v, want %v", d1, now)
	}

	d2 := GetDatePropOrDefault(props, "milli", time.Time{})
	if d2.UnixMilli() != nowMilli {
		t.Errorf("GetDatePropOrDefault(milli) = %v, want %v", d2.UnixMilli(), nowMilli)
	}

	d3 := GetDatePropOrDefault(props, "bad", time.Time{})
	if !d3.IsZero() {
		t.Errorf("GetDatePropOrDefault(bad) = %v, want zero", d3)
	}
}

func TestGetObjectProp(t *testing.T) {
	nested := map[string]interface{}{"a": 1}
	props := map[string]interface{}{
		"obj": nested,
	}

	got := GetObjectPropOrDefault[map[string]interface{}](props, "obj", nil)
	if got["a"] != 1 {
		t.Errorf("GetObjectPropOrDefault = %v, want %v", got, nested)
	}
}

type MyObj struct {
	Val int
}

func TestGetObjectFunctionProp(t *testing.T) {
	nested := map[string]interface{}{"Val": 10}
	props := map[string]interface{}{
		"obj": nested,
	}

	ctor := func(p map[string]interface{}) *MyObj {
		return &MyObj{Val: GetNumberPropOrDefault[int](p, "Val", 0)}
	}

	got := GetObjectFunctionPropOrDefault(props, "obj", ctor, nil)
	if got == nil || got.Val != 10 {
		t.Errorf("GetObjectFunctionPropOrDefault = %v, want Val=10", got)
	}
}

func TestGetArrayProp(t *testing.T) {
	props := map[string]interface{}{
		"strs": []interface{}{"a", "b"},
		"nums": []interface{}{1, 2},
	}

	strs := GetStringArrayPropOrDefault(props, "strs", nil)
	if len(strs) != 2 || strs[0] != "a" {
		t.Errorf("GetStringArrayPropOrDefault = %v", strs)
	}

	// Object array
	objsProp := map[string]interface{}{
		"users": []interface{}{
			map[string]interface{}{"name": "A"},
			map[string]interface{}{"name": "B"},
		},
	}
	type User struct {
		Name string
	}
	userCtor := func(p map[string]interface{}) User {
		return User{Name: GetStringPropOrDefault(p, "name", "")}
	}

	users := GetObjectArrayFunctionPropOrDefault(objsProp, "users", userCtor, nil)
	if len(users) != 2 || users[0].Name != "A" {
		t.Errorf("GetObjectArrayFunctionPropOrDefault = %v", users)
	}
}
