// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: auth.proto

package radwallet

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Verified  bool              `protobuf:"varint,2,opt,name=verified,proto3" json:"verified,omitempty"`
	Exchanges []*ExchangeCommon `protobuf:"bytes,3,rep,name=exchanges,json=configured_exchanges,proto3" json:"exchanges,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Data) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Data) GetExchanges() []*ExchangeCommon {
	if x != nil {
		return x.Exchanges
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,json=access_token,proto3" json:"accessToken,omitempty"`
	TokenType   string `protobuf:"bytes,2,opt,name=tokenType,json=token_type,proto3" json:"tokenType,omitempty"`
	Data        *Data  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *LoginResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirebaseToken string `protobuf:"bytes,1,opt,name=firebase_token,proto3" json:"firebase_token,omitempty"`
	Email         string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetFirebaseToken() string {
	if x != nil {
		return x.FirebaseToken
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Verified bool   `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignupResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SignupResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SignupRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid  bool                 `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Claims *TokenResponse_Claim `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *TokenResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *TokenResponse) GetClaims() *TokenResponse_Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *TokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PEMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*PEMResponse_Key `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *PEMResponse) Reset() {
	*x = PEMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEMResponse) ProtoMessage() {}

func (x *PEMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEMResponse.ProtoReflect.Descriptor instead.
func (*PEMResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *PEMResponse) GetKeys() []*PEMResponse_Key {
	if x != nil {
		return x.Keys
	}
	return nil
}

type TokenResponse_Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub   string `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *TokenResponse_Claim) Reset() {
	*x = TokenResponse_Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse_Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse_Claim) ProtoMessage() {}

func (x *TokenResponse_Claim) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse_Claim.ProtoReflect.Descriptor instead.
func (*TokenResponse_Claim) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5, 0}
}

func (x *TokenResponse_Claim) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *TokenResponse_Claim) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PEMResponse_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E   []byte `protobuf:"bytes,1,opt,name=e,proto3" json:"e,omitempty"`
	Kty string `protobuf:"bytes,2,opt,name=kty,proto3" json:"kty,omitempty"`
	N   string `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	Alg string `protobuf:"bytes,4,opt,name=alg,proto3" json:"alg,omitempty"`
	Kid string `protobuf:"bytes,5,opt,name=kid,proto3" json:"kid,omitempty"`
	Use string `protobuf:"bytes,6,opt,name=use,proto3" json:"use,omitempty"`
}

func (x *PEMResponse_Key) Reset() {
	*x = PEMResponse_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PEMResponse_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PEMResponse_Key) ProtoMessage() {}

func (x *PEMResponse_Key) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PEMResponse_Key.ProtoReflect.Descriptor instead.
func (*PEMResponse_Key) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7, 0}
}

func (x *PEMResponse_Key) GetE() []byte {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *PEMResponse_Key) GetKty() string {
	if x != nil {
		return x.Kty
	}
	return ""
}

func (x *PEMResponse_Key) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *PEMResponse_Key) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *PEMResponse_Key) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *PEMResponse_Key) GetUse() string {
	if x != nil {
		return x.Use
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x61,
	0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x61,
	0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x14,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e,
	0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x52, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x96, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72,
	0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52,
	0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x1a, 0x2f, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x24, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb0,
	0x01, 0x0a, 0x0b, 0x50, 0x45, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72,
	0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x50, 0x45, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x1a, 0x69, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x74, 0x79, 0x12, 0x0c, 0x0a,
	0x01, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x6c, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73,
	0x65, 0x32, 0xe0, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x69, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61,
	0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72,
	0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12,
	0x20, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e, 0x72, 0x61, 0x64, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x3a, 0x01, 0x2a, 0x42, 0x4e, 0x0a, 0x11, 0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2e,
	0x72, 0x61, 0x64, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2d, 0x54, 0x65, 0x63, 0x68, 0x2f,
	0x72, 0x61, 0x64, 0x68, 0x61, 0x72, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x72, 0x61, 0x64, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_auth_proto_goTypes = []interface{}{
	(*Data)(nil),                // 0: radhark.radwallet.Data
	(*LoginResponse)(nil),       // 1: radhark.radwallet.LoginResponse
	(*LoginRequest)(nil),        // 2: radhark.radwallet.LoginRequest
	(*SignupResponse)(nil),      // 3: radhark.radwallet.SignupResponse
	(*SignupRequest)(nil),       // 4: radhark.radwallet.SignupRequest
	(*TokenResponse)(nil),       // 5: radhark.radwallet.TokenResponse
	(*TokenRequest)(nil),        // 6: radhark.radwallet.TokenRequest
	(*PEMResponse)(nil),         // 7: radhark.radwallet.PEMResponse
	(*TokenResponse_Claim)(nil), // 8: radhark.radwallet.TokenResponse.Claim
	(*PEMResponse_Key)(nil),     // 9: radhark.radwallet.PEMResponse.Key
	(*ExchangeCommon)(nil),      // 10: radhark.radwallet.ExchangeCommon
}
var file_auth_proto_depIdxs = []int32{
	10, // 0: radhark.radwallet.Data.exchanges:type_name -> radhark.radwallet.ExchangeCommon
	0,  // 1: radhark.radwallet.LoginResponse.data:type_name -> radhark.radwallet.Data
	8,  // 2: radhark.radwallet.TokenResponse.claims:type_name -> radhark.radwallet.TokenResponse.Claim
	9,  // 3: radhark.radwallet.PEMResponse.keys:type_name -> radhark.radwallet.PEMResponse.Key
	2,  // 4: radhark.radwallet.Auth.Login:input_type -> radhark.radwallet.LoginRequest
	4,  // 5: radhark.radwallet.Auth.Signup:input_type -> radhark.radwallet.SignupRequest
	1,  // 6: radhark.radwallet.Auth.Login:output_type -> radhark.radwallet.LoginResponse
	3,  // 7: radhark.radwallet.Auth.Signup:output_type -> radhark.radwallet.SignupResponse
	6,  // [6:8] is the sub-list for method output_type
	4,  // [4:6] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	file_exchange_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEMResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse_Claim); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PEMResponse_Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
