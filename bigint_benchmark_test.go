package go_objectutils

import (
	"testing"
)

func BenchmarkGetBigInt_String(b *testing.B) {
	props := map[string]interface{}{
		"val": "12345678901234567890",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetBigInt(props, "val")
	}
}

func BenchmarkGetBigInt_Int64(b *testing.B) {
	props := map[string]interface{}{
		"val": int64(1234567890),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetBigInt(props, "val")
	}
}

func BenchmarkGetBigInt_Int(b *testing.B) {
	props := map[string]interface{}{
		"val": 1234567890,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetBigInt(props, "val")
	}
}

func BenchmarkGetBigInt_Float64(b *testing.B) {
	props := map[string]interface{}{
		"val": 1234567890.0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetBigInt(props, "val")
	}
}
