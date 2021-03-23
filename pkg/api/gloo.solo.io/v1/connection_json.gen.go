// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/connection.proto

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
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ConnectionConfig
func (this *ConnectionConfig) MarshalJSON() ([]byte, error) {
	str, err := ConnectionMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionConfig
func (this *ConnectionConfig) UnmarshalJSON(b []byte) error {
	return ConnectionUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionConfig_TcpKeepAlive
func (this *ConnectionConfig_TcpKeepAlive) MarshalJSON() ([]byte, error) {
	str, err := ConnectionMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionConfig_TcpKeepAlive
func (this *ConnectionConfig_TcpKeepAlive) UnmarshalJSON(b []byte) error {
	return ConnectionUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionConfig_HttpProtocolOptions
func (this *ConnectionConfig_HttpProtocolOptions) MarshalJSON() ([]byte, error) {
	str, err := ConnectionMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionConfig_HttpProtocolOptions
func (this *ConnectionConfig_HttpProtocolOptions) UnmarshalJSON(b []byte) error {
	return ConnectionUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ConnectionMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	ConnectionUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)