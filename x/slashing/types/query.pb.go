// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/v43beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/v43/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryParamsRequest is the request type for the Query/Params RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
// method
type QuerySigningInfoRequest struct {
	// cons_address is the address to query signing info of
	ConsAddress string `protobuf:"bytes,1,opt,name=cons_address,json=consAddress,proto3" json:"cons_address,omitempty"`
}

func (m *QuerySigningInfoRequest) Reset()         { *m = QuerySigningInfoRequest{} }
func (m *QuerySigningInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoRequest) ProtoMessage()    {}
func (*QuerySigningInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{2}
}
func (m *QuerySigningInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoRequest.Merge(m, src)
}
func (m *QuerySigningInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoRequest proto.InternalMessageInfo

func (m *QuerySigningInfoRequest) GetConsAddress() string {
	if m != nil {
		return m.ConsAddress
	}
	return ""
}

// QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
type QuerySigningInfoResponse struct {
	// val_signing_info is the signing info of requested val cons address
	ValSigningInfo ValidatorSigningInfo `protobuf:"bytes,1,opt,name=val_signing_info,json=valSigningInfo,proto3" json:"val_signing_info"`
}

func (m *QuerySigningInfoResponse) Reset()         { *m = QuerySigningInfoResponse{} }
func (m *QuerySigningInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoResponse) ProtoMessage()    {}
func (*QuerySigningInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{3}
}
func (m *QuerySigningInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoResponse.Merge(m, src)
}
func (m *QuerySigningInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoResponse proto.InternalMessageInfo

func (m *QuerySigningInfoResponse) GetValSigningInfo() ValidatorSigningInfo {
	if m != nil {
		return m.ValSigningInfo
	}
	return ValidatorSigningInfo{}
}

// QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
// method
type QuerySigningInfosRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySigningInfosRequest) Reset()         { *m = QuerySigningInfosRequest{} }
func (m *QuerySigningInfosRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosRequest) ProtoMessage()    {}
func (*QuerySigningInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{4}
}
func (m *QuerySigningInfosRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosRequest.Merge(m, src)
}
func (m *QuerySigningInfosRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosRequest proto.InternalMessageInfo

func (m *QuerySigningInfosRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
type QuerySigningInfosResponse struct {
	// info is the signing info of all validators
	Info       []ValidatorSigningInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info"`
	Pagination *query.PageResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySigningInfosResponse) Reset()         { *m = QuerySigningInfosResponse{} }
func (m *QuerySigningInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosResponse) ProtoMessage()    {}
func (*QuerySigningInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{5}
}
func (m *QuerySigningInfosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosResponse.Merge(m, src)
}
func (m *QuerySigningInfosResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosResponse proto.InternalMessageInfo

func (m *QuerySigningInfosResponse) GetInfo() []ValidatorSigningInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *QuerySigningInfosResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.slashing.v43beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.slashing.v43beta1.QueryParamsResponse")
	proto.RegisterType((*QuerySigningInfoRequest)(nil), "cosmos.slashing.v43beta1.QuerySigningInfoRequest")
	proto.RegisterType((*QuerySigningInfoResponse)(nil), "cosmos.slashing.v43beta1.QuerySigningInfoResponse")
	proto.RegisterType((*QuerySigningInfosRequest)(nil), "cosmos.slashing.v43beta1.QuerySigningInfosRequest")
	proto.RegisterType((*QuerySigningInfosResponse)(nil), "cosmos.slashing.v43beta1.QuerySigningInfosResponse")
}

func init() {
	proto.RegisterFile("cosmos/slashing/v43beta1/query.proto", fileDescriptor_791b11d41a861ed0)
}

var fileDescriptor_791b11d41a861ed0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x6d, 0x0d, 0x38, 0x29, 0x22, 0x63, 0xa1, 0x35, 0xc8, 0xc6, 0xae, 0x90, 0x16,
	0x35, 0x3b, 0x26, 0x22, 0x5e, 0xec, 0xc1, 0x1e, 0x0c, 0xde, 0x34, 0x8a, 0x07, 0x41, 0xc2, 0x6c,
	0x32, 0x9d, 0x0e, 0x6e, 0x66, 0xb6, 0xfb, 0x36, 0xc1, 0x20, 0x5e, 0x3c, 0x7b, 0x10, 0xfc, 0x0c,
	0x1e, 0x3d, 0xf8, 0x2d, 0x7a, 0x2c, 0x78, 0xf1, 0x24, 0x92, 0xf8, 0x09, 0xfc, 0x04, 0x92, 0x99,
	0x49, 0xb2, 0x25, 0xae, 0xa6, 0x3d, 0xed, 0xf0, 0xe6, 0xfd, 0xdf, 0xff, 0xf7, 0xe6, 0x3d, 0x16,
	0xdf, 0xe8, 0x68, 0xe8, 0x69, 0xa0, 0x10, 0x31, 0x38, 0x94, 0x4a, 0xd0, 0x41, 0x3d, 0xe4, 0x29,
	0xab, 0xd3, 0xa3, 0x3e, 0x4f, 0x86, 0x41, 0x9c, 0xe8, 0x54, 0x93, 0x4d, 0x9b, 0x14, 0x4c, 0x93,
	0x02, 0x97, 0x54, 0xbe, 0xe9, 0xd4, 0x21, 0x03, 0x6e, 0x15, 0x33, 0x7d, 0xcc, 0x84, 0x54, 0x2c,
	0x95, 0x5a, 0xd9, 0x22, 0xe5, 0x0d, 0xa1, 0x85, 0x36, 0x47, 0x3a, 0x39, 0xb9, 0xe8, 0x35, 0xa1,
	0xb5, 0x88, 0x38, 0x65, 0xb1, 0xa4, 0x4c, 0x29, 0x9d, 0x1a, 0x09, 0xb8, 0xdb, 0x6a, 0x1e, 0xdd,
	0x8c, 0xc4, 0xe4, 0xf9, 0x1b, 0x98, 0x3c, 0x9d, 0xb8, 0x3f, 0x61, 0x09, 0xeb, 0x41, 0x8b, 0x1f,
	0xf5, 0x39, 0xa4, 0xfe, 0x73, 0x7c, 0xe5, 0x54, 0x14, 0x62, 0xad, 0x80, 0x93, 0x3d, 0x5c, 0x8c,
	0x4d, 0x64, 0x0b, 0x5d, 0x47, 0xbb, 0xa5, 0x46, 0x25, 0xc8, 0x69, 0x2f, 0xb0, 0xc2, 0xfd, 0xb5,
	0xe3, 0x1f, 0x95, 0x42, 0xcb, 0x89, 0xfc, 0x07, 0x78, 0xd3, 0x54, 0x7d, 0x26, 0x85, 0x92, 0x4a,
	0x3c, 0x56, 0x07, 0xda, 0x19, 0x92, 0x6d, 0xbc, 0xde, 0xd1, 0x0a, 0xda, 0xac, 0xdb, 0x4d, 0x38,
	0xd8, 0xfa, 0x17, 0x5b, 0xa5, 0x49, 0xec, 0xa1, 0x0d, 0xf9, 0x43, 0xbc, 0xb5, 0xa8, 0x76, 0x60,
	0xaf, 0xf0, 0xe5, 0x01, 0x8b, 0xda, 0x60, 0xaf, 0xda, 0x52, 0x1d, 0x68, 0x87, 0x58, 0xcb, 0x45,
	0x7c, 0xc1, 0x22, 0xd9, 0x65, 0xa9, 0x4e, 0x32, 0x05, 0x1d, 0xf0, 0xa5, 0x01, 0x8b, 0x32, 0x51,
	0x3f, 0x5c, 0xb4, 0x9e, 0x3e, 0x15, 0x79, 0x84, 0xf1, 0x7c, 0x60, 0xce, 0xb4, 0x3a, 0x35, 0x9d,
	0x4c, 0x37, 0xb0, 0xfb, 0x30, 0x7f, 0x19, 0xc1, 0x9d, 0xb6, 0x95, 0x51, 0xfa, 0x5f, 0x10, 0xbe,
	0xfa, 0x17, 0x13, 0xd7, 0x60, 0x13, 0xaf, 0xb9, 0xa6, 0x56, 0xcf, 0xdb, 0x94, 0x29, 0x40, 0x9a,
	0xa7, 0x70, 0x57, 0x0c, 0xee, 0xce, 0x7f, 0x71, 0x2d, 0x45, 0x96, 0xb7, 0xf1, 0x7b, 0x15, 0x5f,
	0x30, 0xbc, 0xe4, 0x03, 0xc2, 0x45, 0x3b, 0x6f, 0x72, 0x2b, 0x17, 0x6c, 0x71, 0xc9, 0xca, 0xb7,
	0x97, 0x4b, 0xb6, 0xde, 0xfe, 0xce, 0xfb, 0x6f, 0xbf, 0x3e, 0xad, 0x6c, 0x93, 0x0a, 0xcd, 0xdb,
	0x6c, 0xbb, 0x65, 0xe4, 0x2b, 0xc2, 0xa5, 0x4c, 0xf7, 0xe4, 0xce, 0xbf, 0x6d, 0x16, 0x97, 0xb1,
	0x5c, 0x3f, 0x83, 0xc2, 0xd1, 0xed, 0x19, 0xba, 0xfb, 0xe4, 0x5e, 0x2e, 0x5d, 0x76, 0x37, 0x81,
	0xbe, 0xcd, 0x6e, 0xfb, 0x3b, 0xf2, 0x19, 0xe1, 0xf5, 0xec, 0xdc, 0xc9, 0xf2, 0x08, 0xb3, 0xe7,
	0x6c, 0x9c, 0x45, 0xe2, 0xb0, 0x03, 0x83, 0xbd, 0x4b, 0xaa, 0xcb, 0x61, 0xef, 0x37, 0x8f, 0x47,
	0x1e, 0x3a, 0x19, 0x79, 0xe8, 0xe7, 0xc8, 0x43, 0x1f, 0xc7, 0x5e, 0xe1, 0x64, 0xec, 0x15, 0xbe,
	0x8f, 0xbd, 0xc2, 0xcb, 0x9a, 0x90, 0xe9, 0x61, 0x3f, 0x0c, 0x3a, 0xba, 0x37, 0xad, 0x65, 0x3f,
	0x35, 0xe8, 0xbe, 0xa6, 0x6f, 0xe6, 0x85, 0xd3, 0x61, 0xcc, 0x21, 0x2c, 0x9a, 0xbf, 0xcf, 0xdd,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x0f, 0xc0, 0xe6, 0x45, 0x05, 0x00, 0x00,
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
	// Params queries the parameters of slashing module
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v43beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error) {
	out := new(QuerySigningInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v43beta1.Query/SigningInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error) {
	out := new(QuerySigningInfosResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v43beta1.Query/SigningInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of slashing module
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(context.Context, *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(context.Context, *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SigningInfo(ctx context.Context, req *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfo not implemented")
}
func (*UnimplementedQueryServer) SigningInfos(ctx context.Context, req *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfos not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v43beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v43beta1.Query/SigningInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfo(ctx, req.(*QuerySigningInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v43beta1.Query/SigningInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfos(ctx, req.(*QuerySigningInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.slashing.v43beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SigningInfo",
			Handler:    _Query_SigningInfo_Handler,
		},
		{
			MethodName: "SigningInfos",
			Handler:    _Query_SigningInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/slashing/v43beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsAddress) > 0 {
		i -= len(m.ConsAddress)
		copy(dAtA[i:], m.ConsAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConsAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValSigningInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfosRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QuerySigningInfosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Info) > 0 {
		for iNdEx := len(m.Info) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Info[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySigningInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ValSigningInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfosRequest) Size() (n int) {
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

func (m *QuerySigningInfosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Info) > 0 {
		for _, e := range m.Info {
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySigningInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
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
			m.ConsAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySigningInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValSigningInfo", wireType)
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
			if err := m.ValSigningInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySigningInfosRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfosRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QuerySigningInfosResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySigningInfosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
			m.Info = append(m.Info, ValidatorSigningInfo{})
			if err := m.Info[len(m.Info)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
