// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/manager/auth/manager_auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manager_auth_manager_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_manager_auth_manager_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_api_manager_auth_manager_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AuthRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId            uint32 `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	RoleKeyword       string `protobuf:"bytes,3,opt,name=roleKeyword,proto3" json:"roleKeyword,omitempty"`
	DepartmentId      uint32 `protobuf:"varint,4,opt,name=departmentId,proto3" json:"departmentId,omitempty"`
	DepartmentKeyword string `protobuf:"bytes,5,opt,name=departmentKeyword,proto3" json:"departmentKeyword,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_manager_auth_manager_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_manager_auth_manager_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_manager_auth_manager_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthReply) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthReply) GetRoleId() uint32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *AuthReply) GetRoleKeyword() string {
	if x != nil {
		return x.RoleKeyword
	}
	return ""
}

func (x *AuthReply) GetDepartmentId() uint32 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *AuthReply) GetDepartmentKeyword() string {
	if x != nil {
		return x.DepartmentKeyword
	}
	return ""
}

var File_api_manager_auth_manager_auth_proto protoreflect.FileDescriptor

var file_api_manager_auth_manager_auth_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x1a, 0x72, 0x18, 0x52, 0x03, 0x47,
	0x45, 0x54, 0x52, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x52, 0x03, 0x50, 0x55, 0x54, 0x52, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xaf, 0x01,
	0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72,
	0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_manager_auth_manager_auth_proto_rawDescOnce sync.Once
	file_api_manager_auth_manager_auth_proto_rawDescData = file_api_manager_auth_manager_auth_proto_rawDesc
)

func file_api_manager_auth_manager_auth_proto_rawDescGZIP() []byte {
	file_api_manager_auth_manager_auth_proto_rawDescOnce.Do(func() {
		file_api_manager_auth_manager_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_manager_auth_manager_auth_proto_rawDescData)
	})
	return file_api_manager_auth_manager_auth_proto_rawDescData
}

var file_api_manager_auth_manager_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_manager_auth_manager_auth_proto_goTypes = []interface{}{
	(*AuthRequest)(nil), // 0: manager_auth.AuthRequest
	(*AuthReply)(nil),   // 1: manager_auth.AuthReply
}
var file_api_manager_auth_manager_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_manager_auth_manager_auth_proto_init() }
func file_api_manager_auth_manager_auth_proto_init() {
	if File_api_manager_auth_manager_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_manager_auth_manager_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_api_manager_auth_manager_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
			RawDescriptor: file_api_manager_auth_manager_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_manager_auth_manager_auth_proto_goTypes,
		DependencyIndexes: file_api_manager_auth_manager_auth_proto_depIdxs,
		MessageInfos:      file_api_manager_auth_manager_auth_proto_msgTypes,
	}.Build()
	File_api_manager_auth_manager_auth_proto = out.File
	file_api_manager_auth_manager_auth_proto_rawDesc = nil
	file_api_manager_auth_manager_auth_proto_goTypes = nil
	file_api_manager_auth_manager_auth_proto_depIdxs = nil
}
