// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vippool.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type VipPoolRequest struct {
	VipPools             []*VipPool       `protobuf:"bytes,1,rep,name=VipPools,proto3" json:"VipPools,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *VipPoolRequest) Reset()         { *m = VipPoolRequest{} }
func (m *VipPoolRequest) String() string { return proto.CompactTextString(m) }
func (*VipPoolRequest) ProtoMessage()    {}
func (*VipPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34da3589aa67c0c5, []int{0}
}

func (m *VipPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VipPoolRequest.Unmarshal(m, b)
}
func (m *VipPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VipPoolRequest.Marshal(b, m, deterministic)
}
func (m *VipPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VipPoolRequest.Merge(m, src)
}
func (m *VipPoolRequest) XXX_Size() int {
	return xxx_messageInfo_VipPoolRequest.Size(m)
}
func (m *VipPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VipPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VipPoolRequest proto.InternalMessageInfo

func (m *VipPoolRequest) GetVipPools() []*VipPool {
	if m != nil {
		return m.VipPools
	}
	return nil
}

func (m *VipPoolRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VipPoolResponse struct {
	VipPools             []*VipPool          `protobuf:"bytes,1,rep,name=VipPools,proto3" json:"VipPools,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VipPoolResponse) Reset()         { *m = VipPoolResponse{} }
func (m *VipPoolResponse) String() string { return proto.CompactTextString(m) }
func (*VipPoolResponse) ProtoMessage()    {}
func (*VipPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34da3589aa67c0c5, []int{1}
}

func (m *VipPoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VipPoolResponse.Unmarshal(m, b)
}
func (m *VipPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VipPoolResponse.Marshal(b, m, deterministic)
}
func (m *VipPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VipPoolResponse.Merge(m, src)
}
func (m *VipPoolResponse) XXX_Size() int {
	return xxx_messageInfo_VipPoolResponse.Size(m)
}
func (m *VipPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VipPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VipPoolResponse proto.InternalMessageInfo

func (m *VipPoolResponse) GetVipPools() []*VipPool {
	if m != nil {
		return m.VipPools
	}
	return nil
}

func (m *VipPoolResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VipPoolResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VipPool struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cidr                 string         `protobuf:"bytes,3,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Networkid            string         `protobuf:"bytes,4,opt,name=networkid,proto3" json:"networkid,omitempty"`
	Nodefqdn             string         `protobuf:"bytes,5,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName            string         `protobuf:"bytes,6,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName         string         `protobuf:"bytes,7,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status               *common.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Startip              string         `protobuf:"bytes,9,opt,name=startip,proto3" json:"startip,omitempty"`
	Endip                string         `protobuf:"bytes,10,opt,name=endip,proto3" json:"endip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VipPool) Reset()         { *m = VipPool{} }
func (m *VipPool) String() string { return proto.CompactTextString(m) }
func (*VipPool) ProtoMessage()    {}
func (*VipPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_34da3589aa67c0c5, []int{2}
}

func (m *VipPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VipPool.Unmarshal(m, b)
}
func (m *VipPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VipPool.Marshal(b, m, deterministic)
}
func (m *VipPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VipPool.Merge(m, src)
}
func (m *VipPool) XXX_Size() int {
	return xxx_messageInfo_VipPool.Size(m)
}
func (m *VipPool) XXX_DiscardUnknown() {
	xxx_messageInfo_VipPool.DiscardUnknown(m)
}

var xxx_messageInfo_VipPool proto.InternalMessageInfo

func (m *VipPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VipPool) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VipPool) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *VipPool) GetNetworkid() string {
	if m != nil {
		return m.Networkid
	}
	return ""
}

func (m *VipPool) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *VipPool) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VipPool) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VipPool) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VipPool) GetStartip() string {
	if m != nil {
		return m.Startip
	}
	return ""
}

func (m *VipPool) GetEndip() string {
	if m != nil {
		return m.Endip
	}
	return ""
}

