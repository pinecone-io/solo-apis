// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/grpc_web/grpc_web.proto

package grpc_web

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/transformation"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for GrpcWeb
func (this *GrpcWeb) MarshalJSON() ([]byte, error) {
	str, err := GrpcWebMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for GrpcWeb
func (this *GrpcWeb) UnmarshalJSON(b []byte) error {
	return GrpcWebUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	GrpcWebMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	GrpcWebUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)