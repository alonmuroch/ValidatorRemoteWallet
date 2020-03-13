// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/v1alpha1/RemoteWallet.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FetchValidatingKeysResponse struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchValidatingKeysResponse) Reset()         { *m = FetchValidatingKeysResponse{} }
func (m *FetchValidatingKeysResponse) String() string { return proto.CompactTextString(m) }
func (*FetchValidatingKeysResponse) ProtoMessage()    {}
func (*FetchValidatingKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43296aad5947550b, []int{0}
}

func (m *FetchValidatingKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchValidatingKeysResponse.Unmarshal(m, b)
}
func (m *FetchValidatingKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchValidatingKeysResponse.Marshal(b, m, deterministic)
}
func (m *FetchValidatingKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchValidatingKeysResponse.Merge(m, src)
}
func (m *FetchValidatingKeysResponse) XXX_Size() int {
	return xxx_messageInfo_FetchValidatingKeysResponse.Size(m)
}
func (m *FetchValidatingKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchValidatingKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchValidatingKeysResponse proto.InternalMessageInfo

func (m *FetchValidatingKeysResponse) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type FetchValidatingKeysRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchValidatingKeysRequest) Reset()         { *m = FetchValidatingKeysRequest{} }
func (m *FetchValidatingKeysRequest) String() string { return proto.CompactTextString(m) }
func (*FetchValidatingKeysRequest) ProtoMessage()    {}
func (*FetchValidatingKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43296aad5947550b, []int{1}
}

func (m *FetchValidatingKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchValidatingKeysRequest.Unmarshal(m, b)
}
func (m *FetchValidatingKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchValidatingKeysRequest.Marshal(b, m, deterministic)
}
func (m *FetchValidatingKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchValidatingKeysRequest.Merge(m, src)
}
func (m *FetchValidatingKeysRequest) XXX_Size() int {
	return xxx_messageInfo_FetchValidatingKeysRequest.Size(m)
}
func (m *FetchValidatingKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchValidatingKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchValidatingKeysRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FetchValidatingKeysResponse)(nil), "ethereum.wallet.v1alpha1.FetchValidatingKeysResponse")
	proto.RegisterType((*FetchValidatingKeysRequest)(nil), "ethereum.wallet.v1alpha1.FetchValidatingKeysRequest")
}

func init() {
	proto.RegisterFile("wallet/v1alpha1/RemoteWallet.proto", fileDescriptor_43296aad5947550b)
}

var fileDescriptor_43296aad5947550b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4f, 0xcc, 0xc9,
	0x49, 0x2d, 0xd1, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x0f, 0x4a, 0xcd, 0xcd,
	0x2f, 0x49, 0x0d, 0x07, 0x8b, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0xa4, 0x96, 0x64,
	0xa4, 0x16, 0xa5, 0x96, 0xe6, 0xea, 0x41, 0x14, 0xeb, 0xc1, 0x14, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7,
	0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64,
	0xe6, 0xe7, 0x15, 0x43, 0xf4, 0x29, 0xd9, 0x71, 0x49, 0xbb, 0xa5, 0x96, 0x24, 0x67, 0x84, 0x25,
	0xe6, 0x64, 0xa6, 0x24, 0x96, 0x64, 0xe6, 0xa5, 0x7b, 0xa7, 0x56, 0x16, 0x07, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x73, 0x71, 0x17, 0x94, 0x26, 0xe5, 0x64, 0x26, 0xc7, 0x67,
	0xa7, 0x56, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x04, 0x71, 0x41, 0x84, 0x40, 0x0a, 0x95,
	0x64, 0xb8, 0xa4, 0xb0, 0xea, 0x2f, 0x2c, 0x4d, 0x2d, 0x2e, 0x31, 0xda, 0xc6, 0xc8, 0xc5, 0x83,
	0xec, 0x58, 0xa1, 0xa5, 0x8c, 0x5c, 0xc2, 0x58, 0xd4, 0x0b, 0x99, 0xe8, 0xe1, 0x72, 0xbf, 0x1e,
	0x6e, 0xe3, 0xa5, 0x4c, 0x49, 0xd4, 0x05, 0xf1, 0x94, 0x92, 0x62, 0xd3, 0xe5, 0x27, 0x93, 0x99,
	0xa4, 0x85, 0x24, 0xf5, 0xd1, 0x03, 0x16, 0xe4, 0x49, 0xfd, 0xc4, 0x9c, 0x1c, 0xa7, 0x3a, 0x2e,
	0x99, 0xfc, 0xa2, 0x74, 0x9c, 0xc6, 0x3b, 0x09, 0x22, 0xfb, 0x2a, 0x00, 0x14, 0x92, 0x01, 0x8c,
	0x51, 0xc2, 0x68, 0xc6, 0x59, 0xa7, 0x96, 0x64, 0xac, 0x62, 0x92, 0x70, 0x85, 0x19, 0x12, 0x8e,
	0x6a, 0xc8, 0x29, 0x84, 0x54, 0x0c, 0x44, 0x2a, 0x06, 0x26, 0x95, 0xc4, 0x06, 0x8e, 0x1d, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x0c, 0x29, 0x2a, 0xfb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RemoteWalletClient is the client API for RemoteWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteWalletClient interface {
	FetchValidatingKeys(ctx context.Context, in *FetchValidatingKeysRequest, opts ...grpc.CallOption) (*FetchValidatingKeysResponse, error)
}

type remoteWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteWalletClient(cc grpc.ClientConnInterface) RemoteWalletClient {
	return &remoteWalletClient{cc}
}

func (c *remoteWalletClient) FetchValidatingKeys(ctx context.Context, in *FetchValidatingKeysRequest, opts ...grpc.CallOption) (*FetchValidatingKeysResponse, error) {
	out := new(FetchValidatingKeysResponse)
	err := c.cc.Invoke(ctx, "/ethereum.wallet.v1alpha1.RemoteWallet/FetchValidatingKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteWalletServer is the server API for RemoteWallet service.
type RemoteWalletServer interface {
	FetchValidatingKeys(context.Context, *FetchValidatingKeysRequest) (*FetchValidatingKeysResponse, error)
}

// UnimplementedRemoteWalletServer can be embedded to have forward compatible implementations.
type UnimplementedRemoteWalletServer struct {
}

func (*UnimplementedRemoteWalletServer) FetchValidatingKeys(ctx context.Context, req *FetchValidatingKeysRequest) (*FetchValidatingKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchValidatingKeys not implemented")
}

func RegisterRemoteWalletServer(s *grpc.Server, srv RemoteWalletServer) {
	s.RegisterService(&_RemoteWallet_serviceDesc, srv)
}

func _RemoteWallet_FetchValidatingKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchValidatingKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteWalletServer).FetchValidatingKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.wallet.v1alpha1.RemoteWallet/FetchValidatingKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteWalletServer).FetchValidatingKeys(ctx, req.(*FetchValidatingKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteWallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.wallet.v1alpha1.RemoteWallet",
	HandlerType: (*RemoteWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchValidatingKeys",
			Handler:    _RemoteWallet_FetchValidatingKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/v1alpha1/RemoteWallet.proto",
}
