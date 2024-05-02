// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_nodeinfo.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type OsRegistrationState int32

const (
	OsRegistrationState_notRegistered OsRegistrationState = 0
	OsRegistrationState_registered    OsRegistrationState = 1
)

var OsRegistrationState_name = map[int32]string{
	0: "notRegistered",
	1: "registered",
}

var OsRegistrationState_value = map[string]int32{
	"notRegistered": 0,
	"registered":    1,
}

func (x OsRegistrationState) String() string {
	return proto.EnumName(OsRegistrationState_name, int32(x))
}

func (OsRegistrationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{0}
}

type OsRegistrationStatus struct {
	Status               OsRegistrationState `protobuf:"varint,1,opt,name=status,proto3,enum=moc.common.OsRegistrationState" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OsRegistrationStatus) Reset()         { *m = OsRegistrationStatus{} }
func (m *OsRegistrationStatus) String() string { return proto.CompactTextString(m) }
func (*OsRegistrationStatus) ProtoMessage()    {}
func (*OsRegistrationStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{0}
}

func (m *OsRegistrationStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OsRegistrationStatus.Unmarshal(m, b)
}
func (m *OsRegistrationStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OsRegistrationStatus.Marshal(b, m, deterministic)
}
func (m *OsRegistrationStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsRegistrationStatus.Merge(m, src)
}
func (m *OsRegistrationStatus) XXX_Size() int {
	return xxx_messageInfo_OsRegistrationStatus.Size(m)
}
func (m *OsRegistrationStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_OsRegistrationStatus.DiscardUnknown(m)
}

var xxx_messageInfo_OsRegistrationStatus proto.InternalMessageInfo

func (m *OsRegistrationStatus) GetStatus() OsRegistrationState {
	if m != nil {
		return m.Status
	}
	return OsRegistrationState_notRegistered
}

type Processor struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cores                uint32        `protobuf:"varint,2,opt,name=cores,proto3" json:"cores,omitempty"`
	Speed                string        `protobuf:"bytes,3,opt,name=speed,proto3" json:"speed,omitempty"`
	Type                 ProcessorType `protobuf:"varint,4,opt,name=type,proto3,enum=moc.ProcessorType" json:"type,omitempty"`
	Virtualization       bool          `protobuf:"varint,5,opt,name=virtualization,proto3" json:"virtualization,omitempty"`
	Logicalprocessors    uint32        `protobuf:"varint,6,opt,name=logicalprocessors,proto3" json:"logicalprocessors,omitempty"`
	Hypervisorpresent    bool          `protobuf:"varint,7,opt,name=hypervisorpresent,proto3" json:"hypervisorpresent,omitempty"`
	Manufacturer         string        `protobuf:"bytes,8,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Architecture         Architecture  `protobuf:"varint,9,opt,name=architecture,proto3,enum=moc.Architecture" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Processor) Reset()         { *m = Processor{} }
func (m *Processor) String() string { return proto.CompactTextString(m) }
func (*Processor) ProtoMessage()    {}
func (*Processor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{1}
}

func (m *Processor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Processor.Unmarshal(m, b)
}
func (m *Processor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Processor.Marshal(b, m, deterministic)
}
func (m *Processor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Processor.Merge(m, src)
}
func (m *Processor) XXX_Size() int {
	return xxx_messageInfo_Processor.Size(m)
}
func (m *Processor) XXX_DiscardUnknown() {
	xxx_messageInfo_Processor.DiscardUnknown(m)
}

var xxx_messageInfo_Processor proto.InternalMessageInfo

func (m *Processor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Processor) GetCores() uint32 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *Processor) GetSpeed() string {
	if m != nil {
		return m.Speed
	}
	return ""
}

func (m *Processor) GetType() ProcessorType {
	if m != nil {
		return m.Type
	}
	return ProcessorType_None
}

func (m *Processor) GetVirtualization() bool {
	if m != nil {
		return m.Virtualization
	}
	return false
}

func (m *Processor) GetLogicalprocessors() uint32 {
	if m != nil {
		return m.Logicalprocessors
	}
	return 0
}

func (m *Processor) GetHypervisorpresent() bool {
	if m != nil {
		return m.Hypervisorpresent
	}
	return false
}

func (m *Processor) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *Processor) GetArchitecture() Architecture {
	if m != nil {
		return m.Architecture
	}
	return Architecture_x86
}

