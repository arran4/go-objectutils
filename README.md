# Go Object Utils

![Go](https://github.com/arran4/go-objectutils/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/arran4/go-objectutils)](https://goreportcard.com/report/github.com/arran4/go-objectutils)
[![GoDoc](https://godoc.org/github.com/arran4/go-objectutils?status.svg)](https://godoc.org/github.com/arran4/go-objectutils)

Extensive utilities for extracting typed values from `map[string]interface{}` (commonly used for JSON objects in Go) with type safety, default values, and error handling.

This library is designed to help working with loosely typed data structures (like parsed JSON or YAML) in a safer and more ergonomic way.

## Justification for `Must` (Panic) Patterns

This library provides `Must*` functions that panic on error. While idiomatic Go prefers returning errors, there are specific scenarios where panicking is appropriate or preferred for developer ergonomics, especially when:

1.  **Configuration Loading**: When an application starts, if essential configuration is missing or invalid, it is often better to fail fast (panic) than to proceed with a broken state. `Must` functions simplify this code by removing the need for verbose error checking for every single property.
2.  **Consistency**: This library aims for consistency with sister libraries in other languages:
    *   [tsobjectutils](https://github.com/arran4/tsobjectutils) (TypeScript)
    *   [dartobjectutils](https://github.com/arran4/dartobjectutils) (Dart)

    These libraries use similar patterns to assert presence and type of fields.
3.  **Chaining/Inlining**: `Must` functions return the value directly, allowing them to be used inline in expressions or struct initialization literals where multi-value returns are not permitted.

**However**, for general library usage or handling user input where failure is expected and recoverable, **you should use the error-returning versions** (e.g., `GetString`) as the primary way of interacting with this library.

## Installation

```bash
go get github.com/arran4/go-objectutils
```

## Usage

### Primitives

The library supports all Go primitive types: `string`, `int`, `int64`, `float64`, `bool`, etc.

#### Error Returning (Recommended)

Use these functions when you want to handle errors gracefully.

```go
import "github.com/arran4/go-objectutils"

data := map[string]interface{}{
    "name": "John",
    "age":  30,
}

name, err := go_objectutils.GetString(data, "name")
if err != nil {
    // handle missing or invalid type
}

age, err := go_objectutils.GetNumber[int](data, "age")
if err != nil {
    // handle error
}
```

#### Must (Panic)

Use these when you want to fail fast if the data is invalid.

```go
name := go_objectutils.MustGetString(data, "name")
age := go_objectutils.MustGetNumber[int](data, "age")
```

#### Default Values

Use these to provide a fallback if the property is missing or invalid.

```go
// Returns "Guest" if "name" is missing
name := go_objectutils.GetStringOrDefault(data, "name", "Guest")

// Returns 0 if "count" is missing
count := go_objectutils.GetNumberOrDefault(data, "count", 0)
```

### Pointers

You can retrieve pointers to values, which is useful for distinguishing between "missing/null" (nil pointer) and "zero value" (e.g., empty string or 0).

```go
// Returns *string, or error
namePtr, err := go_objectutils.GetStringPtr(data, "name")

// Returns *int, panics on error
agePtr := go_objectutils.MustGetNumberPtr[int](data, "age")
```

### Arrays / Slices

```go
data := map[string]interface{}{
    "tags": []interface{}{"go", "json"},
}

tags, err := go_objectutils.GetStringArray(data, "tags")
// tags is []string{"go", "json"}
```

### Nested Objects

```go
data := map[string]interface{}{
    "user": map[string]interface{}{
        "id": 123,
    },
}

userObj, err := go_objectutils.GetObject[map[string]interface{}](data, "user")
```

### Dates

Dates are parsed from strings (RFC3339/ISO8601) or timestamps (int/float).

```go
data := map[string]interface{}{
    "created_at": "2023-10-01T12:00:00Z",
}

createdAt, err := go_objectutils.GetDate(data, "created_at")
// createdAt is time.Time
```

## API Reference

The library follows a consistent naming convention:

*   `Get<Type>(props, key) (<Type>, error)`
*   `MustGet<Type>(props, key) <Type>`
*   `Get<Type>OrDefault(props, key, default) <Type>`
*   `Get<Type>Ptr(props, key) (*<Type>, error)`
*   `MustGet<Type>Ptr(props, key) *<Type>`

Supported Types:
*   `String`
*   `Number` (Generics: `int`, `float64`, etc.)
*   `Boolean`
*   `Date` (`time.Time`)
*   `StringArray`
*   `ObjectArray`
*   `DateArray`
*   `Object` / `Map`

## Error Handling

Errors returned are typed:
*   `*go_objectutils.MissingFieldError`: The property does not exist.
*   `*go_objectutils.InvalidTypeError`: The property exists but is not of the expected type.
