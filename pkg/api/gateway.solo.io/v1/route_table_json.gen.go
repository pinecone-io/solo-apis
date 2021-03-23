// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/route_table.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RouteTableSpec
func (this *RouteTableSpec) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTableSpec
func (this *RouteTableSpec) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteTableStatus
func (this *RouteTableStatus) MarshalJSON() ([]byte, error) {
	str, err := RouteTableMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTableStatus
func (this *RouteTableStatus) UnmarshalJSON(b []byte) error {
	return RouteTableUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RouteTableMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	RouteTableUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)