type PhysicalMemory struct {
	SizeBytes            uint64   `protobuf:"varint,1,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhysicalMemory) Reset()         { *m = PhysicalMemory{} }
func (m *PhysicalMemory) String() string { return proto.CompactTextString(m) }
func (*PhysicalMemory) ProtoMessage()    {}
func (*PhysicalMemory) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{2}
}

func (m *PhysicalMemory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalMemory.Unmarshal(m, b)
}
func (m *PhysicalMemory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalMemory.Marshal(b, m, deterministic)
}
func (m *PhysicalMemory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalMemory.Merge(m, src)
}
func (m *PhysicalMemory) XXX_Size() int {
	return xxx_messageInfo_PhysicalMemory.Size(m)
}
func (m *PhysicalMemory) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalMemory.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalMemory proto.InternalMessageInfo

func (m *PhysicalMemory) GetSizeBytes() uint64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type HostGPU struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeBytes            uint32         `protobuf:"varint,2,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
	Count                uint32         `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	PartitionSizeBytes   uint64         `protobuf:"varint,4,opt,name=partitionSizeBytes,proto3" json:"partitionSizeBytes,omitempty"`
	Assignment           AssignmentType `protobuf:"varint,5,opt,name=assignment,proto3,enum=moc.AssignmentType" json:"assignment,omitempty"`
	Id                   string         `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HostGPU) Reset()         { *m = HostGPU{} }
func (m *HostGPU) String() string { return proto.CompactTextString(m) }
func (*HostGPU) ProtoMessage()    {}
func (*HostGPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{3}
}

func (m *HostGPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostGPU.Unmarshal(m, b)
}
func (m *HostGPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostGPU.Marshal(b, m, deterministic)
}
func (m *HostGPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostGPU.Merge(m, src)
}
func (m *HostGPU) XXX_Size() int {
	return xxx_messageInfo_HostGPU.Size(m)
}
func (m *HostGPU) XXX_DiscardUnknown() {
	xxx_messageInfo_HostGPU.DiscardUnknown(m)
}

var xxx_messageInfo_HostGPU proto.InternalMessageInfo

func (m *HostGPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HostGPU) GetSizeBytes() uint32 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *HostGPU) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *HostGPU) GetPartitionSizeBytes() uint64 {
	if m != nil {
		return m.PartitionSizeBytes
	}
	return 0
}

func (m *HostGPU) GetAssignment() AssignmentType {
	if m != nil {
		return m.Assignment
	}
	return AssignmentType_GpuDDA
}

