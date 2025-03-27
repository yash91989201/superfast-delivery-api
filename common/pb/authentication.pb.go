// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: authentication.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRole int32

const (
	AuthRole_CUSTOMER         AuthRole = 0
	AuthRole_DELIVERY_PARTNER AuthRole = 1
	AuthRole_VENDOR           AuthRole = 2
	AuthRole_ADMIN            AuthRole = 3
)

// Enum value maps for AuthRole.
var (
	AuthRole_name = map[int32]string{
		0: "CUSTOMER",
		1: "DELIVERY_PARTNER",
		2: "VENDOR",
		3: "ADMIN",
	}
	AuthRole_value = map[string]int32{
		"CUSTOMER":         0,
		"DELIVERY_PARTNER": 1,
		"VENDOR":           2,
		"ADMIN":            3,
	}
)

func (x AuthRole) Enum() *AuthRole {
	p := new(AuthRole)
	*p = x
	return p
}

func (x AuthRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthRole) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_proto_enumTypes[0].Descriptor()
}

func (AuthRole) Type() protoreflect.EnumType {
	return &file_authentication_proto_enumTypes[0]
}

func (x AuthRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthRole.Descriptor instead.
func (AuthRole) EnumDescriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

type Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         *string                `protobuf:"bytes,2,opt,name=email,proto3,oneof" json:"email,omitempty"`
	EmailVerified bool                   `protobuf:"varint,3,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Phone         *string                `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	AuthRole      AuthRole               `protobuf:"varint,5,opt,name=auth_role,json=authRole,proto3,enum=pb.AuthRole" json:"auth_role,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_authentication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Auth) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Auth) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *Auth) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *Auth) GetAuthRole() AuthRole {
	if x != nil {
		return x.AuthRole
	}
	return AuthRole_CUSTOMER
}

func (x *Auth) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Auth) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Auth) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type Session struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken          string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=access_token_expires_at,json=accessTokenExpiresAt,proto3" json:"access_token_expires_at,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Session) Reset() {
	*x = Session{}
	mi := &file_authentication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Session) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

type SignInWithEmailReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	AuthRole      AuthRole               `protobuf:"varint,2,opt,name=auth_role,json=authRole,proto3,enum=pb.AuthRole" json:"auth_role,omitempty"`
	Otp           *string                `protobuf:"bytes,3,opt,name=otp,proto3,oneof" json:"otp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInWithEmailReq) Reset() {
	*x = SignInWithEmailReq{}
	mi := &file_authentication_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInWithEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithEmailReq) ProtoMessage() {}

func (x *SignInWithEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithEmailReq.ProtoReflect.Descriptor instead.
func (*SignInWithEmailReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *SignInWithEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInWithEmailReq) GetAuthRole() AuthRole {
	if x != nil {
		return x.AuthRole
	}
	return AuthRole_CUSTOMER
}

func (x *SignInWithEmailReq) GetOtp() string {
	if x != nil && x.Otp != nil {
		return *x.Otp
	}
	return ""
}

type SignInWithPhoneReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	AuthRole      AuthRole               `protobuf:"varint,2,opt,name=auth_role,json=authRole,proto3,enum=pb.AuthRole" json:"auth_role,omitempty"`
	Otp           *string                `protobuf:"bytes,3,opt,name=otp,proto3,oneof" json:"otp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInWithPhoneReq) Reset() {
	*x = SignInWithPhoneReq{}
	mi := &file_authentication_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInWithPhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithPhoneReq) ProtoMessage() {}

func (x *SignInWithPhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithPhoneReq.ProtoReflect.Descriptor instead.
func (*SignInWithPhoneReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *SignInWithPhoneReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignInWithPhoneReq) GetAuthRole() AuthRole {
	if x != nil {
		return x.AuthRole
	}
	return AuthRole_CUSTOMER
}

func (x *SignInWithPhoneReq) GetOtp() string {
	if x != nil && x.Otp != nil {
		return *x.Otp
	}
	return ""
}

type SignInWithGoogleReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdToken       string                 `protobuf:"bytes,1,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	AuthRole      AuthRole               `protobuf:"varint,2,opt,name=auth_role,json=authRole,proto3,enum=pb.AuthRole" json:"auth_role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInWithGoogleReq) Reset() {
	*x = SignInWithGoogleReq{}
	mi := &file_authentication_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInWithGoogleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithGoogleReq) ProtoMessage() {}

