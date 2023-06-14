// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_virtualmachinescaleset.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	network "github.com/microsoft/moc/rpc/cloudagent/network"
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

type VirtualMachineScaleSetRequest struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	OperationType                 common.Operation          `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetRequest) Reset()         { *m = VirtualMachineScaleSetRequest{} }
func (m *VirtualMachineScaleSetRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetRequest) ProtoMessage()    {}
func (*VirtualMachineScaleSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{0}
}

func (m *VirtualMachineScaleSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetRequest.Merge(m, src)
}
func (m *VirtualMachineScaleSetRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetRequest.Size(m)
}
func (m *VirtualMachineScaleSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetRequest proto.InternalMessageInfo

func (m *VirtualMachineScaleSetRequest) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualMachineScaleSetResponse struct {
	VirtualMachineScaleSetSystems []*VirtualMachineScaleSet `protobuf:"bytes,1,rep,name=VirtualMachineScaleSetSystems,proto3" json:"VirtualMachineScaleSetSystems,omitempty"`
	Result                        *wrappers.BoolValue       `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                         string                    `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                  `json:"-"`
	XXX_unrecognized              []byte                    `json:"-"`
	XXX_sizecache                 int32                     `json:"-"`
}

func (m *VirtualMachineScaleSetResponse) Reset()         { *m = VirtualMachineScaleSetResponse{} }
func (m *VirtualMachineScaleSetResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSetResponse) ProtoMessage()    {}
func (*VirtualMachineScaleSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{1}
}

func (m *VirtualMachineScaleSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSetResponse.Merge(m, src)
}
func (m *VirtualMachineScaleSetResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSetResponse.Size(m)
}
func (m *VirtualMachineScaleSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSetResponse proto.InternalMessageInfo

func (m *VirtualMachineScaleSetResponse) GetVirtualMachineScaleSetSystems() []*VirtualMachineScaleSet {
	if m != nil {
		return m.VirtualMachineScaleSetSystems
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualMachineScaleSetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Sku struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             int64    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sku) Reset()         { *m = Sku{} }
func (m *Sku) String() string { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()    {}
func (*Sku) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{2}
}

func (m *Sku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sku.Unmarshal(m, b)
}
func (m *Sku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sku.Marshal(b, m, deterministic)
}
func (m *Sku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sku.Merge(m, src)
}
func (m *Sku) XXX_Size() int {
	return xxx_messageInfo_Sku.Size(m)
}
func (m *Sku) XXX_DiscardUnknown() {
	xxx_messageInfo_Sku.DiscardUnknown(m)
}

var xxx_messageInfo_Sku proto.InternalMessageInfo

func (m *Sku) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sku) GetCapacity() int64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type NetworkConfigurationScaleSet struct {
	Interfaces           []*network.NetworkInterface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NetworkConfigurationScaleSet) Reset()         { *m = NetworkConfigurationScaleSet{} }
func (m *NetworkConfigurationScaleSet) String() string { return proto.CompactTextString(m) }
func (*NetworkConfigurationScaleSet) ProtoMessage()    {}
func (*NetworkConfigurationScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{3}
}

func (m *NetworkConfigurationScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Unmarshal(m, b)
}
func (m *NetworkConfigurationScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Marshal(b, m, deterministic)
}
func (m *NetworkConfigurationScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfigurationScaleSet.Merge(m, src)
}
func (m *NetworkConfigurationScaleSet) XXX_Size() int {
	return xxx_messageInfo_NetworkConfigurationScaleSet.Size(m)
}
func (m *NetworkConfigurationScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfigurationScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfigurationScaleSet proto.InternalMessageInfo

func (m *NetworkConfigurationScaleSet) GetInterfaces() []*network.NetworkInterface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type VirtualMachineProfile struct {
	Vmprefix             string                        `protobuf:"bytes,1,opt,name=vmprefix,proto3" json:"vmprefix,omitempty"`
	Network              *NetworkConfigurationScaleSet `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Storage              *StorageConfiguration         `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Os                   *OperatingSystemConfiguration `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Hardware             *HardwareConfiguration        `protobuf:"bytes,5,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Security             *SecurityConfiguration        `protobuf:"bytes,6,opt,name=security,proto3" json:"security,omitempty"`
	GuestAgent           *GuestAgentConfiguration      `protobuf:"bytes,7,opt,name=guestAgent,proto3" json:"guestAgent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *VirtualMachineProfile) Reset()         { *m = VirtualMachineProfile{} }
