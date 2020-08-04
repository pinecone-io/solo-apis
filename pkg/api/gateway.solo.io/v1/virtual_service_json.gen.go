// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gateway/v1/virtual_service.proto

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
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/core/matchers"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VirtualServiceSpec
func (this *VirtualServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualServiceSpec
func (this *VirtualServiceSpec) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualHost
func (this *VirtualHost) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualHost
func (this *VirtualHost) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Route
func (this *Route) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Route
func (this *Route) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DelegateAction
func (this *DelegateAction) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DelegateAction
func (this *DelegateAction) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteTableSelector
func (this *RouteTableSelector) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteTableSelector
func (this *RouteTableSelector) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for VirtualServiceStatus
func (this *VirtualServiceStatus) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualServiceStatus
func (this *VirtualServiceStatus) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	VirtualServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)