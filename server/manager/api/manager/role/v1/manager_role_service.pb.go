// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/manager/role/manager_role_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_manager_role_manager_role_service_proto protoreflect.FileDescriptor

var file_api_manager_role_manager_role_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x09,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x73, 0x12, 0x32, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x69, 0x64, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x8b, 0x01, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x32, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0xa4, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01,
	0x2a, 0x1a, 0x1b, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x88,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x30, 0x0a, 0x1b, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_api_manager_role_manager_role_service_proto_goTypes = []interface{}{
	(*GetRoleMenuIdsRequest)(nil),   // 0: manager.api.manager.role.v1.GetRoleMenuIdsRequest
	(*ListRoleRequest)(nil),         // 1: manager.api.manager.role.v1.ListRoleRequest
	(*CreateRoleRequest)(nil),       // 2: manager.api.manager.role.v1.CreateRoleRequest
	(*UpdateRoleRequest)(nil),       // 3: manager.api.manager.role.v1.UpdateRoleRequest
	(*UpdateRoleMenuRequest)(nil),   // 4: manager.api.manager.role.v1.UpdateRoleMenuRequest
	(*UpdateRoleStatusRequest)(nil), // 5: manager.api.manager.role.v1.UpdateRoleStatusRequest
	(*DeleteRoleRequest)(nil),       // 6: manager.api.manager.role.v1.DeleteRoleRequest
	(*GetRoleRequest)(nil),          // 7: manager.api.manager.role.v1.GetRoleRequest
	(*GetRoleMenuIdsReply)(nil),     // 8: manager.api.manager.role.v1.GetRoleMenuIdsReply
	(*ListRoleReply)(nil),           // 9: manager.api.manager.role.v1.ListRoleReply
	(*CreateRoleReply)(nil),         // 10: manager.api.manager.role.v1.CreateRoleReply
	(*UpdateRoleReply)(nil),         // 11: manager.api.manager.role.v1.UpdateRoleReply
	(*UpdateRoleMenuReply)(nil),     // 12: manager.api.manager.role.v1.UpdateRoleMenuReply
	(*UpdateRoleStatusReply)(nil),   // 13: manager.api.manager.role.v1.UpdateRoleStatusReply
	(*DeleteRoleReply)(nil),         // 14: manager.api.manager.role.v1.DeleteRoleReply
	(*GetRoleReply)(nil),            // 15: manager.api.manager.role.v1.GetRoleReply
}
var file_api_manager_role_manager_role_service_proto_depIdxs = []int32{
	0,  // 0: manager.api.manager.role.v1.Role.GetRoleMenuIds:input_type -> manager.api.manager.role.v1.GetRoleMenuIdsRequest
	1,  // 1: manager.api.manager.role.v1.Role.ListRole:input_type -> manager.api.manager.role.v1.ListRoleRequest
	2,  // 2: manager.api.manager.role.v1.Role.CreateRole:input_type -> manager.api.manager.role.v1.CreateRoleRequest
	3,  // 3: manager.api.manager.role.v1.Role.UpdateRole:input_type -> manager.api.manager.role.v1.UpdateRoleRequest
	4,  // 4: manager.api.manager.role.v1.Role.UpdateRoleMenu:input_type -> manager.api.manager.role.v1.UpdateRoleMenuRequest
	5,  // 5: manager.api.manager.role.v1.Role.UpdateRoleStatus:input_type -> manager.api.manager.role.v1.UpdateRoleStatusRequest
	6,  // 6: manager.api.manager.role.v1.Role.DeleteRole:input_type -> manager.api.manager.role.v1.DeleteRoleRequest
	7,  // 7: manager.api.manager.role.v1.Role.GetRole:input_type -> manager.api.manager.role.v1.GetRoleRequest
	8,  // 8: manager.api.manager.role.v1.Role.GetRoleMenuIds:output_type -> manager.api.manager.role.v1.GetRoleMenuIdsReply
	9,  // 9: manager.api.manager.role.v1.Role.ListRole:output_type -> manager.api.manager.role.v1.ListRoleReply
	10, // 10: manager.api.manager.role.v1.Role.CreateRole:output_type -> manager.api.manager.role.v1.CreateRoleReply
	11, // 11: manager.api.manager.role.v1.Role.UpdateRole:output_type -> manager.api.manager.role.v1.UpdateRoleReply
	12, // 12: manager.api.manager.role.v1.Role.UpdateRoleMenu:output_type -> manager.api.manager.role.v1.UpdateRoleMenuReply
	13, // 13: manager.api.manager.role.v1.Role.UpdateRoleStatus:output_type -> manager.api.manager.role.v1.UpdateRoleStatusReply
	14, // 14: manager.api.manager.role.v1.Role.DeleteRole:output_type -> manager.api.manager.role.v1.DeleteRoleReply
	15, // 15: manager.api.manager.role.v1.Role.GetRole:output_type -> manager.api.manager.role.v1.GetRoleReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_manager_role_manager_role_service_proto_init() }
func file_api_manager_role_manager_role_service_proto_init() {
	if File_api_manager_role_manager_role_service_proto != nil {
		return
	}
	file_api_manager_role_manager_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_manager_role_manager_role_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_manager_role_manager_role_service_proto_goTypes,
		DependencyIndexes: file_api_manager_role_manager_role_service_proto_depIdxs,
	}.Build()
	File_api_manager_role_manager_role_service_proto = out.File
	file_api_manager_role_manager_role_service_proto_rawDesc = nil
	file_api_manager_role_manager_role_service_proto_goTypes = nil
	file_api_manager_role_manager_role_service_proto_depIdxs = nil
}