func (m *VirtualMachineProfile) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineProfile) ProtoMessage()    {}
func (*VirtualMachineProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{4}
}

func (m *VirtualMachineProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineProfile.Unmarshal(m, b)
}
func (m *VirtualMachineProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineProfile.Marshal(b, m, deterministic)
}
func (m *VirtualMachineProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineProfile.Merge(m, src)
}
func (m *VirtualMachineProfile) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineProfile.Size(m)
}
func (m *VirtualMachineProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineProfile.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineProfile proto.InternalMessageInfo

func (m *VirtualMachineProfile) GetVmprefix() string {
	if m != nil {
		return m.Vmprefix
	}
	return ""
}

func (m *VirtualMachineProfile) GetNetwork() *NetworkConfigurationScaleSet {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *VirtualMachineProfile) GetStorage() *StorageConfiguration {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *VirtualMachineProfile) GetOs() *OperatingSystemConfiguration {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *VirtualMachineProfile) GetHardware() *HardwareConfiguration {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *VirtualMachineProfile) GetSecurity() *SecurityConfiguration {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *VirtualMachineProfile) GetGuestAgent() *GuestAgentConfiguration {
	if m != nil {
		return m.GuestAgent
	}
	return nil
}

type VirtualMachineScaleSet struct {
	Name                    string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                      string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Sku                     *Sku                   `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Virtualmachineprofile   *VirtualMachineProfile `protobuf:"bytes,4,opt,name=virtualmachineprofile,proto3" json:"virtualmachineprofile,omitempty"`
	VirtualMachineSystems   []*VirtualMachine      `protobuf:"bytes,5,rep,name=VirtualMachineSystems,proto3" json:"VirtualMachineSystems,omitempty"`
	Nodefqdn                string                 `protobuf:"bytes,6,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName               string                 `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	LocationName            string                 `protobuf:"bytes,8,opt,name=locationName,proto3" json:"locationName,omitempty"`
	Status                  *common.Status         `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	DisableHighAvailability bool                   `protobuf:"varint,11,opt,name=disableHighAvailability,proto3" json:"disableHighAvailability,omitempty"`
	AllowedOwnerNodes       []string               `protobuf:"bytes,12,rep,name=allowedOwnerNodes,proto3" json:"allowedOwnerNodes,omitempty"`
	Tags                    *common.Tags           `protobuf:"bytes,13,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}               `json:"-"`
	XXX_unrecognized        []byte                 `json:"-"`
	XXX_sizecache           int32                  `json:"-"`
}

func (m *VirtualMachineScaleSet) Reset()         { *m = VirtualMachineScaleSet{} }
func (m *VirtualMachineScaleSet) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineScaleSet) ProtoMessage()    {}
func (*VirtualMachineScaleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ced78740362b34c9, []int{5}
}