func (x *SignInWithGoogleReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithGoogleReq.ProtoReflect.Descriptor instead.
func (*SignInWithGoogleReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *SignInWithGoogleReq) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *SignInWithGoogleReq) GetAuthRole() AuthRole {
	if x != nil {
		return x.AuthRole
	}
	return AuthRole_CUSTOMER
}

type SignInRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Auth          *Auth                  `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	AccessToken   string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInRes) Reset() {
	*x = SignInRes{}
	mi := &file_authentication_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRes) ProtoMessage() {}

func (x *SignInRes) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRes.ProtoReflect.Descriptor instead.
func (*SignInRes) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{5}
}

func (x *SignInRes) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SignInRes) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshAccessTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshAccessTokenReq) Reset() {
	*x = RefreshAccessTokenReq{}
	mi := &file_authentication_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshAccessTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenReq) ProtoMessage() {}

func (x *RefreshAccessTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshAccessTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LogOutReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogOutReq) Reset() {
	*x = LogOutReq{}
	mi := &file_authentication_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogOutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOutReq) ProtoMessage() {}

func (x *LogOutReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogOutReq.ProtoReflect.Descriptor instead.
func (*LogOutReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{7}
}

func (x *LogOutReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type GetAuthByIdReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthByIdReq) Reset() {
	*x = GetAuthByIdReq{}
	mi := &file_authentication_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthByIdReq) ProtoMessage() {}

func (x *GetAuthByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthByIdReq.ProtoReflect.Descriptor instead.
func (*GetAuthByIdReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{8}
}

func (x *GetAuthByIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuthReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         *string                `protobuf:"bytes,1,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone         *string                `protobuf:"bytes,2,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthReq) Reset() {
	*x = GetAuthReq{}
	mi := &file_authentication_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthReq) ProtoMessage() {}

func (x *GetAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthReq.ProtoReflect.Descriptor instead.
func (*GetAuthReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{9}
}

func (x *GetAuthReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *GetAuthReq) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

type ValidateSessionReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateSessionReq) Reset() {
	*x = ValidateSessionReq{}
	mi := &file_authentication_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateSessionReq) ProtoMessage() {}

func (x *ValidateSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateSessionReq.ProtoReflect.Descriptor instead.
func (*ValidateSessionReq) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{10}
}

func (x *ValidateSessionReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ValidateSessionRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Auth          *Auth                  `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	SessionId     string                 `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateSessionRes) Reset() {
	*x = ValidateSessionRes{}
	mi := &file_authentication_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateSessionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateSessionRes) ProtoMessage() {}

func (x *ValidateSessionRes) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateSessionRes.ProtoReflect.Descriptor instead.
func (*ValidateSessionRes) Descriptor() ([]byte, []int) {
	return file_authentication_proto_rawDescGZIP(), []int{11}
}

func (x *ValidateSessionRes) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ValidateSessionRes) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

var File_authentication_proto protoreflect.FileDescriptor

var file_authentication_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x04,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x51, 0x0a, 0x17, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15,
	0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6f,
	0x74, 0x70, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6f, 0x74, 0x70, 0x22, 0x74, 0x0a,
	0x12, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x6f, 0x74, 0x70, 0x22, 0x5b, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74,
	0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65,
	0x22, 0x71, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x2a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x56, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x37, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x51, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x2a, 0x45, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x45, 0x4e, 0x44, 0x4f, 0x52, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x32, 0xd4, 0x03, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57,
	0x69, 0x74, 0x68, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x28, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x73, 0x68, 0x39, 0x31, 0x39, 0x38, 0x39, 0x32, 0x30, 0x31, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x66, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_authentication_proto_rawDescOnce sync.Once
	file_authentication_proto_rawDescData []byte
)

func file_authentication_proto_rawDescGZIP() []byte {
	file_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authentication_proto_rawDesc), len(file_authentication_proto_rawDesc)))
	})
	return file_authentication_proto_rawDescData
}

