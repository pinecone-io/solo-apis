// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/status.proto

package annotations

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for StatusAnnotation
func (this *StatusAnnotation) MarshalJSON() ([]byte, error) {
	str, err := StatusMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StatusAnnotation
func (this *StatusAnnotation) UnmarshalJSON(b []byte) error {
	return StatusUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	StatusMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	StatusUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)