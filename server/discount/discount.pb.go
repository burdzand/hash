// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discount.proto

package discount

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Discount struct {
	Pct                  float32  `protobuf:"fixed32,1,opt,name=pct,proto3" json:"pct,omitempty"`
	ValueInCents         int32    `protobuf:"varint,2,opt,name=value_in_cents,json=valueInCents,proto3" json:"value_in_cents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Discount) Reset()         { *m = Discount{} }
func (m *Discount) String() string { return proto.CompactTextString(m) }
func (*Discount) ProtoMessage()    {}
func (*Discount) Descriptor() ([]byte, []int) {
	return fileDescriptor_78985be6a51b316f, []int{0}
}

func (m *Discount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Discount.Unmarshal(m, b)
}
func (m *Discount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Discount.Marshal(b, m, deterministic)
}
func (m *Discount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discount.Merge(m, src)
}
func (m *Discount) XXX_Size() int {
	return xxx_messageInfo_Discount.Size(m)
}
func (m *Discount) XXX_DiscardUnknown() {
	xxx_messageInfo_Discount.DiscardUnknown(m)
}

var xxx_messageInfo_Discount proto.InternalMessageInfo

func (m *Discount) GetPct() float32 {
	if m != nil {
		return m.Pct
	}
	return 0
}

func (m *Discount) GetValueInCents() int32 {
	if m != nil {
		return m.ValueInCents
	}
	return 0
}

type DiscountResquest struct {
	UserID               int32    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ProductID            int32    `protobuf:"varint,2,opt,name=productID,proto3" json:"productID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscountResquest) Reset()         { *m = DiscountResquest{} }
func (m *DiscountResquest) String() string { return proto.CompactTextString(m) }
func (*DiscountResquest) ProtoMessage()    {}
func (*DiscountResquest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78985be6a51b316f, []int{1}
}

func (m *DiscountResquest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscountResquest.Unmarshal(m, b)
}
func (m *DiscountResquest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscountResquest.Marshal(b, m, deterministic)
}
func (m *DiscountResquest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscountResquest.Merge(m, src)
}
func (m *DiscountResquest) XXX_Size() int {
	return xxx_messageInfo_DiscountResquest.Size(m)
}
func (m *DiscountResquest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscountResquest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscountResquest proto.InternalMessageInfo

func (m *DiscountResquest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *DiscountResquest) GetProductID() int32 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func init() {
	proto.RegisterType((*Discount)(nil), "discount.Discount")
	proto.RegisterType((*DiscountResquest)(nil), "discount.DiscountResquest")
}

func init() { proto.RegisterFile("discount.proto", fileDescriptor_78985be6a51b316f) }

var fileDescriptor_78985be6a51b316f = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2c, 0x4e,
	0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x9c,
	0xb8, 0x38, 0x5c, 0xa0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0x82, 0xe4, 0x12, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xa6, 0x20, 0x10, 0x53, 0x48, 0x85, 0x8b, 0xaf, 0x2c, 0x31, 0xa7, 0x34, 0x35, 0x3e, 0x33,
	0x2f, 0x3e, 0x39, 0x35, 0xaf, 0xa4, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x88, 0x07, 0x2c,
	0xea, 0x99, 0xe7, 0x0c, 0x12, 0x53, 0xf2, 0xe0, 0x12, 0x80, 0x99, 0x11, 0x94, 0x5a, 0x5c, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc6, 0xc5, 0x56, 0x5a, 0x9c, 0x5a, 0xe4, 0xe9, 0x02, 0x36, 0x8e,
	0x35, 0x08, 0xca, 0x13, 0x92, 0xe1, 0xe2, 0x2c, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xf1, 0x74,
	0x81, 0x1a, 0x86, 0x10, 0x30, 0xf2, 0xe2, 0xe2, 0x87, 0x99, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x99,
	0x9c, 0x2a, 0x64, 0xce, 0xc5, 0xec, 0x9e, 0x5a, 0x22, 0x24, 0xa5, 0x07, 0xf7, 0x02, 0xba, 0x5d,
	0x52, 0x42, 0x98, 0x72, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xaf, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x9b, 0x2e, 0x98, 0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscountServiceClient is the client API for DiscountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscountServiceClient interface {
	Get(ctx context.Context, in *DiscountResquest, opts ...grpc.CallOption) (*Discount, error)
}

type discountServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscountServiceClient(cc *grpc.ClientConn) DiscountServiceClient {
	return &discountServiceClient{cc}
}

func (c *discountServiceClient) Get(ctx context.Context, in *DiscountResquest, opts ...grpc.CallOption) (*Discount, error) {
	out := new(Discount)
	err := c.cc.Invoke(ctx, "/discount.DiscountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscountServiceServer is the server API for DiscountService service.
type DiscountServiceServer interface {
	Get(context.Context, *DiscountResquest) (*Discount, error)
}

func RegisterDiscountServiceServer(s *grpc.Server, srv DiscountServiceServer) {
	s.RegisterService(&_DiscountService_serviceDesc, srv)
}

func _DiscountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discount.DiscountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscountServiceServer).Get(ctx, req.(*DiscountResquest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discount.DiscountService",
	HandlerType: (*DiscountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DiscountService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discount.proto",
}
