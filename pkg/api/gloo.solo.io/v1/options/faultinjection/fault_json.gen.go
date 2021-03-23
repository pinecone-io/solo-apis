// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/faultinjection/fault.proto

package faultinjection

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RouteAbort
func (this *RouteAbort) MarshalJSON() ([]byte, error) {
	str, err := FaultMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteAbort
func (this *RouteAbort) UnmarshalJSON(b []byte) error {
	return FaultUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteDelay
func (this *RouteDelay) MarshalJSON() ([]byte, error) {
	str, err := FaultMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteDelay
func (this *RouteDelay) UnmarshalJSON(b []byte) error {
	return FaultUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteFaults
func (this *RouteFaults) MarshalJSON() ([]byte, error) {
	str, err := FaultMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteFaults
func (this *RouteFaults) UnmarshalJSON(b []byte) error {
	return FaultUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	FaultMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	FaultUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)