var file_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_authentication_proto_goTypes = []any{
	(AuthRole)(0),                 // 0: pb.AuthRole
	(*Auth)(nil),                  // 1: pb.Auth
	(*Session)(nil),               // 2: pb.Session
	(*SignInWithEmailReq)(nil),    // 3: pb.SignInWithEmailReq
	(*SignInWithPhoneReq)(nil),    // 4: pb.SignInWithPhoneReq
	(*SignInWithGoogleReq)(nil),   // 5: pb.SignInWithGoogleReq
	(*SignInRes)(nil),             // 6: pb.SignInRes
	(*RefreshAccessTokenReq)(nil), // 7: pb.RefreshAccessTokenReq
	(*LogOutReq)(nil),             // 8: pb.LogOutReq
	(*GetAuthByIdReq)(nil),        // 9: pb.GetAuthByIdReq
	(*GetAuthReq)(nil),            // 10: pb.GetAuthReq
	(*ValidateSessionReq)(nil),    // 11: pb.ValidateSessionReq
	(*ValidateSessionRes)(nil),    // 12: pb.ValidateSessionRes
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_authentication_proto_depIdxs = []int32{
	0,  // 0: pb.Auth.auth_role:type_name -> pb.AuthRole
	13, // 1: pb.Auth.created_at:type_name -> google.protobuf.Timestamp
	13, // 2: pb.Auth.updated_at:type_name -> google.protobuf.Timestamp
	13, // 3: pb.Auth.deleted_at:type_name -> google.protobuf.Timestamp
	13, // 4: pb.Session.access_token_expires_at:type_name -> google.protobuf.Timestamp
	0,  // 5: pb.SignInWithEmailReq.auth_role:type_name -> pb.AuthRole
	0,  // 6: pb.SignInWithPhoneReq.auth_role:type_name -> pb.AuthRole
	0,  // 7: pb.SignInWithGoogleReq.auth_role:type_name -> pb.AuthRole
	1,  // 8: pb.SignInRes.auth:type_name -> pb.Auth
	1,  // 9: pb.ValidateSessionRes.auth:type_name -> pb.Auth
	3,  // 10: pb.AuthenticationService.SignInWithEmail:input_type -> pb.SignInWithEmailReq
	4,  // 11: pb.AuthenticationService.SignInWithPhone:input_type -> pb.SignInWithPhoneReq
	5,  // 12: pb.AuthenticationService.SignInWithGoogle:input_type -> pb.SignInWithGoogleReq
	9,  // 13: pb.AuthenticationService.GetAuthById:input_type -> pb.GetAuthByIdReq
	10, // 14: pb.AuthenticationService.GetAuth:input_type -> pb.GetAuthReq
	7,  // 15: pb.AuthenticationService.RefreshAccessToken:input_type -> pb.RefreshAccessTokenReq
	8,  // 16: pb.AuthenticationService.LogOut:input_type -> pb.LogOutReq
	11, // 17: pb.AuthenticationService.ValidateSession:input_type -> pb.ValidateSessionReq
	6,  // 18: pb.AuthenticationService.SignInWithEmail:output_type -> pb.SignInRes
	6,  // 19: pb.AuthenticationService.SignInWithPhone:output_type -> pb.SignInRes
	6,  // 20: pb.AuthenticationService.SignInWithGoogle:output_type -> pb.SignInRes
	1,  // 21: pb.AuthenticationService.GetAuthById:output_type -> pb.Auth
	1,  // 22: pb.AuthenticationService.GetAuth:output_type -> pb.Auth
	6,  // 23: pb.AuthenticationService.RefreshAccessToken:output_type -> pb.SignInRes
	6,  // 24: pb.AuthenticationService.LogOut:output_type -> pb.SignInRes
	12, // 25: pb.AuthenticationService.ValidateSession:output_type -> pb.ValidateSessionRes
	18, // [18:26] is the sub-list for method output_type
	10, // [10:18] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_authentication_proto_init() }
func file_authentication_proto_init() {
	if File_authentication_proto != nil {
		return
	}
	file_authentication_proto_msgTypes[0].OneofWrappers = []any{}
	file_authentication_proto_msgTypes[2].OneofWrappers = []any{}
	file_authentication_proto_msgTypes[3].OneofWrappers = []any{}
	file_authentication_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authentication_proto_rawDesc), len(file_authentication_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_proto_depIdxs,
		EnumInfos:         file_authentication_proto_enumTypes,
		MessageInfos:      file_authentication_proto_msgTypes,
	}.Build()
	File_authentication_proto = out.File
	file_authentication_proto_goTypes = nil
	file_authentication_proto_depIdxs = nil
}
