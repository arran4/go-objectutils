package go_objectutils

import (
	"testing"
)

type MyMap map[string]interface{}

func BenchmarkGetObjectArray_Fail(b *testing.B) {
	props := map[string]interface{}{
		"objects": []interface{}{
			map[string]interface{}{"name": "obj1"},
			map[string]interface{}{"name": "obj2"},
			map[string]interface{}{"name": "obj3"},
			map[string]interface{}{"name": "obj4"},
			map[string]interface{}{"name": "obj5"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This will fail on the first element
		_, _ = GetObjectArray[MyMap](props, "objects")
	}
}

func BenchmarkGetObjectArray_Success(b *testing.B) {
	props := map[string]interface{}{
		"objects": []interface{}{
			MyMap{"name": "obj1"},
			MyMap{"name": "obj2"},
			MyMap{"name": "obj3"},
			MyMap{"name": "obj4"},
			MyMap{"name": "obj5"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetObjectArray[MyMap](props, "objects")
	}
}
