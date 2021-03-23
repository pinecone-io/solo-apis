// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed.gloo/v1/settings.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/fed.solo.io/core/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	_ "github.com/solo-io/solo-apis/pkg/api/multicluster.solo.io/v1alpha1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for FederatedSettingsSpec
func (this *FederatedSettingsSpec) MarshalJSON() ([]byte, error) {
	str, err := SettingsMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedSettingsSpec
func (this *FederatedSettingsSpec) UnmarshalJSON(b []byte) error {
	return SettingsUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FederatedSettingsSpec_Template
func (this *FederatedSettingsSpec_Template) MarshalJSON() ([]byte, error) {
	str, err := SettingsMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedSettingsSpec_Template
func (this *FederatedSettingsSpec_Template) UnmarshalJSON(b []byte) error {
	return SettingsUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for FederatedSettingsStatus
func (this *FederatedSettingsStatus) MarshalJSON() ([]byte, error) {
	str, err := SettingsMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FederatedSettingsStatus
func (this *FederatedSettingsStatus) UnmarshalJSON(b []byte) error {
	return SettingsUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	SettingsMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	SettingsUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)