func (m *HostGPU) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OperatingSystem struct {
	Operatingsystemsku   uint64                `protobuf:"varint,1,opt,name=operatingsystemsku,proto3" json:"operatingsystemsku,omitempty"`
	Ostype               OperatingSystemType   `protobuf:"varint,2,opt,name=ostype,proto3,enum=moc.OperatingSystemType" json:"ostype,omitempty"`
	Osversion            string                `protobuf:"bytes,3,opt,name=osversion,proto3" json:"osversion,omitempty"`
	OsRegistrationStatus *OsRegistrationStatus `protobuf:"bytes,4,opt,name=osRegistrationStatus,proto3" json:"osRegistrationStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OperatingSystem) Reset()         { *m = OperatingSystem{} }
func (m *OperatingSystem) String() string { return proto.CompactTextString(m) }
func (*OperatingSystem) ProtoMessage()    {}
func (*OperatingSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{4}
}

func (m *OperatingSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatingSystem.Unmarshal(m, b)
}
func (m *OperatingSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatingSystem.Marshal(b, m, deterministic)
}
func (m *OperatingSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatingSystem.Merge(m, src)
}
func (m *OperatingSystem) XXX_Size() int {
	return xxx_messageInfo_OperatingSystem.Size(m)
}
func (m *OperatingSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatingSystem.DiscardUnknown(m)
}

var xxx_messageInfo_OperatingSystem proto.InternalMessageInfo

func (m *OperatingSystem) GetOperatingsystemsku() uint64 {
	if m != nil {
		return m.Operatingsystemsku
	}
	return 0
}

func (m *OperatingSystem) GetOstype() OperatingSystemType {
	if m != nil {
		return m.Ostype
	}
	return OperatingSystemType_WINDOWS
}

func (m *OperatingSystem) GetOsversion() string {
	if m != nil {
		return m.Osversion
	}
	return ""
}

func (m *OperatingSystem) GetOsRegistrationStatus() *OsRegistrationStatus {
	if m != nil {
		return m.OsRegistrationStatus
	}
	return nil
}

type NodeInfo struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Capability           *Resources          `protobuf:"bytes,3,opt,name=capability,proto3" json:"capability,omitempty"`
	Availability         *Resources          `protobuf:"bytes,4,opt,name=availability,proto3" json:"availability,omitempty"`
	Ostype               OperatingSystemType `protobuf:"varint,6,opt,name=ostype,proto3,enum=moc.OperatingSystemType" json:"ostype,omitempty"`
	Status               *Status             `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Uptime               int64               `protobuf:"varint,8,opt,name=uptime,proto3" json:"uptime,omitempty"`
	OsInfo               *OperatingSystem    `protobuf:"bytes,9,opt,name=osInfo,proto3" json:"osInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{5}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeInfo) GetCapability() *Resources {
	if m != nil {
		return m.Capability
	}
	return nil
}

func (m *NodeInfo) GetAvailability() *Resources {
	if m != nil {
		return m.Availability
	}
	return nil
}

func (m *NodeInfo) GetOstype() OperatingSystemType {
	if m != nil {
		return m.Ostype
	}
	return OperatingSystemType_WINDOWS
}

func (m *NodeInfo) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *NodeInfo) GetUptime() int64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *NodeInfo) GetOsInfo() *OperatingSystem {
	if m != nil {
		return m.OsInfo
	}
	return nil
}

type Resources struct {
	Processor            *Processor                  `protobuf:"bytes,1,opt,name=processor,proto3" json:"processor,omitempty"`
	Memory               *PhysicalMemory             `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	VmCapabilities       *VirtualMachineCapabilities `protobuf:"bytes,5,opt,name=vmCapabilities,proto3" json:"vmCapabilities,omitempty"`
	HostGPUs             []*HostGPU                  `protobuf:"bytes,7,rep,name=hostGPUs,proto3" json:"hostGPUs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c83f03f7e6831a3, []int{6}
}

func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (m *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(m, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetProcessor() *Processor {
	if m != nil {
		return m.Processor
	}
	return nil
}

func (m *Resources) GetMemory() *PhysicalMemory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Resources) GetVmCapabilities() *VirtualMachineCapabilities {
	if m != nil {
		return m.VmCapabilities
	}
	return nil
}

func (m *Resources) GetHostGPUs() []*HostGPU {
	if m != nil {
		return m.HostGPUs
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.common.OsRegistrationState", OsRegistrationState_name, OsRegistrationState_value)
	proto.RegisterType((*OsRegistrationStatus)(nil), "moc.common.OsRegistrationStatus")
	proto.RegisterType((*Processor)(nil), "moc.common.Processor")
	proto.RegisterType((*PhysicalMemory)(nil), "moc.common.PhysicalMemory")
	proto.RegisterType((*HostGPU)(nil), "moc.common.HostGPU")
	proto.RegisterType((*OperatingSystem)(nil), "moc.common.OperatingSystem")
	proto.RegisterType((*NodeInfo)(nil), "moc.common.NodeInfo")
	proto.RegisterType((*Resources)(nil), "moc.common.Resources")
}

func init() { proto.RegisterFile("moc_common_nodeinfo.proto", fileDescriptor_7c83f03f7e6831a3) }

var fileDescriptor_7c83f03f7e6831a3 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xfe, 0x9b, 0xa6, 0x69, 0x7b, 0xba, 0xf5, 0xef, 0xbc, 0x01, 0x61, 0x20, 0xa8, 0x32, 0x31,
	0x55, 0x08, 0xb5, 0xa8, 0xd5, 0x04, 0x5c, 0x32, 0x2e, 0x06, 0x93, 0xc6, 0x26, 0x6f, 0x70, 0xc1,
	0xcd, 0x94, 0xa5, 0x6e, 0x6b, 0xd1, 0xc4, 0x91, 0xed, 0x54, 0xca, 0xae, 0x78, 0x38, 0xc4, 0x03,
	0xf0, 0x12, 0xbc, 0x06, 0xca, 0x49, 0x9a, 0x26, 0x5d, 0x35, 0x71, 0xd5, 0xf8, 0x3b, 0xdf, 0xf1,
	0xf9, 0xfc, 0x1d, 0x1f, 0x17, 0x1e, 0xfb, 0xc2, 0xbb, 0xf6, 0x84, 0xef, 0x8b, 0xe0, 0x3a, 0x10,
	0x63, 0xc6, 0x83, 0x89, 0xe8, 0x87, 0x52, 0x68, 0x41, 0xc0, 0x17, 0x5e, 0x3f, 0x0d, 0xed, 0x3f,
	0x2a, 0xd0, 0xd2, 0x9f, 0x94, 0xb4, 0xff, 0xac, 0x1c, 0x08, 0x23, 0xcd, 0x8a, 0x71, 0xe7, 0x1c,
	0xf6, 0xce, 0x15, 0x65, 0x53, 0xae, 0xb4, 0x74, 0x35, 0x17, 0xc1, 0xa5, 0x76, 0x75, 0xa4, 0xc8,
	0x1b, 0xb0, 0x14, 0x7e, 0xd9, 0x95, 0x6e, 0xa5, 0xd7, 0x1e, 0x3e, 0xef, 0xaf, 0xaa, 0xf5, 0xef,
	0x66, 0x30, 0x9a, 0xd1, 0x9d, 0xdf, 0x06, 0x34, 0x2f, 0xa4, 0xf0, 0x98, 0x52, 0x42, 0x12, 0x02,
	0x66, 0xe0, 0xfa, 0x0c, 0x37, 0x69, 0x52, 0xfc, 0x26, 0x7b, 0x50, 0xf3, 0x84, 0x64, 0xca, 0x36,
	0xba, 0x95, 0xde, 0x36, 0x4d, 0x17, 0x09, 0xaa, 0x42, 0xc6, 0xc6, 0x76, 0x15, 0xa9, 0xe9, 0x82,
	0x1c, 0x82, 0xa9, 0xe3, 0x90, 0xd9, 0x26, 0x8a, 0x20, 0x28, 0x22, 0xdf, 0xfd, 0x2a, 0x0e, 0x19,
	0xc5, 0x38, 0x39, 0x84, 0xf6, 0x82, 0x4b, 0x1d, 0xb9, 0x73, 0x7e, 0x8b, 0xa2, 0xec, 0x5a, 0xb7,
	0xd2, 0x6b, 0xd0, 0x35, 0x94, 0xbc, 0x82, 0x9d, 0xb9, 0x98, 0x72, 0xcf, 0x9d, 0x87, 0xcb, 0x5d,
	0x94, 0x6d, 0xa1, 0x8e, 0xbb, 0x81, 0x84, 0x3d, 0x8b, 0x43, 0x26, 0x17, 0x5c, 0x09, 0x19, 0x4a,
	0xa6, 0x58, 0xa0, 0xed, 0x3a, 0x6e, 0x7c, 0x37, 0x40, 0x1c, 0xd8, 0xf2, 0xdd, 0x20, 0x9a, 0xb8,
	0x9e, 0x8e, 0x24, 0x93, 0x76, 0x03, 0x0f, 0x52, 0xc2, 0xc8, 0x11, 0x6c, 0xb9, 0xd2, 0x9b, 0x71,
	0xcd, 0x10, 0xb0, 0x9b, 0x78, 0xae, 0x1d, 0x3c, 0xd7, 0xfb, 0x42, 0x80, 0x96, 0x68, 0x4e, 0x1f,
	0xda, 0x17, 0xb3, 0x58, 0x25, 0xf2, 0xce, 0x98, 0x2f, 0x64, 0x4c, 0x9e, 0x42, 0x53, 0xf1, 0x5b,
	0x76, 0x1c, 0x6b, 0x96, 0xb6, 0xc8, 0xa4, 0x2b, 0xc0, 0xf9, 0x55, 0x81, 0xfa, 0x47, 0xa1, 0xf4,
	0xc9, 0xc5, 0x97, 0x8d, 0x2d, 0x28, 0x65, 0xa7, 0x6d, 0x58, 0x01, 0x69, 0x83, 0xa2, 0x40, 0x63,
	0x2b, 0xb0, 0x41, 0x51, 0xa0, 0x49, 0x1f, 0x48, 0xe8, 0x4a, 0xcd, 0xb1, 0xe5, 0x79, 0xb2, 0x89,
	0xa5, 0x37, 0x44, 0xc8, 0x08, 0xc0, 0x55, 0x8a, 0x4f, 0x03, 0x3f, 0x71, 0xad, 0x86, 0x07, 0xdd,
	0x4d, 0x0f, 0x9a, 0xc3, 0xd8, 0xc1, 0x02, 0x8d, 0xb4, 0xc1, 0xe0, 0x63, 0x6c, 0x48, 0x93, 0x1a,
	0x7c, 0xec, 0xfc, 0xa9, 0xc0, 0xff, 0xe7, 0x21, 0x4b, 0x2e, 0x5a, 0x30, 0xbd, 0x8c, 0x95, 0x66,
	0x7e, 0x22, 0x44, 0x2c, 0x21, 0x85, 0x90, 0xfa, 0x1e, 0x65, 0x1e, 0x6c, 0x88, 0x90, 0xd7, 0x60,
	0x09, 0x85, 0xb7, 0xc8, 0x40, 0x11, 0x36, 0x8a, 0x58, 0xdb, 0x15, 0x95, 0x64, 0xbc, 0xc4, 0x1e,
	0xa1, 0x16, 0x4c, 0xaa, 0xe4, 0x22, 0xa5, 0xf7, 0x71, 0x05, 0x90, 0x2b, 0xd8, 0x13, 0x1b, 0x46,
	0x06, 0xad, 0x68, 0x0d, 0xbb, 0xf7, 0x0f, 0x4a, 0xa4, 0xe8, 0xc6, 0x6c, 0xe7, 0xa7, 0x01, 0x8d,
	0xcf, 0x62, 0xcc, 0x3e, 0x05, 0x13, 0xb1, 0xb1, 0x67, 0xa9, 0x35, 0xc6, 0xd2, 0x1a, 0x72, 0x04,
	0xe0, 0xb9, 0xa1, 0x7b, 0xc3, 0xe7, 0x5c, 0xc7, 0xa8, 0xb2, 0x35, 0x7c, 0x50, 0x2c, 0x4e, 0x99,
	0x12, 0x91, 0xf4, 0x98, 0xa2, 0x05, 0x22, 0x79, 0x07, 0x5b, 0xee, 0xc2, 0xe5, 0xf3, 0x65, 0xa2,
	0x79, 0x5f, 0x62, 0x89, 0x5a, 0x30, 0xd2, 0xfa, 0x47, 0x23, 0x0f, 0xf2, 0x57, 0xa4, 0x8e, 0x65,
	0x5a, 0x98, 0x91, 0xf9, 0x90, 0x85, 0xc8, 0x43, 0xb0, 0xa2, 0x50, 0x73, 0x9f, 0xe1, 0xc4, 0x54,
	0x69, 0xb6, 0x22, 0xa3, 0xa4, 0x5c, 0x62, 0x07, 0x4e, 0x49, 0x6b, 0xf8, 0xa4, 0xe4, 0x6c, 0xb9,
	0x2a, 0xcd, 0xa8, 0xce, 0x0f, 0x03, 0x9a, 0xb9, 0x7e, 0x32, 0x82, 0x66, 0x3e, 0xce, 0x68, 0xe6,
	0xda, 0x49, 0xf3, 0xa7, 0x84, 0xae, 0x78, 0x64, 0x08, 0x96, 0x8f, 0x43, 0x86, 0x66, 0xb7, 0x86,
	0xfb, 0xa5, 0x8c, 0xd2, 0x18, 0xd2, 0x8c, 0x49, 0x4e, 0xa0, 0xbd, 0xf0, 0x3f, 0x2c, 0x5d, 0xe6,
	0x4c, 0xe1, 0x85, 0x6f, 0x65, 0xcf, 0xe6, 0xd7, 0xf4, 0x11, 0x3a, 0x73, 0xbd, 0x19, 0x0f, 0x58,
	0x91, 0x46, 0xd7, 0xd2, 0xc8, 0x00, 0x1a, 0xb3, 0x74, 0x70, 0x13, 0xcf, 0xaa, 0xbd, 0x56, 0x36,
	0x33, 0x59, 0xf9, 0x6c, 0xa8, 0x69, 0x4e, 0x3a, 0x35, 0x1b, 0xd5, 0x8e, 0x79, 0x6a, 0x36, 0xcc,
	0x4e, 0xed, 0xd4, 0x6c, 0x58, 0x9d, 0xfa, 0xcb, 0xb7, 0xb0, 0xbb, 0xe1, 0x81, 0x26, 0x3b, 0xb0,
	0x1d, 0x08, 0x9d, 0xe2, 0x4c, 0xb2, 0x71, 0xe7, 0x3f, 0xd2, 0x06, 0x90, 0xab, 0x75, 0xe5, 0xf8,
	0xc5, 0xb7, 0x83, 0x29, 0xd7, 0xb3, 0xe8, 0x26, 0x29, 0x37, 0xf0, 0xb9, 0x27, 0x85, 0x12, 0x13,
	0x3d, 0xf0, 0x85, 0x37, 0x90, 0xa1, 0x37, 0x48, 0x45, 0xdc, 0x58, 0xf8, 0xd7, 0x31, 0xfa, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x85, 0x80, 0xb0, 0x18, 0x9c, 0x06, 0x00, 0x00,
}
