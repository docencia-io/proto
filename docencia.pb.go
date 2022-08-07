// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: docencia.proto

package proto

import (
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

// The ChangePassword request.
type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{0}
}

func (x *ChangePasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChangePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// The ChangePassword response.
type ChangePasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChangePasswordReply) Reset() {
	*x = ChangePasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReply) ProtoMessage() {}

func (x *ChangePasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReply.ProtoReflect.Descriptor instead.
func (*ChangePasswordReply) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{1}
}

func (x *ChangePasswordReply) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ChangePasswordReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The FindOrCreateUser request.
type FindOrCreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Provider string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *FindOrCreateUserRequest) Reset() {
	*x = FindOrCreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOrCreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOrCreateUserRequest) ProtoMessage() {}

func (x *FindOrCreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOrCreateUserRequest.ProtoReflect.Descriptor instead.
func (*FindOrCreateUserRequest) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{2}
}

func (x *FindOrCreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FindOrCreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOrCreateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *FindOrCreateUserRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

// The FindOrCreateUser response.
type FindOrCreateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id       int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *FindOrCreateUserReply) Reset() {
	*x = FindOrCreateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOrCreateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOrCreateUserReply) ProtoMessage() {}

func (x *FindOrCreateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOrCreateUserReply.ProtoReflect.Descriptor instead.
func (*FindOrCreateUserReply) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{3}
}

func (x *FindOrCreateUserReply) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *FindOrCreateUserReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FindOrCreateUserReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FindOrCreateUserReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FindOrCreateUserReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindOrCreateUserReply) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

// The InfoID request.
type InfoIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *InfoIDRequest) Reset() {
	*x = InfoIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoIDRequest) ProtoMessage() {}

func (x *InfoIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoIDRequest.ProtoReflect.Descriptor instead.
func (*InfoIDRequest) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{4}
}

func (x *InfoIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InfoIDRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// The ChangePassword response.
type InfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Role     string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	Info     []byte `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *InfoReply) Reset() {
	*x = InfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_docencia_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReply) ProtoMessage() {}

func (x *InfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_docencia_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReply.ProtoReflect.Descriptor instead.
func (*InfoReply) Descriptor() ([]byte, []int) {
	return file_docencia_proto_rawDescGZIP(), []int{5}
}

func (x *InfoReply) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *InfoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InfoReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InfoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoReply) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *InfoReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *InfoReply) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_docencia_proto protoreflect.FileDescriptor

var file_docencia_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x22, 0x49, 0x0a, 0x15, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7b, 0x0a, 0x17,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x9d, 0x01, 0x0a, 0x15, 0x46, 0x69,
	0x6e, 0x64, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x49, 0x6e, 0x66,
	0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0xa9, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xf3, 0x01, 0x0a,
	0x04, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x52, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x47,
	0x52, 0x50, 0x43, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x10, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x6f, 0x63, 0x65, 0x6e, 0x63, 0x69, 0x61, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_docencia_proto_rawDescOnce sync.Once
	file_docencia_proto_rawDescData = file_docencia_proto_rawDesc
)

func file_docencia_proto_rawDescGZIP() []byte {
	file_docencia_proto_rawDescOnce.Do(func() {
		file_docencia_proto_rawDescData = protoimpl.X.CompressGZIP(file_docencia_proto_rawDescData)
	})
	return file_docencia_proto_rawDescData
}

var file_docencia_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_docencia_proto_goTypes = []interface{}{
	(*ChangePasswordRequest)(nil),   // 0: coreGRPC.ChangePasswordRequest
	(*ChangePasswordReply)(nil),     // 1: coreGRPC.ChangePasswordReply
	(*FindOrCreateUserRequest)(nil), // 2: coreGRPC.FindOrCreateUserRequest
	(*FindOrCreateUserReply)(nil),   // 3: coreGRPC.FindOrCreateUserReply
	(*InfoIDRequest)(nil),           // 4: coreGRPC.InfoIDRequest
	(*InfoReply)(nil),               // 5: coreGRPC.InfoReply
}
var file_docencia_proto_depIdxs = []int32{
	0, // 0: coreGRPC.Core.ChangePassword:input_type -> coreGRPC.ChangePasswordRequest
	2, // 1: coreGRPC.Core.FindOrCreateUser:input_type -> coreGRPC.FindOrCreateUserRequest
	4, // 2: coreGRPC.Core.GetInfoByID:input_type -> coreGRPC.InfoIDRequest
	1, // 3: coreGRPC.Core.ChangePassword:output_type -> coreGRPC.ChangePasswordReply
	3, // 4: coreGRPC.Core.FindOrCreateUser:output_type -> coreGRPC.FindOrCreateUserReply
	5, // 5: coreGRPC.Core.GetInfoByID:output_type -> coreGRPC.InfoReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_docencia_proto_init() }
func file_docencia_proto_init() {
	if File_docencia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_docencia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_docencia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordReply); i {
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
		file_docencia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOrCreateUserRequest); i {
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
		file_docencia_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOrCreateUserReply); i {
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
		file_docencia_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoIDRequest); i {
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
		file_docencia_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReply); i {
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
			RawDescriptor: file_docencia_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_docencia_proto_goTypes,
		DependencyIndexes: file_docencia_proto_depIdxs,
		MessageInfos:      file_docencia_proto_msgTypes,
	}.Build()
	File_docencia_proto = out.File
	file_docencia_proto_rawDesc = nil
	file_docencia_proto_goTypes = nil
	file_docencia_proto_depIdxs = nil
}
