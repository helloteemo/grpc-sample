package json

import "testing"

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Marshal(map[string]interface{}{"a": 1})
	}
}
