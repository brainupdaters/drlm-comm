// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drlm.proto

package drlm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthType int32

const (
	AuthType_LOCAL AuthType = 0
)

var AuthType_name = map[int32]string{
	0: "LOCAL",
}

var AuthType_value = map[string]int32{
	"LOCAL": 0,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}

func (AuthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{0}
}

type UserLoginRequest struct {
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{0}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func (m *UserLoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserLoginResponse struct {
	Tkn                  string               `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
	TknExpiration        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=tkn_expiration,json=tknExpiration,proto3" json:"tkn_expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{1}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetTkn() string {
	if m != nil {
		return m.Tkn
	}
	return ""
}

func (m *UserLoginResponse) GetTknExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.TknExpiration
	}
	return nil
}

type UserTokenRenewRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenRenewRequest) Reset()         { *m = UserTokenRenewRequest{} }
func (m *UserTokenRenewRequest) String() string { return proto.CompactTextString(m) }
func (*UserTokenRenewRequest) ProtoMessage()    {}
func (*UserTokenRenewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{2}
}

func (m *UserTokenRenewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenRenewRequest.Unmarshal(m, b)
}
func (m *UserTokenRenewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenRenewRequest.Marshal(b, m, deterministic)
}
func (m *UserTokenRenewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenRenewRequest.Merge(m, src)
}
func (m *UserTokenRenewRequest) XXX_Size() int {
	return xxx_messageInfo_UserTokenRenewRequest.Size(m)
}
func (m *UserTokenRenewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenRenewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenRenewRequest proto.InternalMessageInfo

type UserTokenRenewResponse struct {
	Tkn                  string               `protobuf:"bytes,1,opt,name=tkn,proto3" json:"tkn,omitempty"`
	TknExpiration        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=tkn_expiration,json=tknExpiration,proto3" json:"tkn_expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserTokenRenewResponse) Reset()         { *m = UserTokenRenewResponse{} }
func (m *UserTokenRenewResponse) String() string { return proto.CompactTextString(m) }
func (*UserTokenRenewResponse) ProtoMessage()    {}
func (*UserTokenRenewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{3}
}

func (m *UserTokenRenewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenRenewResponse.Unmarshal(m, b)
}
func (m *UserTokenRenewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenRenewResponse.Marshal(b, m, deterministic)
}
func (m *UserTokenRenewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenRenewResponse.Merge(m, src)
}
func (m *UserTokenRenewResponse) XXX_Size() int {
	return xxx_messageInfo_UserTokenRenewResponse.Size(m)
}
func (m *UserTokenRenewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenRenewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenRenewResponse proto.InternalMessageInfo

func (m *UserTokenRenewResponse) GetTkn() string {
	if m != nil {
		return m.Tkn
	}
	return ""
}

func (m *UserTokenRenewResponse) GetTknExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.TknExpiration
	}
	return nil
}

type UserAddRequest struct {
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddRequest) Reset()         { *m = UserAddRequest{} }
func (m *UserAddRequest) String() string { return proto.CompactTextString(m) }
func (*UserAddRequest) ProtoMessage()    {}
func (*UserAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{4}
}

func (m *UserAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddRequest.Unmarshal(m, b)
}
func (m *UserAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddRequest.Marshal(b, m, deterministic)
}
func (m *UserAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddRequest.Merge(m, src)
}
func (m *UserAddRequest) XXX_Size() int {
	return xxx_messageInfo_UserAddRequest.Size(m)
}
func (m *UserAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddRequest proto.InternalMessageInfo

func (m *UserAddRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func (m *UserAddRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserAddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddResponse) Reset()         { *m = UserAddResponse{} }
func (m *UserAddResponse) String() string { return proto.CompactTextString(m) }
func (*UserAddResponse) ProtoMessage()    {}
func (*UserAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{5}
}

func (m *UserAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddResponse.Unmarshal(m, b)
}
func (m *UserAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddResponse.Marshal(b, m, deterministic)
}
func (m *UserAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddResponse.Merge(m, src)
}
func (m *UserAddResponse) XXX_Size() int {
	return xxx_messageInfo_UserAddResponse.Size(m)
}
func (m *UserAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddResponse proto.InternalMessageInfo

type UserDeleteRequest struct {
	Usr                  string   `protobuf:"bytes,2,opt,name=usr,proto3" json:"usr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{6}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

type UserDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteResponse) Reset()         { *m = UserDeleteResponse{} }
func (m *UserDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*UserDeleteResponse) ProtoMessage()    {}
func (*UserDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{7}
}

func (m *UserDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteResponse.Unmarshal(m, b)
}
func (m *UserDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteResponse.Marshal(b, m, deterministic)
}
func (m *UserDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteResponse.Merge(m, src)
}
func (m *UserDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_UserDeleteResponse.Size(m)
}
func (m *UserDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteResponse proto.InternalMessageInfo

type UserListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{8}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

type UserListResponse struct {
	Users                []*UserListResponse_User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{9}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*UserListResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserListResponse_User struct {
	Usr                  string               `protobuf:"bytes,1,opt,name=usr,proto3" json:"usr,omitempty"`
	AuthType             AuthType             `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3,enum=drlm.AuthType" json:"auth_type,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserListResponse_User) Reset()         { *m = UserListResponse_User{} }
func (m *UserListResponse_User) String() string { return proto.CompactTextString(m) }
func (*UserListResponse_User) ProtoMessage()    {}
func (*UserListResponse_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4bd9cd91f607bb1, []int{9, 0}
}

func (m *UserListResponse_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse_User.Unmarshal(m, b)
}
func (m *UserListResponse_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse_User.Marshal(b, m, deterministic)
}
func (m *UserListResponse_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse_User.Merge(m, src)
}
func (m *UserListResponse_User) XXX_Size() int {
	return xxx_messageInfo_UserListResponse_User.Size(m)
}
func (m *UserListResponse_User) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse_User.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse_User proto.InternalMessageInfo

func (m *UserListResponse_User) GetUsr() string {
	if m != nil {
		return m.Usr
	}
	return ""
}

func (m *UserListResponse_User) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_LOCAL
}

func (m *UserListResponse_User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *UserListResponse_User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("drlm.AuthType", AuthType_name, AuthType_value)
	proto.RegisterType((*UserLoginRequest)(nil), "drlm.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "drlm.UserLoginResponse")
	proto.RegisterType((*UserTokenRenewRequest)(nil), "drlm.UserTokenRenewRequest")
	proto.RegisterType((*UserTokenRenewResponse)(nil), "drlm.UserTokenRenewResponse")
	proto.RegisterType((*UserAddRequest)(nil), "drlm.UserAddRequest")
	proto.RegisterType((*UserAddResponse)(nil), "drlm.UserAddResponse")
	proto.RegisterType((*UserDeleteRequest)(nil), "drlm.UserDeleteRequest")
	proto.RegisterType((*UserDeleteResponse)(nil), "drlm.UserDeleteResponse")
	proto.RegisterType((*UserListRequest)(nil), "drlm.UserListRequest")
	proto.RegisterType((*UserListResponse)(nil), "drlm.UserListResponse")
	proto.RegisterType((*UserListResponse_User)(nil), "drlm.UserListResponse.User")
}

func init() { proto.RegisterFile("drlm.proto", fileDescriptor_a4bd9cd91f607bb1) }

var fileDescriptor_a4bd9cd91f607bb1 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0xda, 0xae, 0x36, 0x77, 0xb1, 0x76, 0x87, 0x6d, 0x37, 0x44, 0xc1, 0x25, 0x20, 0x2c,
	0x0a, 0x59, 0xac, 0x22, 0x8a, 0x20, 0x04, 0xd7, 0xb7, 0x2c, 0x42, 0xa8, 0xcf, 0x25, 0x6b, 0xae,
	0x6d, 0x48, 0x3b, 0x19, 0x33, 0x37, 0xd4, 0x7d, 0xf6, 0x67, 0xf9, 0x23, 0xfc, 0x4b, 0x92, 0xcc,
	0xe4, 0xc3, 0xb4, 0x2a, 0x3e, 0xf8, 0x36, 0x39, 0x73, 0xce, 0x9d, 0x93, 0x73, 0x0f, 0x40, 0x94,
	0x6d, 0xb6, 0xae, 0xc8, 0x52, 0x4a, 0xd9, 0xb0, 0x38, 0xdb, 0x8f, 0x56, 0x69, 0xba, 0xda, 0xe0,
	0x65, 0x89, 0xdd, 0xe4, 0x9f, 0x2f, 0x29, 0xde, 0xa2, 0xa4, 0x70, 0x2b, 0x14, 0xcd, 0x79, 0x09,
	0x93, 0x8f, 0x12, 0x33, 0x3f, 0x5d, 0xc5, 0x3c, 0xc0, 0x2f, 0x39, 0x4a, 0x62, 0x13, 0x18, 0xe4,
	0x32, 0xb3, 0xfa, 0xe7, 0xc6, 0x85, 0x19, 0x14, 0xc7, 0x02, 0x11, 0xbb, 0xc8, 0x1a, 0x28, 0x44,
	0xec, 0x22, 0x67, 0x0d, 0x27, 0x2d, 0x9d, 0x14, 0x29, 0x97, 0x58, 0xd0, 0x28, 0xe1, 0x96, 0xa1,
	0x68, 0x94, 0x70, 0xe6, 0xc1, 0x98, 0x12, 0xbe, 0xc4, 0xaf, 0x22, 0xce, 0x42, 0x8a, 0x53, 0x5e,
	0x4e, 0x3d, 0x9e, 0xdb, 0xae, 0x32, 0xe6, 0x56, 0xc6, 0xdc, 0x45, 0x65, 0x2c, 0xb8, 0x47, 0x09,
	0x7f, 0x5f, 0x0b, 0x9c, 0x33, 0x98, 0x16, 0x2f, 0x2d, 0xd2, 0x04, 0x79, 0x80, 0x1c, 0x77, 0xda,
	0xa6, 0xb3, 0x85, 0x59, 0xf7, 0xe2, 0x7f, 0xfa, 0x78, 0x01, 0xe3, 0xe2, 0x39, 0x2f, 0x8a, 0xfe,
	0x25, 0xa7, 0x13, 0xb8, 0x5f, 0xab, 0x94, 0x3b, 0xe7, 0xb1, 0x8a, 0xee, 0x0a, 0x37, 0x48, 0xf8,
	0xdb, 0x59, 0xce, 0x29, 0xb0, 0x36, 0x4d, 0x8b, 0xf5, 0x3c, 0x3f, 0x96, 0x54, 0xe5, 0xf0, 0xad,
	0xaf, 0x77, 0x58, 0x62, 0x3a, 0x82, 0x67, 0x70, 0x94, 0x4b, 0xcc, 0xa4, 0x65, 0x9c, 0x0f, 0x2e,
	0x8e, 0xe7, 0x0f, 0xdc, 0xb2, 0x1a, 0x5d, 0x5a, 0x09, 0x04, 0x8a, 0x69, 0x7f, 0x37, 0x60, 0x58,
	0x7c, 0x57, 0x5e, 0x8c, 0xe6, 0xbf, 0x9e, 0x82, 0x19, 0xe6, 0xb4, 0x5e, 0xd2, 0xad, 0xc0, 0xd2,
	0xe3, 0x78, 0x3e, 0x56, 0x13, 0xbd, 0x9c, 0xd6, 0x8b, 0x5b, 0x81, 0xc1, 0x28, 0xd4, 0x27, 0xf6,
	0x1a, 0xe0, 0x53, 0x86, 0x21, 0x61, 0xb4, 0x0c, 0xa9, 0xcc, 0xe2, 0xcf, 0x39, 0x9b, 0x9a, 0xed,
	0x51, 0x21, 0xcd, 0x45, 0x54, 0x49, 0x87, 0x7f, 0x97, 0x6a, 0xb6, 0x47, 0x4f, 0xa6, 0x30, 0xaa,
	0xbc, 0x30, 0x13, 0x8e, 0xfc, 0x0f, 0xef, 0x3c, 0x7f, 0xd2, 0x9b, 0xff, 0xe8, 0xc3, 0xf0, 0x2a,
	0xf0, 0xaf, 0xd9, 0x5b, 0x30, 0xeb, 0xc2, 0xb2, 0x59, 0x2b, 0x8e, 0x56, 0xf3, 0xed, 0xb3, 0x3d,
	0x5c, 0xc7, 0xde, 0x63, 0xd7, 0x6a, 0xfd, 0x4d, 0xdb, 0x58, 0x2b, 0xd3, 0xbd, 0x72, 0xda, 0x0f,
	0x0f, 0x5f, 0xd6, 0xe3, 0x5e, 0xc1, 0x5d, 0xdd, 0x0b, 0x76, 0xda, 0x50, 0x9b, 0x72, 0xd9, 0xd3,
	0x0e, 0x5a, 0x2b, 0x3d, 0x80, 0xa6, 0x17, 0xac, 0xe5, 0xf8, 0x97, 0x42, 0xd9, 0xd6, 0xfe, 0x45,
	0x3d, 0xe2, 0x0d, 0x8c, 0xaa, 0x26, 0xb0, 0x69, 0xb7, 0x19, 0x4a, 0x3e, 0x3b, 0x5c, 0x18, 0xa7,
	0x77, 0x73, 0xa7, 0xdc, 0xc3, 0xf3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xd9, 0x72, 0x8b,
	0x6d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DRLMClient is the client API for DRLM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DRLMClient interface {
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	UserTokenRenew(ctx context.Context, in *UserTokenRenewRequest, opts ...grpc.CallOption) (*UserTokenRenewResponse, error)
	UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
}

type dRLMClient struct {
	cc *grpc.ClientConn
}

func NewDRLMClient(cc *grpc.ClientConn) DRLMClient {
	return &dRLMClient{cc}
}

func (c *dRLMClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserTokenRenew(ctx context.Context, in *UserTokenRenewRequest, opts ...grpc.CallOption) (*UserTokenRenewResponse, error) {
	out := new(UserTokenRenewResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserTokenRenew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserAdd(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserAddResponse, error) {
	out := new(UserAddResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dRLMClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/drlm.DRLM/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DRLMServer is the server API for DRLM service.
type DRLMServer interface {
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	UserTokenRenew(context.Context, *UserTokenRenewRequest) (*UserTokenRenewResponse, error)
	UserAdd(context.Context, *UserAddRequest) (*UserAddResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
}

// UnimplementedDRLMServer can be embedded to have forward compatible implementations.
type UnimplementedDRLMServer struct {
}

func (*UnimplementedDRLMServer) UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedDRLMServer) UserTokenRenew(ctx context.Context, req *UserTokenRenewRequest) (*UserTokenRenewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserTokenRenew not implemented")
}
func (*UnimplementedDRLMServer) UserAdd(ctx context.Context, req *UserAddRequest) (*UserAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (*UnimplementedDRLMServer) UserDelete(ctx context.Context, req *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (*UnimplementedDRLMServer) UserList(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}

func RegisterDRLMServer(s *grpc.Server, srv DRLMServer) {
	s.RegisterService(&_DRLM_serviceDesc, srv)
}

func _DRLM_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserTokenRenew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserTokenRenew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserTokenRenew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserTokenRenew(ctx, req.(*UserTokenRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserAdd(ctx, req.(*UserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DRLM_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DRLMServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drlm.DRLM/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DRLMServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DRLM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drlm.DRLM",
	HandlerType: (*DRLMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _DRLM_UserLogin_Handler,
		},
		{
			MethodName: "UserTokenRenew",
			Handler:    _DRLM_UserTokenRenew_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _DRLM_UserAdd_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _DRLM_UserDelete_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _DRLM_UserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drlm.proto",
}