func init() {
	proto.RegisterType((*VipPoolRequest)(nil), "moc.cloudagent.network.VipPoolRequest")
	proto.RegisterType((*VipPoolResponse)(nil), "moc.cloudagent.network.VipPoolResponse")
	proto.RegisterType((*VipPool)(nil), "moc.cloudagent.network.VipPool")
}

func init() { proto.RegisterFile("vippool.proto", fileDescriptor_34da3589aa67c0c5) }

var fileDescriptor_34da3589aa67c0c5 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x37, 0xdd, 0xdd, 0xb4, 0x7d, 0xed, 0x56, 0x18, 0x44, 0x86, 0x20, 0x5a, 0x22, 0x68, 0x2f,
	0x26, 0x10, 0x15, 0x0f, 0x9e, 0x5c, 0xf0, 0xe0, 0x45, 0x25, 0xca, 0x1e, 0xf4, 0x94, 0x4e, 0xa6,
	0x71, 0x68, 0x32, 0x6f, 0x76, 0x66, 0xb2, 0x8b, 0x67, 0xbf, 0x81, 0x27, 0x3f, 0xae, 0xe4, 0x25,
	0xcd, 0xb2, 0x20, 0xf4, 0xe0, 0x6d, 0xde, 0xfb, 0xfd, 0x79, 0x8f, 0xdf, 0x1b, 0x38, 0xbf, 0x56,
	0xc6, 0x20, 0xd6, 0x89, 0xb1, 0xe8, 0x91, 0x3d, 0x6c, 0x50, 0x24, 0xa2, 0xc6, 0xb6, 0x2c, 0x2a,
	0xa9, 0x7d, 0xa2, 0xa5, 0xbf, 0x41, 0xbb, 0x8f, 0x1e, 0x57, 0x88, 0x55, 0x2d, 0x53, 0x62, 0x6d,
	0xdb, 0x5d, 0x7a, 0x63, 0x0b, 0x63, 0xa4, 0x75, 0xbd, 0x2e, 0x5a, 0x0a, 0x6c, 0x1a, 0xd4, 0x7d,
	0x15, 0xff, 0x0a, 0x60, 0x75, 0xa9, 0xcc, 0x67, 0xc4, 0x3a, 0x97, 0x57, 0xad, 0x74, 0x9e, 0xbd,
	0x85, 0xd9, 0xd0, 0x71, 0x3c, 0x58, 0x9f, 0x6c, 0x16, 0xd9, 0x93, 0xe4, 0xdf, 0xb3, 0x92, 0x83,
	0x72, 0x14, 0xb0, 0x57, 0x70, 0xfe, 0xc9, 0x48, 0x5b, 0x78, 0x85, 0xfa, 0xeb, 0x4f, 0x23, 0xf9,
	0x64, 0x1d, 0x6c, 0x56, 0xd9, 0x8a, 0x1c, 0x46, 0x24, 0xbf, 0x4b, 0x8a, 0xff, 0x04, 0x70, 0x7f,
	0xdc, 0xc2, 0x19, 0xd4, 0x4e, 0xfe, 0xdf, 0x1a, 0x19, 0x84, 0xb9, 0x74, 0x6d, 0xed, 0x69, 0xfe,
	0x22, 0x8b, 0x92, 0x3e, 0x95, 0xe4, 0x90, 0x4a, 0x72, 0x81, 0x58, 0x5f, 0x16, 0x75, 0x2b, 0xf3,
	0x81, 0xc9, 0x1e, 0xc0, 0xd9, 0x7b, 0x6b, 0xd1, 0xf2, 0x93, 0x75, 0xb0, 0x99, 0xe7, 0x7d, 0x11,
	0xff, 0x9e, 0xc0, 0x74, 0xb0, 0x65, 0x0c, 0x4e, 0x75, 0xd1, 0x48, 0x1e, 0x10, 0x81, 0xde, 0x6c,
	0x05, 0x13, 0x55, 0xd2, 0x94, 0x79, 0x3e, 0x51, 0x65, 0xc7, 0x11, 0xaa, 0x3c, 0x98, 0xd0, 0x9b,
	0x3d, 0x82, 0xf9, 0xb0, 0xaa, 0x2a, 0xf9, 0x29, 0x01, 0xb7, 0x0d, 0x16, 0xc1, 0x4c, 0x63, 0x29,
	0x77, 0x57, 0xa5, 0xe6, 0x67, 0x04, 0x8e, 0x75, 0xa7, 0xac, 0x2c, 0xb6, 0xe6, 0x63, 0x37, 0x36,
	0xec, 0x95, 0x63, 0x83, 0xc5, 0xb0, 0xac, 0x51, 0x50, 0x8c, 0x44, 0x98, 0x12, 0xe1, 0x4e, 0x8f,
	0x3d, 0x85, 0xd0, 0xf9, 0xc2, 0xb7, 0x8e, 0xcf, 0x28, 0x89, 0x05, 0x85, 0xf8, 0x85, 0x5a, 0xf9,
	0x00, 0x31, 0x0e, 0x53, 0xe7, 0x0b, 0xeb, 0x95, 0xe1, 0x73, 0xf2, 0x38, 0x94, 0x5d, 0x28, 0x52,
	0x97, 0xca, 0x70, 0xe8, 0x43, 0xa1, 0x22, 0xdb, 0xc3, 0x72, 0xc8, 0xe4, 0x5d, 0x77, 0x08, 0xf6,
	0x1d, 0xc2, 0x0f, 0xfa, 0x1a, 0xf7, 0x92, 0x3d, 0x3b, 0x76, 0xa3, 0xfe, 0x93, 0x45, 0xcf, 0x8f,
	0xf2, 0xfa, 0x6f, 0x10, 0xdf, 0xbb, 0x78, 0xf3, 0xed, 0x75, 0xa5, 0xfc, 0x8f, 0x76, 0x9b, 0x08,
	0x6c, 0xd2, 0x46, 0x09, 0x8b, 0x0e, 0x77, 0x3e, 0x6d, 0x50, 0xbc, 0xa0, 0x73, 0xa6, 0xd6, 0x88,
	0xf4, 0xd6, 0x2a, 0x1d, 0xac, 0xb6, 0x21, 0xa1, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x06,
	0x52, 0x9b, 0xf6, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VipPoolAgentClient is the client API for VipPoolAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VipPoolAgentClient interface {
	Invoke(ctx context.Context, in *VipPoolRequest, opts ...grpc.CallOption) (*VipPoolResponse, error)
}

type vipPoolAgentClient struct {
	cc *grpc.ClientConn
}

func NewVipPoolAgentClient(cc *grpc.ClientConn) VipPoolAgentClient {
	return &vipPoolAgentClient{cc}
}

func (c *vipPoolAgentClient) Invoke(ctx context.Context, in *VipPoolRequest, opts ...grpc.CallOption) (*VipPoolResponse, error) {
	out := new(VipPoolResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.VipPoolAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VipPoolAgentServer is the server API for VipPoolAgent service.
type VipPoolAgentServer interface {
	Invoke(context.Context, *VipPoolRequest) (*VipPoolResponse, error)
}

// UnimplementedVipPoolAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVipPoolAgentServer struct {
}

func (*UnimplementedVipPoolAgentServer) Invoke(ctx context.Context, req *VipPoolRequest) (*VipPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVipPoolAgentServer(s *grpc.Server, srv VipPoolAgentServer) {
	s.RegisterService(&_VipPoolAgent_serviceDesc, srv)
}

func _VipPoolAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VipPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VipPoolAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.VipPoolAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VipPoolAgentServer).Invoke(ctx, req.(*VipPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VipPoolAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.VipPoolAgent",
	HandlerType: (*VipPoolAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VipPoolAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vippool.proto",
}
