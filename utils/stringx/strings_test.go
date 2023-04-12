package stringx

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestZeroCopyString2Bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "123",
			args: args{
				s: "hello,teemo",
			},
			want: []byte("hello,teemo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroCopyString2Bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroCopyString2Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZeroCopyBytes2String(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "123",
			args: args{[]byte(`teemo`)},
			want: "teemo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroCopyBytes2String(tt.args.b); got != tt.want {
				t.Errorf("ZeroCopyBytes2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkZeroCopyString2Bytes(b *testing.B) {
	s := uuid.New().String()
	for i := 0; i < b.N; i++ {
		_ = ZeroCopyString2Bytes(s)
	}
}

func BenchmarkString2Bytes(b *testing.B) {
	s := uuid.New().String()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkZeroCopyBytes2String(b *testing.B) {
	s := uuid.New().String()
	bytes := []byte(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ZeroCopyBytes2String(bytes)
	}
}

func BenchmarkBytes2String(b *testing.B) {
	s := uuid.New().String()
	bytes := []byte(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(bytes)
	}
}
