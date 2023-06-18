// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: icon/v1/icon_service.proto

package iconv1

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

type Icon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link      string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Kindof    string `protobuf:"bytes,3,opt,name=kindof,proto3" json:"kindof,omitempty"`
	Owneruuid string `protobuf:"bytes,4,opt,name=owneruuid,proto3" json:"owneruuid,omitempty"`
}

func (x *Icon) Reset() {
	*x = Icon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icon_v1_icon_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Icon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Icon) ProtoMessage() {}

func (x *Icon) ProtoReflect() protoreflect.Message {
	mi := &file_icon_v1_icon_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Icon.ProtoReflect.Descriptor instead.
func (*Icon) Descriptor() ([]byte, []int) {
	return file_icon_v1_icon_service_proto_rawDescGZIP(), []int{0}
}

func (x *Icon) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Icon) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Icon) GetKindof() string {
	if x != nil {
		return x.Kindof
	}
	return ""
}

func (x *Icon) GetOwneruuid() string {
	if x != nil {
		return x.Owneruuid
	}
	return ""
}

type GetIconRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetIconRequest) Reset() {
	*x = GetIconRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icon_v1_icon_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIconRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIconRequest) ProtoMessage() {}

func (x *GetIconRequest) ProtoReflect() protoreflect.Message {
	mi := &file_icon_v1_icon_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIconRequest.ProtoReflect.Descriptor instead.
func (*GetIconRequest) Descriptor() ([]byte, []int) {
	return file_icon_v1_icon_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetIconRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetIconResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon *Icon `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *GetIconResponse) Reset() {
	*x = GetIconResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icon_v1_icon_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIconResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIconResponse) ProtoMessage() {}

func (x *GetIconResponse) ProtoReflect() protoreflect.Message {
	mi := &file_icon_v1_icon_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIconResponse.ProtoReflect.Descriptor instead.
func (*GetIconResponse) Descriptor() ([]byte, []int) {
	return file_icon_v1_icon_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetIconResponse) GetIcon() *Icon {
	if x != nil {
		return x.Icon
	}
	return nil
}

type DeleteIconRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Owneruuid string `protobuf:"bytes,2,opt,name=owneruuid,proto3" json:"owneruuid,omitempty"`
}

func (x *DeleteIconRequest) Reset() {
	*x = DeleteIconRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icon_v1_icon_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIconRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIconRequest) ProtoMessage() {}

func (x *DeleteIconRequest) ProtoReflect() protoreflect.Message {
	mi := &file_icon_v1_icon_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIconRequest.ProtoReflect.Descriptor instead.
func (*DeleteIconRequest) Descriptor() ([]byte, []int) {
	return file_icon_v1_icon_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteIconRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DeleteIconRequest) GetOwneruuid() string {
	if x != nil {
		return x.Owneruuid
	}
	return ""
}

type DeleteIconResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteIconResponse) Reset() {
	*x = DeleteIconResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_icon_v1_icon_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIconResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIconResponse) ProtoMessage() {}

func (x *DeleteIconResponse) ProtoReflect() protoreflect.Message {
	mi := &file_icon_v1_icon_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIconResponse.ProtoReflect.Descriptor instead.
func (*DeleteIconResponse) Descriptor() ([]byte, []int) {
	return file_icon_v1_icon_service_proto_rawDescGZIP(), []int{4}
}

var File_icon_v1_icon_service_proto protoreflect.FileDescriptor

var file_icon_v1_icon_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x63, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x63,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x64, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x69, 0x6e, 0x64, 0x6f, 0x66, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x69, 0x6e, 0x64, 0x6f, 0x66, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x75, 0x75, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x63, 0x6f,
	0x6e, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x75, 0x75, 0x69, 0x64, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x96, 0x01, 0x0a, 0x0b, 0x49, 0x63, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x12,
	0x17, 0x2e, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x63, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x63, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x63,
	0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x65, 0x74,
	0x7a, 0x79, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x63, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x63, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x63, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_icon_v1_icon_service_proto_rawDescOnce sync.Once
	file_icon_v1_icon_service_proto_rawDescData = file_icon_v1_icon_service_proto_rawDesc
)

func file_icon_v1_icon_service_proto_rawDescGZIP() []byte {
	file_icon_v1_icon_service_proto_rawDescOnce.Do(func() {
		file_icon_v1_icon_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_icon_v1_icon_service_proto_rawDescData)
	})
	return file_icon_v1_icon_service_proto_rawDescData
}

var file_icon_v1_icon_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_icon_v1_icon_service_proto_goTypes = []interface{}{
	(*Icon)(nil),               // 0: icon.v1.Icon
	(*GetIconRequest)(nil),     // 1: icon.v1.GetIconRequest
	(*GetIconResponse)(nil),    // 2: icon.v1.GetIconResponse
	(*DeleteIconRequest)(nil),  // 3: icon.v1.DeleteIconRequest
	(*DeleteIconResponse)(nil), // 4: icon.v1.DeleteIconResponse
}
var file_icon_v1_icon_service_proto_depIdxs = []int32{
	0, // 0: icon.v1.GetIconResponse.icon:type_name -> icon.v1.Icon
	1, // 1: icon.v1.IconService.GetIcon:input_type -> icon.v1.GetIconRequest
	3, // 2: icon.v1.IconService.DeleteIcon:input_type -> icon.v1.DeleteIconRequest
	2, // 3: icon.v1.IconService.GetIcon:output_type -> icon.v1.GetIconResponse
	4, // 4: icon.v1.IconService.DeleteIcon:output_type -> icon.v1.DeleteIconResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_icon_v1_icon_service_proto_init() }
func file_icon_v1_icon_service_proto_init() {
	if File_icon_v1_icon_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_icon_v1_icon_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Icon); i {
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
		file_icon_v1_icon_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIconRequest); i {
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
		file_icon_v1_icon_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIconResponse); i {
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
		file_icon_v1_icon_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIconRequest); i {
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
		file_icon_v1_icon_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIconResponse); i {
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
			RawDescriptor: file_icon_v1_icon_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_icon_v1_icon_service_proto_goTypes,
		DependencyIndexes: file_icon_v1_icon_service_proto_depIdxs,
		MessageInfos:      file_icon_v1_icon_service_proto_msgTypes,
	}.Build()
	File_icon_v1_icon_service_proto = out.File
	file_icon_v1_icon_service_proto_rawDesc = nil
	file_icon_v1_icon_service_proto_goTypes = nil
	file_icon_v1_icon_service_proto_depIdxs = nil
}
