// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/distance/proto/distance.proto

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

// Mesasage that refers to a map position tuple (latitude / longitude)
type MapPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *MapPosition) Reset() {
	*x = MapPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distance_proto_distance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapPosition) ProtoMessage() {}

func (x *MapPosition) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distance_proto_distance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapPosition.ProtoReflect.Descriptor instead.
func (*MapPosition) Descriptor() ([]byte, []int) {
	return file_pkg_distance_proto_distance_proto_rawDescGZIP(), []int{0}
}

func (x *MapPosition) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *MapPosition) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// The request message containing two tuple of lat and lon
type DistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      *MapPosition `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination *MapPosition `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *DistanceRequest) Reset() {
	*x = DistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distance_proto_distance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceRequest) ProtoMessage() {}

func (x *DistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distance_proto_distance_proto_msgTypes[1]
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
	return file_pkg_distance_proto_distance_proto_rawDescGZIP(), []int{1}
}

func (x *DistanceRequest) GetOrigin() *MapPosition {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *DistanceRequest) GetDestination() *MapPosition {
	if x != nil {
		return x.Destination
	}
	return nil
}

// The response message containing the distance between the two positions in
// meters format
type DistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin      string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Distance    int32  `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *DistanceResponse) Reset() {
	*x = DistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distance_proto_distance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistanceResponse) ProtoMessage() {}

func (x *DistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distance_proto_distance_proto_msgTypes[2]
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
	return file_pkg_distance_proto_distance_proto_rawDescGZIP(), []int{2}
}

func (x *DistanceResponse) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *DistanceResponse) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *DistanceResponse) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

// Address request
type AddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_distance_proto_distance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_distance_proto_distance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_pkg_distance_proto_distance_proto_rawDescGZIP(), []int{3}
}

func (x *AddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_pkg_distance_proto_distance_proto protoreflect.FileDescriptor

var file_pkg_distance_proto_distance_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x47, 0x0a,
	0x0b, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x79, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x68, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x96, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x63, 0x6d, 0x65, 0x2d, 0x73, 0x6b, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_distance_proto_distance_proto_rawDescOnce sync.Once
	file_pkg_distance_proto_distance_proto_rawDescData = file_pkg_distance_proto_distance_proto_rawDesc
)

func file_pkg_distance_proto_distance_proto_rawDescGZIP() []byte {
	file_pkg_distance_proto_distance_proto_rawDescOnce.Do(func() {
		file_pkg_distance_proto_distance_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_distance_proto_distance_proto_rawDescData)
	})
	return file_pkg_distance_proto_distance_proto_rawDescData
}

var file_pkg_distance_proto_distance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_distance_proto_distance_proto_goTypes = []interface{}{
	(*MapPosition)(nil),      // 0: distance.MapPosition
	(*DistanceRequest)(nil),  // 1: distance.DistanceRequest
	(*DistanceResponse)(nil), // 2: distance.DistanceResponse
	(*AddressRequest)(nil),   // 3: distance.AddressRequest
}
var file_pkg_distance_proto_distance_proto_depIdxs = []int32{
	0, // 0: distance.DistanceRequest.origin:type_name -> distance.MapPosition
	0, // 1: distance.DistanceRequest.destination:type_name -> distance.MapPosition
	1, // 2: distance.Distance.FindDistance:input_type -> distance.DistanceRequest
	3, // 3: distance.Distance.FindGeometry:input_type -> distance.AddressRequest
	2, // 4: distance.Distance.FindDistance:output_type -> distance.DistanceResponse
	0, // 5: distance.Distance.FindGeometry:output_type -> distance.MapPosition
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_distance_proto_distance_proto_init() }
func file_pkg_distance_proto_distance_proto_init() {
	if File_pkg_distance_proto_distance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_distance_proto_distance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapPosition); i {
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
		file_pkg_distance_proto_distance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_distance_proto_distance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_distance_proto_distance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressRequest); i {
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
			RawDescriptor: file_pkg_distance_proto_distance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_distance_proto_distance_proto_goTypes,
		DependencyIndexes: file_pkg_distance_proto_distance_proto_depIdxs,
		MessageInfos:      file_pkg_distance_proto_distance_proto_msgTypes,
	}.Build()
	File_pkg_distance_proto_distance_proto = out.File
	file_pkg_distance_proto_distance_proto_rawDesc = nil
	file_pkg_distance_proto_distance_proto_goTypes = nil
	file_pkg_distance_proto_distance_proto_depIdxs = nil
}
