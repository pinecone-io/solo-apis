// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/http_path/http_path.proto

package http_path

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/udpa/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for HttpPath
func (this *HttpPath) MarshalJSON() ([]byte, error) {
	str, err := HttpPathMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HttpPath
func (this *HttpPath) UnmarshalJSON(b []byte) error {
	return HttpPathUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	HttpPathMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	HttpPathUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)