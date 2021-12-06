// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: surge.proto

package surgegrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DemandAreaKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DemandAreaKey) Reset() {
	*x = DemandAreaKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemandAreaKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemandAreaKey) ProtoMessage() {}

func (x *DemandAreaKey) ProtoReflect() protoreflect.Message {
	mi := &file_surge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemandAreaKey.ProtoReflect.Descriptor instead.
func (*DemandAreaKey) Descriptor() ([]byte, []int) {
	return file_surge_proto_rawDescGZIP(), []int{0}
}

func (x *DemandAreaKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SurgeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate float32 `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *SurgeRate) Reset() {
	*x = SurgeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_surge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurgeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurgeRate) ProtoMessage() {}

func (x *SurgeRate) ProtoReflect() protoreflect.Message {
	mi := &file_surge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurgeRate.ProtoReflect.Descriptor instead.
func (*SurgeRate) Descriptor() ([]byte, []int) {
	return file_surge_proto_rawDescGZIP(), []int{1}
}

func (x *SurgeRate) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_surge_proto protoreflect.FileDescriptor

var file_surge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x75, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73,
	0x75, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x75, 0x72, 0x67, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22,
	0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x75, 0x72, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x32, 0x5e, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x75, 0x72, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x75, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x75, 0x72, 0x67,
	0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x65,
	0x61, 0x4b, 0x65, 0x79, 0x1a, 0x1b, 0x2e, 0x73, 0x75, 0x72, 0x67, 0x65, 0x2e, 0x73, 0x75, 0x72,
	0x67, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x72, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x73, 0x65, 0x66, 0x73, 0x73, 0x2f, 0x63, 0x61, 0x62, 0x2d, 0x70, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x75, 0x72, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_surge_proto_rawDescOnce sync.Once
	file_surge_proto_rawDescData = file_surge_proto_rawDesc
)

func file_surge_proto_rawDescGZIP() []byte {
	file_surge_proto_rawDescOnce.Do(func() {
		file_surge_proto_rawDescData = protoimpl.X.CompressGZIP(file_surge_proto_rawDescData)
	})
	return file_surge_proto_rawDescData
}

var file_surge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_surge_proto_goTypes = []interface{}{
	(*DemandAreaKey)(nil), // 0: surge.surge.grpc.DemandAreaKey
	(*SurgeRate)(nil),     // 1: surge.surge.grpc.SurgeRate
}
var file_surge_proto_depIdxs = []int32{
	0, // 0: surge.surge.grpc.SurgeService.GetSurgeRate:input_type -> surge.surge.grpc.DemandAreaKey
	1, // 1: surge.surge.grpc.SurgeService.GetSurgeRate:output_type -> surge.surge.grpc.SurgeRate
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_surge_proto_init() }
func file_surge_proto_init() {
	if File_surge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_surge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemandAreaKey); i {
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
		file_surge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurgeRate); i {
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
			RawDescriptor: file_surge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_surge_proto_goTypes,
		DependencyIndexes: file_surge_proto_depIdxs,
		MessageInfos:      file_surge_proto_msgTypes,
	}.Build()
	File_surge_proto = out.File
	file_surge_proto_rawDesc = nil
	file_surge_proto_goTypes = nil
	file_surge_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SurgeServiceClient is the client API for SurgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SurgeServiceClient interface {
	GetSurgeRate(ctx context.Context, in *DemandAreaKey, opts ...grpc.CallOption) (*SurgeRate, error)
}

type surgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSurgeServiceClient(cc grpc.ClientConnInterface) SurgeServiceClient {
	return &surgeServiceClient{cc}
}

func (c *surgeServiceClient) GetSurgeRate(ctx context.Context, in *DemandAreaKey, opts ...grpc.CallOption) (*SurgeRate, error) {
	out := new(SurgeRate)
	err := c.cc.Invoke(ctx, "/surge.surge.grpc.SurgeService/GetSurgeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurgeServiceServer is the server API for SurgeService service.
type SurgeServiceServer interface {
	GetSurgeRate(context.Context, *DemandAreaKey) (*SurgeRate, error)
}

// UnimplementedSurgeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSurgeServiceServer struct {
}

func (*UnimplementedSurgeServiceServer) GetSurgeRate(context.Context, *DemandAreaKey) (*SurgeRate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSurgeRate not implemented")
}

func RegisterSurgeServiceServer(s *grpc.Server, srv SurgeServiceServer) {
	s.RegisterService(&_SurgeService_serviceDesc, srv)
}

func _SurgeService_GetSurgeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemandAreaKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurgeServiceServer).GetSurgeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/surge.surge.grpc.SurgeService/GetSurgeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurgeServiceServer).GetSurgeRate(ctx, req.(*DemandAreaKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _SurgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "surge.surge.grpc.SurgeService",
	HandlerType: (*SurgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSurgeRate",
			Handler:    _SurgeService_GetSurgeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "surge.proto",
}
