// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed/v1/failover.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The State of a reconciled object
type FailoverSchemeStatus_State int32

const (
	// Waiting to be processed.
	FailoverSchemeStatus_PENDING FailoverSchemeStatus_State = 0
	// Currently processing.
	FailoverSchemeStatus_PROCESSING FailoverSchemeStatus_State = 1
	// Invalid parameters supplied, will not continue.
	FailoverSchemeStatus_INVALID FailoverSchemeStatus_State = 2
	// Failed during processing.
	FailoverSchemeStatus_FAILED FailoverSchemeStatus_State = 3
	// Finished processing successfully.
	FailoverSchemeStatus_ACCEPTED FailoverSchemeStatus_State = 4
)

var FailoverSchemeStatus_State_name = map[int32]string{
	0: "PENDING",
	1: "PROCESSING",
	2: "INVALID",
	3: "FAILED",
	4: "ACCEPTED",
}

var FailoverSchemeStatus_State_value = map[string]int32{
	"PENDING":    0,
	"PROCESSING": 1,
	"INVALID":    2,
	"FAILED":     3,
	"ACCEPTED":   4,
}

func (x FailoverSchemeStatus_State) String() string {
	return proto.EnumName(FailoverSchemeStatus_State_name, int32(x))
}

func (FailoverSchemeStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_29a7fd16046f32dd, []int{1, 0}
}

