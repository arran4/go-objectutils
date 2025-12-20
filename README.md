# Go Object Utils

![Go](https://github.com/arran4/go-objectutils/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/arran4/go-objectutils)](https://goreportcard.com/report/github.com/arran4/go-objectutils)
[![GoDoc](https://godoc.org/github.com/arran4/go-objectutils?status.svg)](https://godoc.org/github.com/arran4/go-objectutils)

Extensive utilities for extracting typed values from `map[string]interface{}` (commonly used for JSON objects in Go) with type safety, default values, and error handling.

This library is designed to help working with loosely typed data structures (like parsed JSON or YAML) in a safer and more ergonomic way, providing a bridge between dynamic data and Go's static typing.

## Installation

```bash
go get github.com/arran4/go-objectutils
```

## Concept: Error Handling Patterns

This library offers three primary ways to access data, allowing you to choose the best strategy for your specific use case (e.g., failing fast, safe handling, or default fallbacks).

### 1. Error Returning (`Get*`)

**Use Case:** General application logic where missing data is expected or recoverable.

These functions return the value and an `error`. The error will be `*MissingFieldError` if the key doesn't exist, or `*InvalidTypeError` if the value cannot be converted to the target type.

```go
val, err := go_objectutils.GetString(data, "key")
if err != nil {
    // Handle error
}
```

### 2. Panic (`MustGet*`)

**Use Case:** Configuration loading, scripts, or scenarios where the application cannot proceed without this data. Fails fast.

These functions return the value directly but will `panic` if an error occurs. This keeps the code concise when safety is guaranteed by other means or when crashing is the desired behavior on failure.

```go
val := go_objectutils.MustGetString(data, "key")
```

### 3. Default Values (`Get*OrDefault`)

**Use Case:** Optional configuration, UI rendering where defaults are acceptable.

These functions return the value if found and valid; otherwise, they return the provided default value. Errors are swallowed.

```go
val := go_objectutils.GetStringOrDefault(data, "key", "default")
```

### 4. Pointers (`Get*Ptr`)

**Use Case:** Distinguishing between a "zero value" (e.g., `""`, `0`, `false`) and a "missing/null" value.

```go
valPtr, err := go_objectutils.GetStringPtr(data, "key")
if valPtr == nil {
    // Key was missing or null
}
```

## Function Listing

### Strings

| Function | Description |
| :--- | :--- |
| `GetString` | Returns `string` or error. |
| `MustGetString` | Returns `string` or panics. |
| `GetStringOrDefault` | Returns `string` or default value. |
| `GetStringPtr` | Returns `*string` or error. |
| `MustGetStringPtr` | Returns `*string` or panics. |
| `GetStringPtrOrDefault` | Returns `*string` or default value. |

### Numbers (Generics)

Supports `int`, `int8`...`int64`, `uint`...`uint64`, `float32`, `float64`.
*Note: Also handles automatic conversion from string representations of numbers.*

| Function | Description |
| :--- | :--- |
| `GetNumber[T]` | Returns number of type `T` or error. |
| `MustGetNumber[T]` | Returns number of type `T` or panics. |
| `GetNumberOrDefault[T]` | Returns number of type `T` or default value. |
| `GetNumberPtr[T]` | Returns `*T` or error. |
| `MustGetNumberPtr[T]` | Returns `*T` or panics. |
| `GetNumberPtrOrDefault[T]` | Returns `*T` or default value. |

### Booleans

| Function | Description |
| :--- | :--- |
| `GetBoolean` | Returns `bool` or error. |
| `MustGetBoolean` | Returns `bool` or panics. |
| `GetBooleanOrDefault` | Returns `bool` or default value. |
| `GetBooleanPtr` | Returns `*bool` or error. |
| `MustGetBooleanPtr` | Returns `*bool` or panics. |
| `GetBooleanPtrOrDefault` | Returns `*bool` or default value. |

### Dates

Parses `time.Time` from strings (RFC3339) or timestamps (int/float milliseconds).

| Function | Description |
| :--- | :--- |
| `GetDate` | Returns `time.Time` or error. |
| `MustGetDate` | Returns `time.Time` or panics. |
| `GetDateOrDefault` | Returns `time.Time` or default value. |
| `GetDatePtr` | Returns `*time.Time` or error. |
| `MustGetDatePtr` | Returns `*time.Time` or panics. |
| `GetDatePtrOrDefault` | Returns `*time.Time` or default value. |

### BigInt

Extracts `math/big.Int` from strings or numbers.

| Function | Description |
| :--- | :--- |
| `GetBigInt` | Returns `*big.Int` or error. |
| `MustGetBigInt` | Returns `*big.Int` or panics. |
| `GetBigIntOrDefault` | Returns `*big.Int` or default value. |

### Arrays / Slices

| Function | Description |
| :--- | :--- |
| `GetStringArray` | Returns `[]string` or error. |
| `MustGetStringArray` | Returns `[]string` or panics. |
| `GetStringArrayOrDefault` | Returns `[]string` or default value. |
| `GetDateArray` | Returns `[]time.Time` or error. |
| `MustGetDateArray` | Returns `[]time.Time` or panics. |
| `GetDateArrayOrDefault` | Returns `[]time.Time` or default value. |
| `GetObjectArray[T]` | Returns `[]T` or error. Useful for lists of sub-objects. |
| `MustGetObjectArray[T]` | Returns `[]T` or panics. |
| `GetObjectArrayOrDefault[T]` | Returns `[]T` or default value. |

### Objects / Maps

| Function | Description |
| :--- | :--- |
| `GetObject[T]` | Returns object cast to type `T` (usually `map[string]interface{}`) or error. |
| `MustGetObject[T]` | Returns object cast to type `T` or panics. |
| `GetObjectOrDefault[T]` | Returns object cast to type `T` or default value. |
| `GetObjectPtr[T]` | Returns `*T` or error. |
| `MustGetObjectPtr[T]` | Returns `*T` or panics. |
| `GetObjectPtrOrDefault[T]` | Returns `*T` or default value. |
| `GetMap[K, V]` | Returns `map[K]V` or error. |
| `MustGetMap[K, V]` | Returns `map[K]V` or panics. |

## Use Cases & Examples

### Scenario 1: Parsing Configuration

When loading a configuration file (e.g., parsed from JSON), you often want to fail if required fields are missing and use defaults for others.

```go
configData := map[string]interface{}{
    "host": "localhost",
    "port": 8080,
    // "timeout": missing, will use default
}

// Fail if host is missing
host := go_objectutils.MustGetString(configData, "host")

// Fail if port is missing or not a number
port := go_objectutils.MustGetNumber[int](configData, "port")

// Use default if timeout is missing
timeout := go_objectutils.GetNumberOrDefault[int](configData, "timeout", 30)
```

### Scenario 2: Handling API Responses

When processing an external API response, you should handle potential schema changes gracefully using the error-returning functions.

```go
response := map[string]interface{}{
    "user": map[string]interface{}{
        "id": "12345",
        "active": true,
    },
    "tags": []interface{}{"admin", "editor"},
}

// Safely access nested object
if userMap, err := go_objectutils.GetObject[map[string]interface{}](response, "user"); err == nil {
    // Note: ID is a string in JSON but we want an int
    id, _ := go_objectutils.GetNumber[int](userMap, "id")

    active := go_objectutils.GetBooleanOrDefault(userMap, "active", false)
    fmt.Printf("User ID: %d, Active: %v\n", id, active)
}

// Safely access array
tags, _ := go_objectutils.GetStringArrayOrDefault(response, "tags", []string{})
```

### Scenario 3: Nullable Fields (Pointers)

Use pointer variants to check if a field was explicitly set to null or missing, versus just being the zero value.

```go
updateData := map[string]interface{}{
    "description": "",   // User wants to clear description
    "age": nil,         // User didn't update age
}

// Check description
descPtr, _ := go_objectutils.GetStringPtr(updateData, "description")
if descPtr != nil {
    fmt.Printf("Updating description to: '%s'\n", *descPtr)
}

// Check age
agePtr, _ := go_objectutils.GetNumberPtr[int](updateData, "age")
if agePtr == nil {
    fmt.Println("Age not updated")
}
```

### Scenario 4: Custom Object Mapping

You can use `GetObjectArray` with a constructor function approach (via legacy helpers or manual iteration) or simple type casting if your structures match.

```go
data := map[string]interface{}{
    "users": []interface{}{
        map[string]interface{}{"name": "Alice"},
        map[string]interface{}{"name": "Bob"},
    },
}

// Extract list of user maps
users := go_objectutils.GetObjectArrayOrDefault[map[string]interface{}](data, "users", nil)

for _, u := range users {
    name := go_objectutils.GetStringOrDefault(u, "name", "Unknown")
    fmt.Println(name)
}
```

## Legacy Functions

The library includes several legacy functions for backward compatibility. These are generally more verbose aliases or specific utility wrappers.

*   `GetStringPropOrDefault` -> `GetStringOrDefault`
*   `GetNumberPropOrDefault` -> `GetNumberOrDefault`
*   `GetBooleanPropOrDefault` -> `GetBooleanOrDefault`
*   `GetDatePropOrDefault` -> `GetDateOrDefault`
*   `GetObjectPropOrDefault` -> `GetObjectOrDefault`
*   `GetStringPropOrThrow`, `GetNumberPropOrThrow`, etc. -> Custom error message wrappers around `Get*`.
*   `Get*PropOrDefaultFunction` -> Uses a function to generate the default value (lazy evaluation).
