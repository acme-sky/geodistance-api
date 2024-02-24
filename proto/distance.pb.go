// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/distance.proto

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

// The request message containing two tuple of lat and lon
type DistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat1 float32 `protobuf:"fixed32,1,opt,name=lat1,proto3" json:"lat1,omitempty"`
	Lon1 float32 `protobuf:"fixed32,2,opt,name=lon1,proto3" json:"lon1,omitempty"`
	Lat2 float32 `protobuf:"fixed32,3,opt,name=lat2,proto3" json:"lat2,omitempty"`
	Lon2 float32 `protobuf:"fixed32,4,opt,name=lon2,proto3" json:"lon2,omitempty"`
}

func (x *DistanceRequest) Reset() {
	*x = DistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_distance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceRequest) ProtoMessage() {}

func (x *DistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_distance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistanceRequest.ProtoReflect.Descriptor instead.
func (*DistanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_distance_proto_rawDescGZIP(), []int{0}
}

func (x *DistanceRequest) GetLat1() float32 {
	if x != nil {
		return x.Lat1
	}
	return 0
}

func (x *DistanceRequest) GetLon1() float32 {
	if x != nil {
		return x.Lon1
	}
	return 0
}

func (x *DistanceRequest) GetLat2() float32 {
	if x != nil {
		return x.Lat2
	}
	return 0
}

func (x *DistanceRequest) GetLon2() float32 {
	if x != nil {
		return x.Lon2
	}
	return 0
}

// The response message containing the distance between the two positions in
// meters format
type DistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance float32 `protobuf:"fixed32,1,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *DistanceResponse) Reset() {
	*x = DistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_distance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceResponse) ProtoMessage() {}

func (x *DistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_distance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistanceResponse.ProtoReflect.Descriptor instead.
func (*DistanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_distance_proto_rawDescGZIP(), []int{1}
}

func (x *DistanceResponse) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_proto_distance_proto protoreflect.FileDescriptor

var file_proto_distance_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x61, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x6c, 0x61, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x74, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x61, 0x74, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c,
	0x6f, 0x6e, 0x32, 0x22, 0x2e, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x32, 0x52, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x19,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2d, 0x73, 0x6b, 0x79, 0x2f, 0x67,
	0x65, 0x6f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_distance_proto_rawDescOnce sync.Once
	file_proto_distance_proto_rawDescData = file_proto_distance_proto_rawDesc
)

func file_proto_distance_proto_rawDescGZIP() []byte {
	file_proto_distance_proto_rawDescOnce.Do(func() {
		file_proto_distance_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_distance_proto_rawDescData)
	})
	return file_proto_distance_proto_rawDescData
}

var file_proto_distance_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_distance_proto_goTypes = []interface{}{
	(*DistanceRequest)(nil),  // 0: distance.DistanceRequest
	(*DistanceResponse)(nil), // 1: distance.DistanceResponse
}
var file_proto_distance_proto_depIdxs = []int32{
	0, // 0: distance.Distance.GetDistance:input_type -> distance.DistanceRequest
	1, // 1: distance.Distance.GetDistance:output_type -> distance.DistanceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_distance_proto_init() }
func file_proto_distance_proto_init() {
	if File_proto_distance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_distance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistanceRequest); i {
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
		file_proto_distance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistanceResponse); i {
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
			RawDescriptor: file_proto_distance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_distance_proto_goTypes,
		DependencyIndexes: file_proto_distance_proto_depIdxs,
		MessageInfos:      file_proto_distance_proto_msgTypes,
	}.Build()
	File_proto_distance_proto = out.File
	file_proto_distance_proto_rawDesc = nil
	file_proto_distance_proto_goTypes = nil
	file_proto_distance_proto_depIdxs = nil
}
