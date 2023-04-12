package json

import (
	"errors"
	"io"

	"github.com/gosexy/to"
	jsoniter "github.com/json-iterator/go"
)

var (
	// ErrNPE error for nil pointer
	ErrNPE = errors.New(`invalid memory address or nil pointer dereference`)
)

// Marshal is a convenience function that converts a map[string]interface{} or struct to a JSON []byte.
// implements json.Marshaler
func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, ErrNPE
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

// MarshalToString is a convenience function that converts a interface{} to a string.
func MarshalToString(v interface{}) (string, error) {
	_byte, err := Marshal(v)
	return to.String(_byte), err
}

// MarshalToStringNoError is a convenience function that converts a interface{} to a string.
func MarshalToStringNoError(v interface{}) string {
	str, _ := MarshalToString(v)
	return str
}

// MarshalNoError is a convenience function that converts a interface{} to a byte array
func MarshalNoError(v interface{}) []byte {
	str, _ := Marshal(v)
	return str
}

// Unmarshal is a convenience function that converts a JSON []byte to a map[string]interface{} or struct.
// implements json.Unmarshaler
func Unmarshal(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

// UnmarshalString is a convenience function that converts a JSON string to a map[string]interface{} or struct.
func UnmarshalString(data string, v interface{}) error {
	return Unmarshal([]byte(data), v)
}

// NewDecoder returns a Decoder which reads JSON stream from "r".
func NewDecoder(r io.Reader) *jsoniter.Decoder {
	return jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(r)
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func  NewEncoder(w io.Writer) *jsoniter.Encoder {
	return jsoniter.ConfigCompatibleWithStandardLibrary.NewEncoder(w)
}

// DeepCopy dest must be pointer
func DeepCopy(src interface{}, dest interface{}) error {
	bytes := MarshalNoError(src)
	return Unmarshal(bytes, dest)
}
