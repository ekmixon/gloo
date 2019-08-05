// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Optional, if not set, Gloo will try to use the default AWS secret specified by environment variables.
	// If a secret is not provided, the environment must specify both the AWS access key and secret.
	// The environment variables used to indicate the AWS account can be:
	// - for the access key: "AWS_ACCESS_KEY_ID" or "AWS_ACCESS_KEY"
	// - for the secret: "AWS_SECRET_ACCESS_KEY" or "AWS_SECRET_KEY"
	// If set, a [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	// Gloo will create an EC2 API client with this credential. You may choose to use a credential with limited access
	// in conjunction with a list of Roles, specified by their Amazon Resource Number (ARN).
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// Optional, Amazon Resource Number (ARN) referring to IAM Role that should be assumed when the Upstream
	// queries for eligible EC2 instances.
	// If provided, Gloo will create an EC2 API client with the provided role. If not provided, Gloo will not assume
	// a role.
	RoleArn string `protobuf:"bytes,7,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	// deprecated: use role_arn. If you do use this field, only the first element will be read
	RoleArns []string `protobuf:"bytes,6,rep,name=role_arns,json=roleArns,proto3" json:"role_arns,omitempty"`
	// List of tag filters for selecting instances
	// An instance must match all the filters in order to be selected
	// Filter keys are not case-sensitive
	Filters []*TagFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// If set, will use the EC2 public IP address. Defaults to the private IP address.
	PublicIp bool `protobuf:"varint,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	// If set, will use this port on EC2 instances. Defaults to port 80.
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetRoleArn() string {
	if m != nil {
		return m.RoleArn
	}
	return ""
}

func (m *UpstreamSpec) GetRoleArns() []string {
	if m != nil {
		return m.RoleArns
	}
	return nil
}

func (m *UpstreamSpec) GetFilters() []*TagFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *UpstreamSpec) GetPublicIp() bool {
	if m != nil {
		return m.PublicIp
	}
	return false
}

func (m *UpstreamSpec) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type TagFilter struct {
	// Types that are valid to be assigned to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec                 isTagFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter.Unmarshal(m, b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return xxx_messageInfo_TagFilter.Size(m)
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

type isTagFilter_Spec interface {
	isTagFilter_Spec()
	Equal(interface{}) bool
}

type TagFilter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}
type TagFilter_KvPair_ struct {
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof"`
}

func (*TagFilter_Key) isTagFilter_Spec()     {}
func (*TagFilter_KvPair_) isTagFilter_Spec() {}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TagFilter) GetKey() string {
	if x, ok := m.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := m.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TagFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TagFilter_OneofMarshaler, _TagFilter_OneofUnmarshaler, _TagFilter_OneofSizer, []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
}

func _TagFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TagFilter)
	// spec
	switch x := m.Spec.(type) {
	case *TagFilter_Key:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Key)
	case *TagFilter_KvPair_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KvPair); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TagFilter.Spec has unexpected type %T", x)
	}
	return nil
}

func _TagFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TagFilter)
	switch tag {
	case 1: // spec.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Spec = &TagFilter_Key{x}
		return true, err
	case 2: // spec.kv_pair
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TagFilter_KvPair)
		err := b.DecodeMessage(msg)
		m.Spec = &TagFilter_KvPair_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TagFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TagFilter)
	// spec
	switch x := m.Spec.(type) {
	case *TagFilter_Key:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Key)))
		n += len(x.Key)
	case *TagFilter_KvPair_:
		s := proto.Size(x.KvPair)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TagFilter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter_KvPair) Reset()         { *m = TagFilter_KvPair{} }
func (m *TagFilter_KvPair) String() string { return proto.CompactTextString(m) }
func (*TagFilter_KvPair) ProtoMessage()    {}
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1, 0}
}
func (m *TagFilter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter_KvPair.Unmarshal(m, b)
}
func (m *TagFilter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter_KvPair.Marshal(b, m, deterministic)
}
func (m *TagFilter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter_KvPair.Merge(m, src)
}
func (m *TagFilter_KvPair) XXX_Size() int {
	return xxx_messageInfo_TagFilter_KvPair.Size(m)
}
func (m *TagFilter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter_KvPair proto.InternalMessageInfo

func (m *TagFilter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TagFilter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*TagFilter)(nil), "aws_ec2.plugins.gloo.solo.io.TagFilter")
	proto.RegisterType((*TagFilter_KvPair)(nil), "aws_ec2.plugins.gloo.solo.io.TagFilter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto", fileDescriptor_fc1fd6f1173c4563)
}

