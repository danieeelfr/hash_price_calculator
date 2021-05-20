// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: discountpb/discount.proto

package discountpb

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

type DiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *DiscountRequest) Reset() {
	*x = DiscountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discountpb_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountRequest) ProtoMessage() {}

func (x *DiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discountpb_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountRequest.ProtoReflect.Descriptor instead.
func (*DiscountRequest) Descriptor() ([]byte, []int) {
	return file_discountpb_discount_proto_rawDescGZIP(), []int{0}
}

func (x *DiscountRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DiscountRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type DiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discount float64 `protobuf:"fixed64,1,opt,name=discount,proto3" json:"discount,omitempty"`
}

func (x *DiscountResponse) Reset() {
	*x = DiscountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discountpb_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountResponse) ProtoMessage() {}

func (x *DiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discountpb_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountResponse.ProtoReflect.Descriptor instead.
func (*DiscountResponse) Descriptor() ([]byte, []int) {
	return file_discountpb_discount_proto_rawDescGZIP(), []int{1}
}

func (x *DiscountResponse) GetDiscount() float64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

var File_discountpb_discount_proto protoreflect.FileDescriptor

var file_discountpb_discount_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x59, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_discountpb_discount_proto_rawDescOnce sync.Once
	file_discountpb_discount_proto_rawDescData = file_discountpb_discount_proto_rawDesc
)

func file_discountpb_discount_proto_rawDescGZIP() []byte {
	file_discountpb_discount_proto_rawDescOnce.Do(func() {
		file_discountpb_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discountpb_discount_proto_rawDescData)
	})
	return file_discountpb_discount_proto_rawDescData
}

var file_discountpb_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discountpb_discount_proto_goTypes = []interface{}{
	(*DiscountRequest)(nil),  // 0: discount.DiscountRequest
	(*DiscountResponse)(nil), // 1: discount.DiscountResponse
}
var file_discountpb_discount_proto_depIdxs = []int32{
	0, // 0: discount.DiscountService.GetDiscount:input_type -> discount.DiscountRequest
	1, // 1: discount.DiscountService.GetDiscount:output_type -> discount.DiscountResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_discountpb_discount_proto_init() }
func file_discountpb_discount_proto_init() {
	if File_discountpb_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discountpb_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountRequest); i {
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
		file_discountpb_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountResponse); i {
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
			RawDescriptor: file_discountpb_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discountpb_discount_proto_goTypes,
		DependencyIndexes: file_discountpb_discount_proto_depIdxs,
		MessageInfos:      file_discountpb_discount_proto_msgTypes,
	}.Build()
	File_discountpb_discount_proto = out.File
	file_discountpb_discount_proto_rawDesc = nil
	file_discountpb_discount_proto_goTypes = nil
	file_discountpb_discount_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiscountServiceClient is the client API for DiscountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscountServiceClient interface {
	GetDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error)
}

type discountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscountServiceClient(cc grpc.ClientConnInterface) DiscountServiceClient {
	return &discountServiceClient{cc}
}

func (c *discountServiceClient) GetDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error) {
	out := new(DiscountResponse)
	err := c.cc.Invoke(ctx, "/discount.DiscountService/GetDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServiceServer is the server API for DiscountService service.
type DiscountServiceServer interface {
	GetDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error)
}

// UnimplementedDiscountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiscountServiceServer struct {
}

func (*UnimplementedDiscountServiceServer) GetDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}

func RegisterDiscountServiceServer(s *grpc.Server, srv DiscountServiceServer) {
	s.RegisterService(&_DiscountService_serviceDesc, srv)
}

func _DiscountService_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServiceServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.DiscountService/GetDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServiceServer).GetDiscount(ctx, req.(*DiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discount.DiscountService",
	HandlerType: (*DiscountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiscount",
			Handler:    _DiscountService_GetDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discountpb/discount.proto",
}