//
//FailoverSpec is the core portion of the API for enabling failover between Gloo Upstreams in gloo-fed.
//This API is heavily inspired by the Failover API present in the Gloo Upstream which can be found in
//`api/gloo/v1/upstream`.
//
//The source Upstream below is the initial primary target of traffic. The type of endpoints vary by the type
//of Upstream specified. Each target specified is then configured as a failover endpoint in the case that
//the prmiary Upstream becomes unhealthy. The priority of the failover endpoints is inferred from the
//order in which the Upstreams are specified. source = [0], targets = [1-n].
//
//Example:
//
//primary:
//cluster: primary
//name: primary
//namespace: primary
//failover_groups:
//- priority_group:
//- cluster: A
//upstreams:
//- name: one
//namespace: one
//- cluster: B
//upstreams:
//- name: two
//namespace: two
//- priority_group:
//- cluster: C
//upstreams:
//- name: one
//namespace: one
//- cluster: D
//upstreams:
//- name: two
//namespace: two
type FailoverSchemeSpec struct {
	// The upstream which will be configured for failover.
	Primary              *v1.ClusterObjectRef                    `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	FailoverGroups       []*FailoverSchemeSpec_FailoverEndpoints `protobuf:"bytes,2,rep,name=failover_groups,json=failoverGroups,proto3" json:"failover_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *FailoverSchemeSpec) Reset()         { *m = FailoverSchemeSpec{} }
func (m *FailoverSchemeSpec) String() string { return proto.CompactTextString(m) }
func (*FailoverSchemeSpec) ProtoMessage()    {}
func (*FailoverSchemeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_29a7fd16046f32dd, []int{0}
}
func (m *FailoverSchemeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverSchemeSpec.Unmarshal(m, b)
}
func (m *FailoverSchemeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverSchemeSpec.Marshal(b, m, deterministic)
}
func (m *FailoverSchemeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverSchemeSpec.Merge(m, src)
}
func (m *FailoverSchemeSpec) XXX_Size() int {
	return xxx_messageInfo_FailoverSchemeSpec.Size(m)
}
func (m *FailoverSchemeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverSchemeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverSchemeSpec proto.InternalMessageInfo

func (m *FailoverSchemeSpec) GetPrimary() *v1.ClusterObjectRef {
	if m != nil {
		return m.Primary
	}
	return nil
}

func (m *FailoverSchemeSpec) GetFailoverGroups() []*FailoverSchemeSpec_FailoverEndpoints {
	if m != nil {
		return m.FailoverGroups
	}
	return nil
}

type FailoverSchemeSpec_FailoverEndpoints struct {
	PriorityGroup        []*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets `protobuf:"bytes,2,rep,name=priority_group,json=priorityGroup,proto3" json:"priority_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *FailoverSchemeSpec_FailoverEndpoints) Reset()         { *m = FailoverSchemeSpec_FailoverEndpoints{} }
func (m *FailoverSchemeSpec_FailoverEndpoints) String() string { return proto.CompactTextString(m) }
func (*FailoverSchemeSpec_FailoverEndpoints) ProtoMessage()    {}
func (*FailoverSchemeSpec_FailoverEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_29a7fd16046f32dd, []int{0, 0}
}
func (m *FailoverSchemeSpec_FailoverEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints.Unmarshal(m, b)
}
func (m *FailoverSchemeSpec_FailoverEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints.Marshal(b, m, deterministic)
}
func (m *FailoverSchemeSpec_FailoverEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints.Merge(m, src)
}
func (m *FailoverSchemeSpec_FailoverEndpoints) XXX_Size() int {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints.Size(m)
}
func (m *FailoverSchemeSpec_FailoverEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints proto.InternalMessageInfo

func (m *FailoverSchemeSpec_FailoverEndpoints) GetPriorityGroup() []*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets {
	if m != nil {
		return m.PriorityGroup
	}
	return nil
}

type FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets struct {
	// (REQUIRED) Cluster on which the endpoints for this Group can be found
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// A list of Upstream targets, each of these targets must exist on the cluster specified in this message
	Upstreams []*v1.ObjectRef `protobuf:"bytes,2,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	// (optional) locality load balancing weight assigned to the specified upstreams.
	// Locality load balancing will add a special load balancing weight among all
	// targets within a given priority, who are located in the zame zone.
	// See envoy Locality Weighted Load Balancing for more information:
	// https://www.envoyproxy.io/docs/envoy/v1.14.1/intro/arch_overview/upstream/load_balancing/locality_weight#arch-overview-load-balancing-locality-weighted-lb
	LocalityWeight       *types.UInt32Value `protobuf:"bytes,3,opt,name=locality_weight,json=localityWeight,proto3" json:"locality_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) Reset() {
	*m = FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets{}
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) String() string {
	return proto.CompactTextString(m)
}
func (*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) ProtoMessage() {}
func (*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) Descriptor() ([]byte, []int) {
	return fileDescriptor_29a7fd16046f32dd, []int{0, 0, 0}
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.Unmarshal(m, b)
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.Marshal(b, m, deterministic)
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.Merge(m, src)
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) XXX_Size() int {
	return xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.Size(m)
}
func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets proto.InternalMessageInfo

func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetUpstreams() []*v1.ObjectRef {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) GetLocalityWeight() *types.UInt32Value {
	if m != nil {
		return m.LocalityWeight
	}
	return nil
}

type FailoverSchemeStatus struct {
	// The current state of the resource
	State FailoverSchemeStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=fed.solo.io.FailoverSchemeStatus_State" json:"state,omitempty"`
	// A human readable message about the current state of the object
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The most recently observed generation of the resource. This value corresponds to the `metadata.generation` of
	// a kubernetes resource
	ObservedGeneration int64 `protobuf:"varint,3,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The time at which this status was recorded
	ProcessingTime       *types.Timestamp `protobuf:"bytes,4,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FailoverSchemeStatus) Reset()         { *m = FailoverSchemeStatus{} }
func (m *FailoverSchemeStatus) String() string { return proto.CompactTextString(m) }
func (*FailoverSchemeStatus) ProtoMessage()    {}
func (*FailoverSchemeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_29a7fd16046f32dd, []int{1}
}
func (m *FailoverSchemeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverSchemeStatus.Unmarshal(m, b)
}
func (m *FailoverSchemeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverSchemeStatus.Marshal(b, m, deterministic)
}
func (m *FailoverSchemeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverSchemeStatus.Merge(m, src)
}
func (m *FailoverSchemeStatus) XXX_Size() int {
	return xxx_messageInfo_FailoverSchemeStatus.Size(m)
}
func (m *FailoverSchemeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverSchemeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverSchemeStatus proto.InternalMessageInfo

func (m *FailoverSchemeStatus) GetState() FailoverSchemeStatus_State {
	if m != nil {
		return m.State
	}
	return FailoverSchemeStatus_PENDING
}

func (m *FailoverSchemeStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FailoverSchemeStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *FailoverSchemeStatus) GetProcessingTime() *types.Timestamp {
	if m != nil {
		return m.ProcessingTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("fed.solo.io.FailoverSchemeStatus_State", FailoverSchemeStatus_State_name, FailoverSchemeStatus_State_value)
	proto.RegisterType((*FailoverSchemeSpec)(nil), "fed.solo.io.FailoverSchemeSpec")
	proto.RegisterType((*FailoverSchemeSpec_FailoverEndpoints)(nil), "fed.solo.io.FailoverSchemeSpec.FailoverEndpoints")
	proto.RegisterType((*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets)(nil), "fed.solo.io.FailoverSchemeSpec.FailoverEndpoints.LocalityLbTargets")
	proto.RegisterType((*FailoverSchemeStatus)(nil), "fed.solo.io.FailoverSchemeStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo-fed/fed/v1/failover.proto", fileDescriptor_29a7fd16046f32dd)
}

var fileDescriptor_29a7fd16046f32dd = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0x1b, 0x3d,
	0x14, 0xfd, 0x92, 0xf0, 0xf3, 0xe1, 0xb4, 0x21, 0xb8, 0x2c, 0xa2, 0x14, 0x51, 0x44, 0x17, 0x65,
	0x83, 0x47, 0x84, 0x55, 0x2b, 0xa1, 0x8a, 0x4e, 0x06, 0x14, 0x35, 0x02, 0x34, 0xa1, 0x54, 0x62,
	0x13, 0x39, 0x93, 0x1b, 0x63, 0x98, 0x89, 0x2d, 0xdb, 0x13, 0xc8, 0x1b, 0x75, 0x59, 0xb1, 0xea,
	0xf3, 0x54, 0x7d, 0x85, 0xee, 0x2b, 0x7b, 0x66, 0x08, 0x22, 0x15, 0x52, 0x17, 0x91, 0x73, 0x7f,
	0xce, 0xbd, 0xe7, 0x1c, 0x8f, 0xd1, 0x47, 0xc6, 0xcd, 0x55, 0x3a, 0x20, 0x91, 0x48, 0x3c, 0x2d,
	0x62, 0xb1, 0xcb, 0x45, 0x76, 0x52, 0xc9, 0xb5, 0x47, 0x25, 0xf7, 0x58, 0x2c, 0xc4, 0xee, 0x08,
	0x86, 0x9e, 0xfd, 0x4d, 0xf6, 0xbc, 0x11, 0xe5, 0xb1, 0x98, 0x80, 0x22, 0x52, 0x09, 0x23, 0x70,
	0x75, 0x04, 0x43, 0x62, 0x11, 0x84, 0x8b, 0xe6, 0x3a, 0x13, 0x4c, 0xb8, 0xbc, 0x67, 0xff, 0x65,
	0x2d, 0x4d, 0x0c, 0x77, 0x26, 0x4b, 0xc2, 0x9d, 0xc9, 0x73, 0x9b, 0x4c, 0x08, 0x16, 0x83, 0xe7,
	0xa2, 0x41, 0x3a, 0xf2, 0x6e, 0x15, 0x95, 0x12, 0x94, 0xce, 0xeb, 0x6f, 0x9e, 0xd6, 0x0d, 0x4f,
	0x40, 0x1b, 0x9a, 0xc8, 0xbc, 0xe1, 0xb5, 0xbe, 0x99, 0xb4, 0x1c, 0xc1, 0x48, 0x28, 0xb0, 0xc4,
	0xec, 0x99, 0x15, 0xb7, 0x7f, 0x55, 0x10, 0x3e, 0xca, 0x79, 0xf6, 0xa2, 0x2b, 0x48, 0xa0, 0x27,
	0x21, 0xc2, 0x07, 0x68, 0x59, 0x2a, 0x9e, 0x50, 0x35, 0x6d, 0x94, 0xb6, 0x4a, 0x3b, 0xd5, 0xd6,
	0x5b, 0xe2, 0x40, 0x76, 0x54, 0xa1, 0x81, 0xf8, 0x71, 0xaa, 0x0d, 0xa8, 0xd3, 0xc1, 0x35, 0x44,
	0x26, 0x84, 0x51, 0x58, 0x60, 0xf0, 0x25, 0x5a, 0x2d, 0xc4, 0xf7, 0x99, 0x12, 0xa9, 0xd4, 0x8d,
	0xf2, 0x56, 0x65, 0xa7, 0xda, 0xda, 0x23, 0x8f, 0x4c, 0x20, 0xf3, 0x8b, 0x1f, 0x52, 0xc1, 0x78,
	0x28, 0x05, 0x1f, 0x1b, 0x1d, 0xd6, 0x8a, 0x49, 0xc7, 0x6e, 0x50, 0xf3, 0xbe, 0x8c, 0xd6, 0xe6,
	0xba, 0xf0, 0x35, 0xaa, 0x49, 0xc5, 0x85, 0xe2, 0x66, 0x9a, 0x6d, 0xcc, 0x17, 0xfa, 0xff, 0xbc,
	0x90, 0x74, 0x45, 0x44, 0x63, 0x6e, 0xa6, 0xdd, 0xc1, 0x39, 0x55, 0x0c, 0x8c, 0x0e, 0x5f, 0x16,
	0xa3, 0x1d, 0x85, 0xe6, 0xf7, 0x12, 0x5a, 0x9b, 0x6b, 0xc2, 0x0d, 0xb4, 0x1c, 0x65, 0x86, 0x38,
	0xcb, 0x56, 0xc2, 0x22, 0xc4, 0x1f, 0xd0, 0x4a, 0x2a, 0xb5, 0x51, 0x40, 0x93, 0xc2, 0x87, 0x8d,
	0xbf, 0xd8, 0x39, 0xf3, 0x71, 0xd6, 0x8e, 0x03, 0xb4, 0x1a, 0xe7, 0xab, 0xfa, 0xb7, 0xc0, 0xd9,
	0x95, 0x69, 0x54, 0xdc, 0x85, 0x6c, 0x90, 0xec, 0xde, 0x49, 0x71, 0xef, 0xe4, 0x4b, 0x67, 0x6c,
	0xf6, 0x5b, 0x17, 0x34, 0x4e, 0x21, 0xac, 0x15, 0xa0, 0xaf, 0x0e, 0xb3, 0x7d, 0x5f, 0x46, 0xeb,
	0x4f, 0xc4, 0x1b, 0x6a, 0x52, 0x8d, 0x0f, 0xd0, 0xa2, 0x36, 0xd4, 0x80, 0xe3, 0x5c, 0x6b, 0xbd,
	0x7b, 0xce, 0x2e, 0x87, 0x20, 0xf6, 0x80, 0x30, 0x43, 0x59, 0xd1, 0x09, 0x68, 0x4d, 0x19, 0x34,
	0xca, 0x99, 0xe8, 0x3c, 0xc4, 0x1e, 0x7a, 0x25, 0x06, 0x1a, 0xd4, 0x04, 0x86, 0x7d, 0x06, 0x63,
	0x50, 0xd4, 0x70, 0x31, 0x76, 0xe4, 0x2b, 0x21, 0x2e, 0x4a, 0xc7, 0x0f, 0x15, 0xec, 0xa3, 0x55,
	0xa9, 0x44, 0x04, 0x5a, 0xf3, 0x31, 0xeb, 0xdb, 0x8f, 0xb8, 0xb1, 0xe0, 0x94, 0x36, 0xe7, 0x94,
	0x9e, 0x17, 0x5f, 0x78, 0x58, 0x9b, 0x41, 0x6c, 0x72, 0xfb, 0x33, 0x5a, 0x74, 0xfc, 0x70, 0x15,
	0x2d, 0x9f, 0x05, 0x27, 0xed, 0xce, 0xc9, 0x71, 0xfd, 0x3f, 0x5c, 0x43, 0xe8, 0x2c, 0x3c, 0xf5,
	0x83, 0x5e, 0xcf, 0xc6, 0x25, 0x5b, 0xec, 0x9c, 0x5c, 0x1c, 0x76, 0x3b, 0xed, 0x7a, 0x19, 0x23,
	0xb4, 0x74, 0x74, 0xd8, 0xe9, 0x06, 0xed, 0x7a, 0x05, 0xbf, 0x40, 0xff, 0x1f, 0xfa, 0x7e, 0x70,
	0x76, 0x1e, 0xb4, 0xeb, 0x0b, 0x9f, 0xfc, 0x1f, 0xbf, 0x17, 0x4a, 0xdf, 0x7e, 0x6e, 0x96, 0x2e,
	0xdf, 0x3f, 0xfb, 0xf6, 0xe5, 0x0d, 0x73, 0xcf, 0xeb, 0x91, 0x75, 0xf6, 0x95, 0x99, 0xa9, 0x04,
	0x3d, 0x58, 0x72, 0xac, 0xf7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x3c, 0x70, 0xf2, 0x3f,
	0x04, 0x00, 0x00,
}

func (this *FailoverSchemeSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverSchemeSpec)
	if !ok {
		that2, ok := that.(FailoverSchemeSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Primary.Equal(that1.Primary) {
		return false
	}
	if len(this.FailoverGroups) != len(that1.FailoverGroups) {
		return false
	}
	for i := range this.FailoverGroups {
		if !this.FailoverGroups[i].Equal(that1.FailoverGroups[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverSchemeSpec_FailoverEndpoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverSchemeSpec_FailoverEndpoints)
	if !ok {
		that2, ok := that.(FailoverSchemeSpec_FailoverEndpoints)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PriorityGroup) != len(that1.PriorityGroup) {
		return false
	}
	for i := range this.PriorityGroup {
		if !this.PriorityGroup[i].Equal(that1.PriorityGroup[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets)
	if !ok {
		that2, ok := that.(FailoverSchemeSpec_FailoverEndpoints_LocalityLbTargets)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Cluster != that1.Cluster {
		return false
	}
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if !this.LocalityWeight.Equal(that1.LocalityWeight) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverSchemeStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverSchemeStatus)
	if !ok {
		that2, ok := that.(FailoverSchemeStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if !this.ProcessingTime.Equal(that1.ProcessingTime) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}