// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sdk/blocksdk/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryLaneRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryLaneRequest) Reset()         { *m = QueryLaneRequest{} }
func (m *QueryLaneRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLaneRequest) ProtoMessage()    {}
func (*QueryLaneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{0}
}
func (m *QueryLaneRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLaneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLaneRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLaneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLaneRequest.Merge(m, src)
}
func (m *QueryLaneRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLaneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLaneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLaneRequest proto.InternalMessageInfo

func (m *QueryLaneRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryLaneResponse struct {
	Lane *Lane `protobuf:"bytes,1,opt,name=lane,proto3" json:"lane,omitempty"`
}

func (m *QueryLaneResponse) Reset()         { *m = QueryLaneResponse{} }
func (m *QueryLaneResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLaneResponse) ProtoMessage()    {}
func (*QueryLaneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{1}
}
func (m *QueryLaneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLaneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLaneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLaneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLaneResponse.Merge(m, src)
}
func (m *QueryLaneResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLaneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLaneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLaneResponse proto.InternalMessageInfo

func (m *QueryLaneResponse) GetLane() *Lane {
	if m != nil {
		return m.Lane
	}
	return nil
}

type QueryAllLanesRequest struct {
}

func (m *QueryAllLanesRequest) Reset()         { *m = QueryAllLanesRequest{} }
func (m *QueryAllLanesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllLanesRequest) ProtoMessage()    {}
func (*QueryAllLanesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{2}
}
func (m *QueryAllLanesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllLanesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllLanesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllLanesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllLanesRequest.Merge(m, src)
}
func (m *QueryAllLanesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllLanesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllLanesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllLanesRequest proto.InternalMessageInfo

type QueryAllLanesResponse struct {
	Lanes []*Lane `protobuf:"bytes,1,rep,name=lanes,proto3" json:"lanes,omitempty"`
}

func (m *QueryAllLanesResponse) Reset()         { *m = QueryAllLanesResponse{} }
func (m *QueryAllLanesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllLanesResponse) ProtoMessage()    {}
func (*QueryAllLanesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c631e2e97c81d34, []int{3}
}
func (m *QueryAllLanesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllLanesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllLanesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllLanesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllLanesResponse.Merge(m, src)
}
func (m *QueryAllLanesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllLanesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllLanesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllLanesResponse proto.InternalMessageInfo

func (m *QueryAllLanesResponse) GetLanes() []*Lane {
	if m != nil {
		return m.Lanes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryLaneRequest)(nil), "sdk.blocksdk.v1.QueryLaneRequest")
	proto.RegisterType((*QueryLaneResponse)(nil), "sdk.blocksdk.v1.QueryLaneResponse")
	proto.RegisterType((*QueryAllLanesRequest)(nil), "sdk.blocksdk.v1.QueryAllLanesRequest")
	proto.RegisterType((*QueryAllLanesResponse)(nil), "sdk.blocksdk.v1.QueryAllLanesResponse")
}

func init() { proto.RegisterFile("sdk/blocksdk/v1/query.proto", fileDescriptor_0c631e2e97c81d34) }

var fileDescriptor_0c631e2e97c81d34 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0xc9, 0xd6,
	0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x06, 0x31, 0xca, 0x0c, 0xf5, 0x0b, 0x4b, 0x53, 0x8b, 0x2a, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x8b, 0x53, 0xb2, 0xf5, 0x60, 0x92, 0x7a, 0x65, 0x86,
	0x52, 0x72, 0xe8, 0xaa, 0xe1, 0x92, 0x60, 0x0d, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9,
	0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5,
	0x50, 0x59, 0xe9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0x88, 0x15, 0x68, 0x76, 0x29, 0x29, 0x71,
	0x09, 0x04, 0x82, 0xb8, 0x3e, 0x89, 0x79, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a,
	0x76, 0x5c, 0x82, 0x48, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x34, 0xb9, 0x58, 0x72,
	0x12, 0xf3, 0x52, 0xc1, 0xca, 0xb8, 0x8d, 0x44, 0xf5, 0xd0, 0xdc, 0xac, 0x07, 0x56, 0x0c, 0x56,
	0xa2, 0x24, 0xc6, 0x25, 0x02, 0xd6, 0xef, 0x98, 0x93, 0x03, 0x12, 0x2d, 0x86, 0xda, 0xa3, 0xe4,
	0xc2, 0x25, 0x8a, 0x26, 0x0e, 0x35, 0x5b, 0x9b, 0x8b, 0x15, 0xa4, 0xb1, 0x58, 0x82, 0x51, 0x81,
	0x19, 0xb7, 0xe1, 0x10, 0x35, 0x46, 0x9d, 0x4c, 0x5c, 0xac, 0x60, 0x63, 0x84, 0x6a, 0xb8, 0x58,
	0x40, 0x12, 0x42, 0x8a, 0x18, 0xea, 0xd1, 0xbd, 0x28, 0xa5, 0x84, 0x4f, 0x09, 0xc4, 0x15, 0x4a,
	0xba, 0x1d, 0xcf, 0x37, 0x68, 0x31, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x49, 0x48, 0x01, 0x12,
	0xe6, 0xba, 0xe8, 0x31, 0x01, 0x72, 0x83, 0x7e, 0x75, 0x66, 0x4a, 0xad, 0x50, 0x23, 0x23, 0x17,
	0x2b, 0xd8, 0x1b, 0x42, 0xaa, 0xd8, 0x0d, 0x47, 0xf3, 0xbe, 0x94, 0x1a, 0x21, 0x65, 0x50, 0x77,
	0x68, 0x22, 0xdc, 0x21, 0x27, 0x24, 0x83, 0xc7, 0x1d, 0xc5, 0x4e, 0x1e, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x5f, 0x9c, 0x9d, 0x59, 0xa0, 0x9b, 0x9b, 0x5a, 0x86, 0x64, 0x54, 0x05, 0xc2, 0xb0,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xf2, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x0a, 0x6c, 0x8f, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Lane(ctx context.Context, in *QueryLaneRequest, opts ...grpc.CallOption) (*QueryLaneResponse, error)
	Lanes(ctx context.Context, in *QueryAllLanesRequest, opts ...grpc.CallOption) (*QueryAllLanesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Lane(ctx context.Context, in *QueryLaneRequest, opts ...grpc.CallOption) (*QueryLaneResponse, error) {
	out := new(QueryLaneResponse)
	err := c.cc.Invoke(ctx, "/sdk.blocksdk.v1.Query/Lane", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Lanes(ctx context.Context, in *QueryAllLanesRequest, opts ...grpc.CallOption) (*QueryAllLanesResponse, error) {
	out := new(QueryAllLanesResponse)
	err := c.cc.Invoke(ctx, "/sdk.blocksdk.v1.Query/Lanes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Lane(context.Context, *QueryLaneRequest) (*QueryLaneResponse, error)
	Lanes(context.Context, *QueryAllLanesRequest) (*QueryAllLanesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Lane(ctx context.Context, req *QueryLaneRequest) (*QueryLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lane not implemented")
}
func (*UnimplementedQueryServer) Lanes(ctx context.Context, req *QueryAllLanesRequest) (*QueryAllLanesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lanes not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Lane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Lane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.blocksdk.v1.Query/Lane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Lane(ctx, req.(*QueryLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Lanes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllLanesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Lanes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sdk.blocksdk.v1.Query/Lanes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Lanes(ctx, req.(*QueryAllLanesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sdk.blocksdk.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lane",
			Handler:    _Query_Lane_Handler,
		},
		{
			MethodName: "Lanes",
			Handler:    _Query_Lanes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdk/blocksdk/v1/query.proto",
}

func (m *QueryLaneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLaneRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLaneRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLaneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLaneResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLaneResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Lane != nil {
		{
			size, err := m.Lane.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllLanesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllLanesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllLanesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllLanesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllLanesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllLanesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for iNdEx := len(m.Lanes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Lanes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryLaneRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLaneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lane != nil {
		l = m.Lane.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllLanesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllLanesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Lanes) > 0 {
		for _, e := range m.Lanes {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryLaneRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLaneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLaneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryLaneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryLaneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLaneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lane", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lane == nil {
				m.Lane = &Lane{}
			}
			if err := m.Lane.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllLanesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllLanesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllLanesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllLanesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllLanesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllLanesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lanes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lanes = append(m.Lanes, &Lane{})
			if err := m.Lanes[len(m.Lanes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
