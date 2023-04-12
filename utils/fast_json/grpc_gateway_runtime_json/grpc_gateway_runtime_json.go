package grpc_gateway_runtime_json

import (
	grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	json "github.com/helloteemo/utils/fast_json"
	"io"
)

// FastJsonBuiltin is a Marshaller which marshals/unmarshal into/from JSON
// with the standard "encoding/json" package of Golang.
// Although it is generally faster for simple proto messages than JSONPb,
// it does not support advanced features of protobuf, e.g. map, oneof, ....
//
// The NewEncoder and NewDecoder types return *json.Encoder and
// *json.Decoder respectively.
type FastJsonBuiltin struct {
}

// Marshal marshals "v" into JSON
func (f *FastJsonBuiltin) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal unmarshal JSON data into "v".
func (f *FastJsonBuiltin) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NewDecoder returns a Decoder which reads JSON stream from "r".
func (f *FastJsonBuiltin) NewDecoder(r io.Reader) grpc_gateway_runtime.Decoder {
	return json.NewDecoder(r)
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (f *FastJsonBuiltin) NewEncoder(w io.Writer) grpc_gateway_runtime.Encoder {
	return json.NewEncoder(w)
}

// ContentType always Returns "application/json".
func (f *FastJsonBuiltin) ContentType(v interface{}) string {
	return "application/json"
}

var delimiter = []byte("\n")

// Delimiter for newline encoded JSON streams.
func (f *FastJsonBuiltin) Delimiter() []byte {
	return delimiter
}