func (m *VirtualMachineScaleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineScaleSet.Unmarshal(m, b)
}
func (m *VirtualMachineScaleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineScaleSet.Marshal(b, m, deterministic)
}
func (m *VirtualMachineScaleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineScaleSet.Merge(m, src)
}
func (m *VirtualMachineScaleSet) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineScaleSet.Size(m)
}
func (m *VirtualMachineScaleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineScaleSet.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineScaleSet proto.InternalMessageInfo

func (m *VirtualMachineScaleSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetSku() *Sku {
	if m != nil {
		return m.Sku
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualmachineprofile() *VirtualMachineProfile {
	if m != nil {
		return m.Virtualmachineprofile
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetVirtualMachineSystems() []*VirtualMachine {
	if m != nil {
		return m.VirtualMachineSystems
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualMachineScaleSet) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetDisableHighAvailability() bool {
	if m != nil {
		return m.DisableHighAvailability
	}
	return false
}

func (m *VirtualMachineScaleSet) GetAllowedOwnerNodes() []string {
	if m != nil {
		return m.AllowedOwnerNodes
	}
	return nil
}

func (m *VirtualMachineScaleSet) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*VirtualMachineScaleSetRequest)(nil), "moc.cloudagent.compute.VirtualMachineScaleSetRequest")
	proto.RegisterType((*VirtualMachineScaleSetResponse)(nil), "moc.cloudagent.compute.VirtualMachineScaleSetResponse")
	proto.RegisterType((*Sku)(nil), "moc.cloudagent.compute.Sku")
	proto.RegisterType((*NetworkConfigurationScaleSet)(nil), "moc.cloudagent.compute.NetworkConfigurationScaleSet")
	proto.RegisterType((*VirtualMachineProfile)(nil), "moc.cloudagent.compute.VirtualMachineProfile")
	proto.RegisterType((*VirtualMachineScaleSet)(nil), "moc.cloudagent.compute.VirtualMachineScaleSet")
}

func init() {
	proto.RegisterFile("moc_cloudagent_virtualmachinescaleset.proto", fileDescriptor_ced78740362b34c9)
}

var fileDescriptor_ced78740362b34c9 = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0x4b, 0x6e, 0xe3, 0x36,
	0x18, 0xc7, 0x2b, 0xdb, 0xf1, 0x83, 0x4e, 0x02, 0x94, 0x68, 0x12, 0xc1, 0x79, 0xc0, 0x50, 0xd0,
	0xc2, 0x40, 0x13, 0x09, 0x75, 0x93, 0xa2, 0xdb, 0xa4, 0xaf, 0x64, 0x51, 0xa7, 0xa0, 0x83, 0x2c,
	0x8a, 0x02, 0x01, 0x2d, 0xd1, 0x32, 0x61, 0x49, 0x54, 0x48, 0xca, 0xae, 0x6f, 0x32, 0x47, 0x99,
	0x3b, 0xcc, 0x25, 0xe6, 0x18, 0xb3, 0x1c, 0x88, 0xa2, 0x9c, 0xc8, 0x63, 0x19, 0xc9, 0x6a, 0x56,
	0x36, 0xf9, 0xfd, 0xbf, 0x1f, 0xbf, 0x17, 0x29, 0xf0, 0x63, 0xc8, 0xdc, 0x47, 0x37, 0x60, 0x89,
	0x87, 0x7d, 0x12, 0xc9, 0xc7, 0x19, 0xe5, 0x32, 0xc1, 0x41, 0x88, 0xdd, 0x09, 0x8d, 0x88, 0x70,
	0x71, 0x40, 0x04, 0x91, 0x76, 0xcc, 0x99, 0x64, 0x70, 0x3f, 0x64, 0xae, 0xfd, 0x2c, 0xb6, 0x5d,
	0x16, 0xc6, 0x89, 0x24, 0x9d, 0x13, 0x9f, 0x31, 0x3f, 0x20, 0x8e, 0x52, 0x8d, 0x92, 0xb1, 0x33,
	0xe7, 0x38, 0x8e, 0x09, 0x17, 0x99, 0x5f, 0xe7, 0x74, 0xe3, 0x21, 0x5a, 0x74, 0xa0, 0x44, 0x2c,
	0x0c, 0x59, 0xa4, 0x7f, 0xb4, 0xe1, 0xfb, 0x15, 0xef, 0x88, 0xc8, 0x39, 0xe3, 0x53, 0x1a, 0x49,
	0xc2, 0xc7, 0xd8, 0xd5, 0xfe, 0xd6, 0x07, 0x03, 0x1c, 0x3f, 0x64, 0xe0, 0xbf, 0x33, 0xf0, 0x30,
	0x8d, 0x7e, 0x48, 0x24, 0x22, 0x4f, 0x09, 0x11, 0x12, 0xca, 0x32, 0xc1, 0x70, 0x21, 0x24, 0x09,
	0x85, 0x69, 0x74, 0xab, 0xbd, 0x76, 0xdf, 0xb6, 0xd7, 0xa7, 0x69, 0x97, 0xd0, 0x37, 0x43, 0xe1,
	0x05, 0xd8, 0xb9, 0x8b, 0x09, 0xc7, 0x92, 0xb2, 0xe8, 0x7e, 0x11, 0x13, 0xb3, 0xd2, 0x35, 0x7a,
	0xbb, 0xfd, 0x5d, 0x75, 0xca, 0xd2, 0x82, 0x8a, 0x22, 0xeb, 0xa3, 0x01, 0x4e, 0xca, 0xb2, 0x11,
	0x31, 0x8b, 0x04, 0xf9, 0x4a, 0xe9, 0xf4, 0x41, 0x1d, 0x11, 0x91, 0x04, 0x52, 0xe5, 0xd1, 0xee,
	0x77, 0xec, 0xac, 0xf9, 0x76, 0xde, 0x7c, 0xfb, 0x9a, 0xb1, 0xe0, 0x01, 0x07, 0x09, 0x41, 0x5a,
	0x09, 0xbf, 0x03, 0x5b, 0x7f, 0x70, 0xce, 0xb8, 0x59, 0xed, 0x1a, 0xbd, 0x16, 0xca, 0x16, 0xd6,
	0x25, 0xa8, 0x0e, 0xa7, 0x09, 0x84, 0xa0, 0x16, 0xe1, 0x90, 0x98, 0x86, 0xb2, 0xa9, 0xff, 0xb0,
	0x03, 0x9a, 0x2e, 0x8e, 0xb1, 0x4b, 0xe5, 0x42, 0x1d, 0x53, 0x45, 0xcb, 0xb5, 0x35, 0x01, 0x47,
	0x83, 0x6c, 0x02, 0x7e, 0x63, 0xd1, 0x98, 0xfa, 0x49, 0x56, 0xb5, 0x3c, 0x4e, 0x78, 0x03, 0xc0,
	0x72, 0x34, 0xf2, 0x1a, 0xf4, 0x56, 0x6b, 0xa0, 0x67, 0xc8, 0xd6, 0xa4, 0xdb, 0xdc, 0x01, 0xbd,
	0xf0, 0xb5, 0x3e, 0x55, 0xc1, 0x5e, 0xb1, 0x18, 0xff, 0x70, 0x36, 0xa6, 0x81, 0x8a, 0x6f, 0x16,
	0xc6, 0x9c, 0x8c, 0xe9, 0xff, 0x3a, 0xee, 0xe5, 0x1a, 0x0e, 0x40, 0x43, 0xd3, 0x75, 0x85, 0x2e,
	0xca, 0x1a, 0xb0, 0x29, 0x0d, 0x94, 0x43, 0xe0, 0x9f, 0xa0, 0x21, 0x24, 0xe3, 0xd8, 0x27, 0xaa,
	0x7c, 0xed, 0xfe, 0x59, 0x19, 0x6f, 0x98, 0xc9, 0x0a, 0x3c, 0x94, 0x3b, 0xc3, 0xdf, 0x41, 0x85,
	0x09, 0xb3, 0xb6, 0x39, 0x24, 0x3d, 0x84, 0x91, 0x9f, 0xb5, 0xbb, 0x88, 0xaa, 0x30, 0x01, 0x6f,
	0x41, 0x73, 0x82, 0xb9, 0x37, 0xc7, 0x9c, 0x98, 0x5b, 0x8a, 0x75, 0x5e, 0xc6, 0xba, 0xd1, 0xba,
	0x22, 0x64, 0xe9, 0x9e, 0xa2, 0x04, 0x71, 0x13, 0x9e, 0x36, 0xb9, 0xbe, 0x19, 0x35, 0xd4, 0xba,
	0x15, 0x54, 0xee, 0x0e, 0xef, 0x00, 0xf0, 0xd3, 0x2b, 0x7e, 0x95, 0x7a, 0x99, 0x0d, 0x05, 0x73,
	0xca, 0x60, 0x7f, 0x2d, 0x95, 0x45, 0xdc, 0x0b, 0x84, 0xf5, 0xbe, 0x06, 0xf6, 0xd7, 0xdf, 0x83,
	0xb5, 0xf3, 0xba, 0x0b, 0x2a, 0xd4, 0x53, 0xed, 0x6e, 0xa1, 0x0a, 0xf5, 0xe0, 0x39, 0xa8, 0x8a,
	0x69, 0xa2, 0xfb, 0x75, 0x58, 0x9a, 0xd5, 0x34, 0x41, 0xa9, 0x0e, 0xba, 0x60, 0xaf, 0xf8, 0x24,
	0xc6, 0xd9, 0x9c, 0xe9, 0x6e, 0x9d, 0xbf, 0xee, 0x06, 0xeb, 0xe1, 0x44, 0xeb, 0x59, 0xf0, 0xbf,
	0xd5, 0x61, 0xce, 0x9f, 0x89, 0x2d, 0x75, 0x45, 0x7e, 0x78, 0xdd, 0x21, 0x68, 0x3d, 0x24, 0xbd,
	0x11, 0x11, 0xf3, 0xc8, 0xf8, 0xc9, 0x8b, 0x54, 0x33, 0x5b, 0x68, 0xb9, 0x86, 0x47, 0xa0, 0xe5,
	0x73, 0x96, 0xc4, 0x83, 0xb4, 0x6c, 0x0d, 0x65, 0x7c, 0xde, 0x80, 0x16, 0xd8, 0x0e, 0x98, 0xab,
	0x5a, 0xa0, 0x04, 0x4d, 0x25, 0x28, 0xec, 0xc1, 0x53, 0x50, 0x17, 0x12, 0xcb, 0x44, 0x98, 0x2d,
	0x55, 0x91, 0xb6, 0x0a, 0x76, 0xa8, 0xb6, 0x90, 0x36, 0xc1, 0x5f, 0xc1, 0x81, 0x47, 0x05, 0x1e,
	0x05, 0xe4, 0x86, 0xfa, 0x93, 0xab, 0x19, 0xa6, 0x01, 0x1e, 0xd1, 0x20, 0x1d, 0xaf, 0x76, 0xd7,
	0xe8, 0x35, 0x51, 0x99, 0x19, 0x9e, 0x81, 0x6f, 0x71, 0x10, 0xb0, 0x39, 0xf1, 0xee, 0xe6, 0x11,
	0xe1, 0x03, 0xe6, 0x11, 0x61, 0x6e, 0x77, 0xab, 0xbd, 0x16, 0xfa, 0xd2, 0x00, 0x8f, 0x41, 0x4d,
	0x62, 0x5f, 0x98, 0x3b, 0x2a, 0x94, 0x96, 0x0a, 0xe5, 0x1e, 0xfb, 0x02, 0xa9, 0xed, 0xfe, 0x3b,
	0x03, 0x1c, 0xae, 0x1f, 0x1d, 0x35, 0x5a, 0x70, 0x01, 0xea, 0xb7, 0xd1, 0x8c, 0x4d, 0x09, 0xbc,
	0x7c, 0xe3, 0xcb, 0x9c, 0x7d, 0xc6, 0x3a, 0xbf, 0xbc, 0xd5, 0x2d, 0xfb, 0x5e, 0x58, 0xdf, 0x5c,
	0xff, 0xf4, 0xaf, 0xe3, 0x53, 0x39, 0x49, 0x46, 0xa9, 0x8b, 0x13, 0x52, 0x97, 0x33, 0xc1, 0xc6,
	0xd2, 0x09, 0x99, 0xeb, 0xf0, 0xd8, 0x75, 0x9e, 0x99, 0x8e, 0x66, 0x8e, 0xea, 0xea, 0x59, 0xff,
	0xf9, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xe2, 0xfb, 0x9a, 0x28, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineScaleSetAgentClient is the client API for VirtualMachineScaleSetAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineScaleSetAgentClient interface {
	Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error)
}

type virtualMachineScaleSetAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineScaleSetAgentClient(cc *grpc.ClientConn) VirtualMachineScaleSetAgentClient {
	return &virtualMachineScaleSetAgentClient{cc}
}

func (c *virtualMachineScaleSetAgentClient) Invoke(ctx context.Context, in *VirtualMachineScaleSetRequest, opts ...grpc.CallOption) (*VirtualMachineScaleSetResponse, error) {
	out := new(VirtualMachineScaleSetResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.compute.VirtualMachineScaleSetAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineScaleSetAgentServer is the server API for VirtualMachineScaleSetAgent service.
type VirtualMachineScaleSetAgentServer interface {
	Invoke(context.Context, *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error)
}

// UnimplementedVirtualMachineScaleSetAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineScaleSetAgentServer struct {
}

func (*UnimplementedVirtualMachineScaleSetAgentServer) Invoke(ctx context.Context, req *VirtualMachineScaleSetRequest) (*VirtualMachineScaleSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualMachineScaleSetAgentServer(s *grpc.Server, srv VirtualMachineScaleSetAgentServer) {
	s.RegisterService(&_VirtualMachineScaleSetAgent_serviceDesc, srv)
}

func _VirtualMachineScaleSetAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineScaleSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.compute.VirtualMachineScaleSetAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineScaleSetAgentServer).Invoke(ctx, req.(*VirtualMachineScaleSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineScaleSetAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.compute.VirtualMachineScaleSetAgent",
	HandlerType: (*VirtualMachineScaleSetAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualMachineScaleSetAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_virtualmachinescaleset.proto",
}
