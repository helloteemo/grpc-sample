package stringx

import (
	"errors"
	"reflect"
	"strconv"
	"unsafe"
)

var (
	// ErrInvalidStartPosition is an error that indicates the start position is invalid.
	ErrInvalidStartPosition = errors.New("start position is invalid")
	// ErrInvalidStopPosition is an error that indicates the stop position is invalid.
	ErrInvalidStopPosition = errors.New("stop position is invalid")
)

// Contains checks if str is in list.
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}

// Filter filters chars from s with given filter function.
func Filter(s string, filter func(r rune) bool) string {
	var n int
	chars := []rune(s)
	for i, x := range chars {
		if n < i {
			chars[n] = x
		}
		if !filter(x) {
			n++
		}
	}

	return string(chars[:n])
}

// HasEmpty checks if there are empty strings in args.
func HasEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}

	return false
}

// NotEmpty checks if all strings are not empty in args.
func NotEmpty(args ...string) bool {
	return !HasEmpty(args...)
}

// Remove removes given strs from strings.
func Remove(strings []string, strs ...string) []string {
	out := append([]string(nil), strings...)

	for _, str := range strs {
		var n int
		for _, v := range out {
			if v != str {
				out[n] = v
				n++
			}
		}
		out = out[:n]
	}

	return out
}

// Reverse reverses s.
func Reverse(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// Substr returns runes between start and stop [start, stop) regardless of the chars are ascii or utf8.
func Substr(str string, start int, stop int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", ErrInvalidStartPosition
	}

	if stop < 0 || stop > length {
		return "", ErrInvalidStopPosition
	}

	return string(rs[start:stop]), nil
}

// TakeOne returns valid string if not empty or later one.
func TakeOne(valid, or string) string {
	if len(valid) > 0 {
		return valid
	}

	return or
}

// TakeWithPriority returns the first not empty result from fns.
func TakeWithPriority(fns ...func() string) string {
	for _, fn := range fns {
		val := fn()
		if len(val) > 0 {
			return val
		}
	}

	return ""
}

// Union merges the strings in first and second.
func Union(first, second []string) []string {
	set := make(map[string]struct{})

	for _, each := range first {
		set[each] = struct{}{}
	}
	for _, each := range second {
		set[each] = struct{}{}
	}

	merged := make([]string, 0, len(set))
	for k := range set {
		merged = append(merged, k)
	}

	return merged
}

func Int64ToString(in int64) string {
	return strconv.Itoa(int(in))
}

func StringToInt64(in string) int64 {
	num, _ := strconv.Atoi(in)
	return int64(num)
}

// ZeroCopyString2Bytes 0拷贝把字符串替换为bytes
// 注意: 这个函数不是深拷贝
func ZeroCopyString2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// ZeroCopyBytes2String 0拷贝把byte转为string
// 注意: 这个函数不是深拷贝
func ZeroCopyBytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
