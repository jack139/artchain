// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trans/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// this line is used by starport scaffolding # 3
type QueryGetTransactionRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetTransactionRequest) Reset()         { *m = QueryGetTransactionRequest{} }
func (m *QueryGetTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTransactionRequest) ProtoMessage()    {}
func (*QueryGetTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60381eeb3321a7dc, []int{0}
}
func (m *QueryGetTransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTransactionRequest.Merge(m, src)
}
func (m *QueryGetTransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTransactionRequest proto.InternalMessageInfo

func (m *QueryGetTransactionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetTransactionResponse struct {
	Transaction *Transaction `protobuf:"bytes,1,opt,name=Transaction,json=transaction,proto3" json:"Transaction,omitempty"`
}

func (m *QueryGetTransactionResponse) Reset()         { *m = QueryGetTransactionResponse{} }
func (m *QueryGetTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTransactionResponse) ProtoMessage()    {}
func (*QueryGetTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60381eeb3321a7dc, []int{1}
}
func (m *QueryGetTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTransactionResponse.Merge(m, src)
}
func (m *QueryGetTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTransactionResponse proto.InternalMessageInfo

func (m *QueryGetTransactionResponse) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type QueryAllTransactionRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTransactionRequest) Reset()         { *m = QueryAllTransactionRequest{} }
func (m *QueryAllTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTransactionRequest) ProtoMessage()    {}
func (*QueryAllTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60381eeb3321a7dc, []int{2}
}
func (m *QueryAllTransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTransactionRequest.Merge(m, src)
}
func (m *QueryAllTransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTransactionRequest proto.InternalMessageInfo

func (m *QueryAllTransactionRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTransactionResponse struct {
	Transaction []*Transaction      `protobuf:"bytes,1,rep,name=Transaction,json=transaction,proto3" json:"Transaction,omitempty"`
	Pagination  *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTransactionResponse) Reset()         { *m = QueryAllTransactionResponse{} }
func (m *QueryAllTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTransactionResponse) ProtoMessage()    {}
func (*QueryAllTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60381eeb3321a7dc, []int{3}
}
func (m *QueryAllTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTransactionResponse.Merge(m, src)
}
func (m *QueryAllTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTransactionResponse proto.InternalMessageInfo

func (m *QueryAllTransactionResponse) GetTransaction() []*Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *QueryAllTransactionResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetTransactionRequest)(nil), "jack139.artchain.trans.QueryGetTransactionRequest")
	proto.RegisterType((*QueryGetTransactionResponse)(nil), "jack139.artchain.trans.QueryGetTransactionResponse")
	proto.RegisterType((*QueryAllTransactionRequest)(nil), "jack139.artchain.trans.QueryAllTransactionRequest")
	proto.RegisterType((*QueryAllTransactionResponse)(nil), "jack139.artchain.trans.QueryAllTransactionResponse")
}

func init() { proto.RegisterFile("trans/query.proto", fileDescriptor_60381eeb3321a7dc) }

var fileDescriptor_60381eeb3321a7dc = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x29, 0x4a, 0xcc,
	0x2b, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb,
	0x4a, 0x4c, 0xce, 0x36, 0x34, 0xb6, 0xd4, 0x4b, 0x2c, 0x2a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3,
	0x03, 0xab, 0x91, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xe8, 0x92, 0xd2, 0x4a, 0xce,
	0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0x85, 0x18, 0xa7, 0x5f, 0x66, 0x98, 0x94,
	0x5a, 0x92, 0x68, 0xa8, 0x5f, 0x90, 0x98, 0x9e, 0x99, 0x07, 0x56, 0x0c, 0x55, 0x2b, 0x0e, 0xb1,
	0x14, 0x4c, 0x26, 0x26, 0x23, 0x24, 0x94, 0x74, 0xb8, 0xa4, 0x02, 0x41, 0x5a, 0xdd, 0x53, 0x4b,
	0x42, 0x10, 0x92, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0x29, 0x5c, 0xd2, 0x58, 0x55,
	0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xb9, 0x72, 0x71, 0x23, 0x09, 0x83, 0xf5, 0x71, 0x1b,
	0x29, 0xeb, 0x61, 0xf7, 0x9d, 0x1e, 0xb2, 0x09, 0xdc, 0x48, 0x2e, 0x53, 0x4a, 0x81, 0xba, 0xc9,
	0x31, 0x27, 0x07, 0x8b, 0x9b, 0xdc, 0xb8, 0xb8, 0x10, 0xde, 0x83, 0xda, 0xa1, 0xa6, 0x07, 0x09,
	0x0b, 0x3d, 0x50, 0x58, 0xe8, 0x41, 0x82, 0x16, 0x1a, 0x16, 0x7a, 0x01, 0x89, 0xe9, 0xa9, 0x50,
	0xbd, 0x41, 0x48, 0x3a, 0x95, 0xd6, 0x32, 0x42, 0x3d, 0x83, 0x6e, 0x0d, 0x2e, 0xcf, 0x30, 0x93,
	0xe3, 0x19, 0x21, 0x77, 0x14, 0xe7, 0x32, 0x81, 0x9d, 0xab, 0x4e, 0xd0, 0xb9, 0x10, 0x37, 0x20,
	0xbb, 0xd7, 0xe8, 0x36, 0x13, 0x17, 0x2b, 0xd8, 0xbd, 0x42, 0x2b, 0x18, 0x51, 0x9c, 0x26, 0x64,
	0x84, 0xcb, 0x51, 0xb8, 0x63, 0x56, 0xca, 0x98, 0x24, 0x3d, 0x10, 0xe7, 0x28, 0x19, 0x34, 0x5d,
	0x7e, 0x32, 0x99, 0x49, 0x4b, 0x48, 0x43, 0x1f, 0xaa, 0x59, 0x1f, 0xa6, 0x59, 0x1f, 0x23, 0x7d,
	0xe9, 0x57, 0x67, 0xa6, 0xd4, 0x0a, 0x2d, 0x63, 0xe4, 0xe2, 0x43, 0x32, 0xc9, 0x31, 0x27, 0x87,
	0x80, 0x6b, 0xb1, 0xc6, 0x39, 0x01, 0xd7, 0x62, 0x8f, 0x40, 0x25, 0x6d, 0xb0, 0x6b, 0x55, 0x85,
	0x94, 0x89, 0x70, 0xad, 0x93, 0xf3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x62, 0x1a, 0x54, 0x01, 0x33,
	0xaa, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x9c, 0xa7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x1f, 0x10, 0xe7, 0xe3, 0x03, 0x00, 0x00,
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
	// this line is used by starport scaffolding # 2
	Transaction(ctx context.Context, in *QueryGetTransactionRequest, opts ...grpc.CallOption) (*QueryGetTransactionResponse, error)
	TransactionAll(ctx context.Context, in *QueryAllTransactionRequest, opts ...grpc.CallOption) (*QueryAllTransactionResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Transaction(ctx context.Context, in *QueryGetTransactionRequest, opts ...grpc.CallOption) (*QueryGetTransactionResponse, error) {
	out := new(QueryGetTransactionResponse)
	err := c.cc.Invoke(ctx, "/jack139.artchain.trans.Query/Transaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TransactionAll(ctx context.Context, in *QueryAllTransactionRequest, opts ...grpc.CallOption) (*QueryAllTransactionResponse, error) {
	out := new(QueryAllTransactionResponse)
	err := c.cc.Invoke(ctx, "/jack139.artchain.trans.Query/TransactionAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// this line is used by starport scaffolding # 2
	Transaction(context.Context, *QueryGetTransactionRequest) (*QueryGetTransactionResponse, error)
	TransactionAll(context.Context, *QueryAllTransactionRequest) (*QueryAllTransactionResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Transaction(ctx context.Context, req *QueryGetTransactionRequest) (*QueryGetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transaction not implemented")
}
func (*UnimplementedQueryServer) TransactionAll(ctx context.Context, req *QueryAllTransactionRequest) (*QueryAllTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Transaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Transaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jack139.artchain.trans.Query/Transaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Transaction(ctx, req.(*QueryGetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TransactionAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TransactionAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jack139.artchain.trans.Query/TransactionAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TransactionAll(ctx, req.(*QueryAllTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jack139.artchain.trans.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transaction",
			Handler:    _Query_Transaction_Handler,
		},
		{
			MethodName: "TransactionAll",
			Handler:    _Query_TransactionAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trans/query.proto",
}

func (m *QueryGetTransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Transaction != nil {
		{
			size, err := m.Transaction.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllTransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Transaction) > 0 {
		for iNdEx := len(m.Transaction) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transaction[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryGetTransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transaction != nil {
		l = m.Transaction.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Transaction) > 0 {
		for _, e := range m.Transaction {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetTransactionRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryGetTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
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
			if m.Transaction == nil {
				m.Transaction = &Transaction{}
			}
			if err := m.Transaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllTransactionRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
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
			m.Transaction = append(m.Transaction, &Transaction{})
			if err := m.Transaction[len(m.Transaction)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