var fileDescriptor_fc1fd6f1173c4563 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9b, 0xd4, 0x89, 0xb7, 0x20, 0xa1, 0x55, 0x85, 0xdc, 0x80, 0x90, 0xd5, 0x0b, 0x3e,
	0xc0, 0x1a, 0xc2, 0x85, 0x23, 0xed, 0xa1, 0x6a, 0xe1, 0x82, 0x16, 0xb8, 0x70, 0xb1, 0x9c, 0xd5,
	0x78, 0x59, 0xec, 0x66, 0x56, 0xb3, 0xeb, 0x20, 0xfe, 0x81, 0xcf, 0xe0, 0xc0, 0x77, 0xf1, 0x25,
	0xc8, 0xbb, 0x49, 0xc5, 0x01, 0xaa, 0x9c, 0x66, 0xde, 0xcc, 0x9b, 0xf7, 0x46, 0xa3, 0x61, 0x6f,
	0xb5, 0xf1, 0x5f, 0x86, 0x95, 0x50, 0x78, 0x53, 0x39, 0xec, 0xf1, 0xb9, 0xc1, 0x4a, 0xf7, 0x88,
	0x95, 0x25, 0xfc, 0x0a, 0xca, 0xbb, 0x88, 0x1a, 0x6b, 0xaa, 0xcd, 0xcb, 0xca, 0xf6, 0x83, 0x36,
	0x6b, 0x57, 0x35, 0xdf, 0x5c, 0x05, 0x6a, 0x39, 0xc6, 0x1a, 0xd4, 0x52, 0x58, 0x42, 0x8f, 0xfc,
	0xf1, 0x2d, 0x8c, 0x34, 0x31, 0x8e, 0x8a, 0x51, 0x55, 0x18, 0x5c, 0x9c, 0x68, 0xd4, 0x18, 0x88,
	0xd5, 0x98, 0xc5, 0x99, 0xc5, 0xb3, 0x7f, 0xf8, 0x87, 0xd8, 0x19, 0xbf, 0x73, 0x25, 0x68, 0x23,
	0xfb, 0xec, 0xc7, 0x21, 0xbb, 0xf7, 0xc9, 0x3a, 0x4f, 0xd0, 0xdc, 0x7c, 0xb0, 0xa0, 0xf8, 0x43,
	0x96, 0x12, 0x68, 0x83, 0xeb, 0x3c, 0x29, 0x92, 0x32, 0x93, 0x5b, 0xc4, 0x5f, 0x33, 0xe6, 0x40,
	0x11, 0xf8, 0x9a, 0xa0, 0xcd, 0x0f, 0x8b, 0xa4, 0x3c, 0x5e, 0x9e, 0x0a, 0x85, 0x04, 0xbb, 0x7d,
	0x84, 0x04, 0x87, 0x03, 0x29, 0x90, 0xd0, 0xca, 0x2c, 0x92, 0x25, 0xb4, 0xfc, 0x94, 0xcd, 0x09,
	0x7b, 0xa8, 0x1b, 0x5a, 0xe7, 0xb3, 0xa0, 0x39, 0x1b, 0xf1, 0x39, 0xad, 0xf9, 0x23, 0x96, 0xed,
	0x5a, 0x2e, 0x4f, 0x8b, 0x49, 0x99, 0xc9, 0xf9, 0xb6, 0xe7, 0xf8, 0x39, 0x9b, 0xb5, 0xa6, 0xf7,
	0x40, 0x2e, 0x9f, 0x14, 0x93, 0xf2, 0x78, 0xf9, 0x54, 0xdc, 0x75, 0x0e, 0xf1, 0xb1, 0xd1, 0x97,
	0x81, 0x2f, 0x77, 0x73, 0xa3, 0xbe, 0x1d, 0x56, 0xbd, 0x51, 0xb5, 0xb1, 0xf9, 0xb4, 0x48, 0xca,
	0xb9, 0x9c, 0xc7, 0xc2, 0xb5, 0xe5, 0x9c, 0x4d, 0x2d, 0x92, 0xcf, 0x8f, 0x8a, 0xa4, 0xbc, 0x2f,
	0x43, 0x7e, 0xf6, 0x33, 0x61, 0xd9, 0xad, 0x0e, 0xe7, 0x6c, 0xd2, 0xc1, 0xf7, 0x78, 0x88, 0xab,
	0x03, 0x39, 0x02, 0x7e, 0xcd, 0x66, 0xdd, 0xa6, 0xb6, 0x8d, 0xa1, 0xed, 0x11, 0xc4, 0x9e, 0x5b,
	0x89, 0x77, 0x9b, 0xf7, 0x8d, 0xa1, 0xab, 0x03, 0x99, 0x76, 0x21, 0x5b, 0xbc, 0x60, 0x69, 0xac,
	0xf1, 0x07, 0x7f, 0x19, 0x45, 0x9b, 0x13, 0x76, 0xb4, 0x69, 0xfa, 0x01, 0x82, 0x49, 0x26, 0x23,
	0xb8, 0x48, 0xd9, 0xd4, 0x59, 0x50, 0x17, 0x97, 0xbf, 0x7e, 0x3f, 0x49, 0x3e, 0xbf, 0xd9, 0xef,
	0xd3, 0x6c, 0xa7, 0xff, 0xf3, 0x6d, 0xab, 0x34, 0x3c, 0xc1, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x61, 0x9f, 0x3b, 0xc0, 0xb4, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if this.RoleArn != that1.RoleArn {
		return false
	}
	if len(this.RoleArns) != len(that1.RoleArns) {
		return false
	}
	for i := range this.RoleArns {
		if this.RoleArns[i] != that1.RoleArns[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if this.PublicIp != that1.PublicIp {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
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
	if that1.Spec == nil {
		if this.Spec != nil {
			return false
		}
	} else if this.Spec == nil {
		return false
	} else if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_Key)
	if !ok {
		that2, ok := that.(TagFilter_Key)
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
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *TagFilter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair_)
	if !ok {
		that2, ok := that.(TagFilter_KvPair_)
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
	if !this.KvPair.Equal(that1.KvPair) {
		return false
	}
	return true
}
func (this *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
