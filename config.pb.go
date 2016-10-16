// Code generated by protoc-gen-go.
// source: v2ray.com/core/config.proto
// DO NOT EDIT!

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/config.proto

It has these top-level messages:
	AllocationStrategyConcurrency
	AllocationStrategyRefresh
	AllocationStrategy
	InboundConnectionConfig
	OutboundConnectionConfig
	Config
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_app_router "v2ray.com/core/app/router"
import v2ray_core_app_dns "v2ray.com/core/app/dns"
import v2ray_core_common_loader "v2ray.com/core/common/loader"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net2 "v2ray.com/core/common/net"
import v2ray_core_common_log "v2ray.com/core/common/log"
import v2ray_core_transport_internet "v2ray.com/core/transport/internet"
import v2ray_core_transport "v2ray.com/core/transport"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigFormat int32

const (
	ConfigFormat_Protobuf ConfigFormat = 0
	ConfigFormat_JSON     ConfigFormat = 1
)

var ConfigFormat_name = map[int32]string{
	0: "Protobuf",
	1: "JSON",
}
var ConfigFormat_value = map[string]int32{
	"Protobuf": 0,
	"JSON":     1,
}

func (x ConfigFormat) String() string {
	return proto.EnumName(ConfigFormat_name, int32(x))
}
func (ConfigFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategy_Type int32

const (
	// Always allocate all connection handlers.
	AllocationStrategy_Always AllocationStrategy_Type = 0
	// Randomly allocate specific range of handlers.
	AllocationStrategy_Random AllocationStrategy_Type = 1
	// External. Not supported yet.
	AllocationStrategy_External AllocationStrategy_Type = 2
)

var AllocationStrategy_Type_name = map[int32]string{
	0: "Always",
	1: "Random",
	2: "External",
}
var AllocationStrategy_Type_value = map[string]int32{
	"Always":   0,
	"Random":   1,
	"External": 2,
}

func (x AllocationStrategy_Type) String() string {
	return proto.EnumName(AllocationStrategy_Type_name, int32(x))
}
func (AllocationStrategy_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type AllocationStrategyConcurrency struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyConcurrency) Reset()                    { *m = AllocationStrategyConcurrency{} }
func (m *AllocationStrategyConcurrency) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyConcurrency) ProtoMessage()               {}
func (*AllocationStrategyConcurrency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategyRefresh struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyRefresh) Reset()                    { *m = AllocationStrategyRefresh{} }
func (m *AllocationStrategyRefresh) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyRefresh) ProtoMessage()               {}
func (*AllocationStrategyRefresh) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AllocationStrategy struct {
	Type AllocationStrategy_Type `protobuf:"varint,1,opt,name=type,enum=v2ray.core.AllocationStrategy_Type" json:"type,omitempty"`
	// Number of handlers (ports) running in parallel.
	Concurrency *AllocationStrategyConcurrency `protobuf:"bytes,2,opt,name=concurrency" json:"concurrency,omitempty"`
	// Number of minutes before a handler is regenerated.
	Refresh *AllocationStrategyRefresh `protobuf:"bytes,3,opt,name=refresh" json:"refresh,omitempty"`
}

func (m *AllocationStrategy) Reset()                    { *m = AllocationStrategy{} }
func (m *AllocationStrategy) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategy) ProtoMessage()               {}
func (*AllocationStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AllocationStrategy) GetConcurrency() *AllocationStrategyConcurrency {
	if m != nil {
		return m.Concurrency
	}
	return nil
}

func (m *AllocationStrategy) GetRefresh() *AllocationStrategyRefresh {
	if m != nil {
		return m.Refresh
	}
	return nil
}

// Config for an inbound connection handler.
type InboundConnectionConfig struct {
	Settings               *v2ray_core_common_loader.TypedSettings     `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	PortRange              *v2ray_core_common_net.PortRange            `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	ListenOn               *v2ray_core_common_net2.IPOrDomain          `protobuf:"bytes,3,opt,name=listen_on,json=listenOn" json:"listen_on,omitempty"`
	Tag                    string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	AllocationStrategy     *AllocationStrategy                         `protobuf:"bytes,5,opt,name=allocation_strategy,json=allocationStrategy" json:"allocation_strategy,omitempty"`
	StreamSettings         *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,6,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	AllowPassiveConnection bool                                        `protobuf:"varint,7,opt,name=allow_passive_connection,json=allowPassiveConnection" json:"allow_passive_connection,omitempty"`
}

func (m *InboundConnectionConfig) Reset()                    { *m = InboundConnectionConfig{} }
func (m *InboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*InboundConnectionConfig) ProtoMessage()               {}
func (*InboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InboundConnectionConfig) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *InboundConnectionConfig) GetPortRange() *v2ray_core_common_net.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *InboundConnectionConfig) GetListenOn() *v2ray_core_common_net2.IPOrDomain {
	if m != nil {
		return m.ListenOn
	}
	return nil
}

func (m *InboundConnectionConfig) GetAllocationStrategy() *AllocationStrategy {
	if m != nil {
		return m.AllocationStrategy
	}
	return nil
}

func (m *InboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

type OutboundConnectionConfig struct {
	Settings       *v2ray_core_common_loader.TypedSettings     `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	SendThrough    *v2ray_core_common_net2.IPOrDomain          `protobuf:"bytes,2,opt,name=send_through,json=sendThrough" json:"send_through,omitempty"`
	StreamSettings *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,3,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	Tag            string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
}

func (m *OutboundConnectionConfig) Reset()                    { *m = OutboundConnectionConfig{} }
func (m *OutboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*OutboundConnectionConfig) ProtoMessage()               {}
func (*OutboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OutboundConnectionConfig) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *OutboundConnectionConfig) GetSendThrough() *v2ray_core_common_net2.IPOrDomain {
	if m != nil {
		return m.SendThrough
	}
	return nil
}

func (m *OutboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

type Config struct {
	Inbound   []*InboundConnectionConfig    `protobuf:"bytes,1,rep,name=inbound" json:"inbound,omitempty"`
	Outbound  []*OutboundConnectionConfig   `protobuf:"bytes,2,rep,name=outbound" json:"outbound,omitempty"`
	Log       *v2ray_core_common_log.Config `protobuf:"bytes,3,opt,name=log" json:"log,omitempty"`
	Router    *v2ray_core_app_router.Config `protobuf:"bytes,4,opt,name=router" json:"router,omitempty"`
	Dns       *v2ray_core_app_dns.Config    `protobuf:"bytes,5,opt,name=dns" json:"dns,omitempty"`
	Transport *v2ray_core_transport.Config  `protobuf:"bytes,6,opt,name=transport" json:"transport,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Config) GetInbound() []*InboundConnectionConfig {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Config) GetOutbound() []*OutboundConnectionConfig {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Config) GetLog() *v2ray_core_common_log.Config {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetRouter() *v2ray_core_app_router.Config {
	if m != nil {
		return m.Router
	}
	return nil
}

func (m *Config) GetDns() *v2ray_core_app_dns.Config {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *Config) GetTransport() *v2ray_core_transport.Config {
	if m != nil {
		return m.Transport
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationStrategyConcurrency)(nil), "v2ray.core.AllocationStrategyConcurrency")
	proto.RegisterType((*AllocationStrategyRefresh)(nil), "v2ray.core.AllocationStrategyRefresh")
	proto.RegisterType((*AllocationStrategy)(nil), "v2ray.core.AllocationStrategy")
	proto.RegisterType((*InboundConnectionConfig)(nil), "v2ray.core.InboundConnectionConfig")
	proto.RegisterType((*OutboundConnectionConfig)(nil), "v2ray.core.OutboundConnectionConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.Config")
	proto.RegisterEnum("v2ray.core.ConfigFormat", ConfigFormat_name, ConfigFormat_value)
	proto.RegisterEnum("v2ray.core.AllocationStrategy_Type", AllocationStrategy_Type_name, AllocationStrategy_Type_value)
}

func init() { proto.RegisterFile("v2ray.com/core/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x95, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0xc6, 0x23, 0xdb, 0x71, 0xec, 0xe3, 0x2c, 0x33, 0xb8, 0x61, 0xd3, 0xb2, 0x65, 0xf0, 0x9c,
	0x7f, 0x5e, 0x1b, 0x48, 0xa8, 0x8b, 0xa0, 0x45, 0x81, 0x36, 0x4d, 0x9c, 0x16, 0x48, 0x0b, 0xd4,
	0x06, 0x9d, 0xab, 0xde, 0x18, 0x8c, 0xc4, 0x28, 0x02, 0x24, 0x52, 0x20, 0xe9, 0xa4, 0x7e, 0x84,
	0x5e, 0xf4, 0xe5, 0x7a, 0xdb, 0x97, 0x29, 0x28, 0xd1, 0xf2, 0xff, 0xb4, 0x40, 0xd1, 0x3b, 0x4a,
	0xfc, 0x7e, 0x87, 0xe4, 0xf7, 0x1d, 0x4a, 0xf0, 0xf7, 0x6d, 0x5b, 0x90, 0x91, 0xe3, 0xf1, 0xd8,
	0xf5, 0xb8, 0xa0, 0xae, 0xc7, 0xd9, 0x75, 0x18, 0x38, 0x89, 0xe0, 0x8a, 0x23, 0x18, 0x4f, 0x0a,
	0xba, 0x7d, 0x30, 0x27, 0x24, 0x49, 0xe2, 0x0a, 0x3e, 0x54, 0x54, 0xcc, 0x30, 0xdb, 0xbb, 0x4b,
	0x74, 0x3e, 0x93, 0xb3, 0xa2, 0xc3, 0x85, 0x55, 0xe3, 0x98, 0x33, 0x37, 0xe2, 0xc4, 0xa7, 0xc2,
	0x55, 0xa3, 0x84, 0x1a, 0xe1, 0xde, 0x72, 0x21, 0xa3, 0xca, 0x4d, 0xb8, 0x50, 0xf7, 0x97, 0xd3,
	0x2a, 0xe2, 0xfb, 0x82, 0x4a, 0x69, 0x84, 0x07, 0xab, 0xd6, 0x0d, 0x66, 0xf7, 0xe7, 0xcc, 0xe9,
	0x94, 0x20, 0x4c, 0xea, 0x05, 0xdd, 0x90, 0x29, 0x2a, 0x74, 0xe1, 0x19, 0xfd, 0xfe, 0x4a, 0xfd,
	0xb4, 0xac, 0x79, 0x0c, 0x3b, 0xa7, 0x51, 0xc4, 0x3d, 0xa2, 0x42, 0xce, 0xfa, 0x4a, 0x10, 0x45,
	0x83, 0x51, 0x87, 0x33, 0x6f, 0x28, 0x04, 0x65, 0xde, 0x08, 0xfd, 0x0e, 0xeb, 0xb7, 0x24, 0x1a,
	0x52, 0xdb, 0x6a, 0x58, 0xad, 0x5f, 0x70, 0xf6, 0xd0, 0x7c, 0x04, 0x7f, 0x2d, 0x62, 0x98, 0x5e,
	0x0b, 0x2a, 0x6f, 0x56, 0x20, 0x1f, 0x0b, 0x80, 0x16, 0x19, 0xf4, 0x04, 0x4a, 0xda, 0xdc, 0x54,
	0xbb, 0xd5, 0xde, 0x75, 0x26, 0xf9, 0x3a, 0x8b, 0x6a, 0xe7, 0x72, 0x94, 0x50, 0x9c, 0x02, 0xe8,
	0x2d, 0xd4, 0xbc, 0xc9, 0x3e, 0xed, 0x42, 0xc3, 0x6a, 0xd5, 0xda, 0xff, 0xdf, 0xcf, 0x4f, 0x1d,
	0x0c, 0x4f, 0xd3, 0xe8, 0x04, 0x36, 0x44, 0xb6, 0x7b, 0xbb, 0x98, 0x16, 0xda, 0xbf, 0xbf, 0x90,
	0x39, 0x2a, 0x1e, 0x53, 0xcd, 0x23, 0x28, 0xe9, 0xbd, 0x21, 0x80, 0xf2, 0x69, 0x74, 0x47, 0x46,
	0xb2, 0xbe, 0xa6, 0xc7, 0x98, 0x30, 0x9f, 0xc7, 0x75, 0x0b, 0x6d, 0x42, 0xe5, 0xd5, 0x07, 0x9d,
	0x13, 0x89, 0xea, 0x85, 0xe6, 0xe7, 0x22, 0xfc, 0x79, 0xc1, 0xae, 0xf8, 0x90, 0xf9, 0x1d, 0xce,
	0x18, 0xf5, 0x74, 0xed, 0x4e, 0x9a, 0x0b, 0xea, 0x40, 0x45, 0x52, 0xa5, 0x42, 0x16, 0xc8, 0xd4,
	0x94, 0x5a, 0xfb, 0x70, 0x7a, 0x2f, 0x59, 0x7f, 0x38, 0x59, 0x5f, 0xa6, 0x7e, 0xf8, 0x7d, 0x23,
	0xc7, 0x39, 0x88, 0x4e, 0x00, 0x74, 0xd6, 0x03, 0x41, 0x58, 0x40, 0x8d, 0x37, 0x8d, 0x25, 0x65,
	0x18, 0x55, 0x4e, 0x8f, 0x0b, 0x85, 0xb5, 0x0e, 0x57, 0x93, 0xf1, 0x10, 0xbd, 0x80, 0x6a, 0x14,
	0x4a, 0x45, 0xd9, 0x80, 0x33, 0x63, 0xc9, 0x7f, 0x2b, 0xf8, 0x8b, 0x5e, 0x57, 0x9c, 0xf3, 0x98,
	0x84, 0x0c, 0x57, 0x32, 0xa6, 0xcb, 0x50, 0x1d, 0x8a, 0x8a, 0x04, 0x76, 0xa9, 0x61, 0xb5, 0xaa,
	0x58, 0x0f, 0x51, 0x17, 0x7e, 0x23, 0xb9, 0x8f, 0x03, 0x69, 0x8c, 0xb4, 0xd7, 0xd3, 0xda, 0xff,
	0x7e, 0xc3, 0x6e, 0x44, 0x16, 0x3b, 0xe7, 0x12, 0x7e, 0x95, 0x4a, 0x50, 0x12, 0x0f, 0x72, 0xbf,
	0xca, 0x69, 0xb1, 0x87, 0xd3, 0xc5, 0xf2, 0xbe, 0x77, 0xc6, 0xf7, 0xc4, 0xe9, 0xa7, 0x54, 0x66,
	0x37, 0xde, 0xca, 0x6a, 0x8c, 0x3d, 0x44, 0x4f, 0xc1, 0xd6, 0x6b, 0xdd, 0x0d, 0x12, 0x22, 0x65,
	0x78, 0x4b, 0x07, 0x5e, 0x1e, 0x90, 0xbd, 0xd1, 0xb0, 0x5a, 0x15, 0xfc, 0x47, 0x3a, 0xdf, 0xcb,
	0xa6, 0x27, 0xf1, 0x35, 0x3f, 0x15, 0xc0, 0xee, 0x0e, 0xd5, 0x4f, 0x4c, 0xf5, 0x1c, 0x36, 0x25,
	0x65, 0xfe, 0x40, 0xdd, 0x08, 0x3e, 0x0c, 0x6e, 0x4c, 0xae, 0xdf, 0x91, 0x4b, 0x4d, 0x63, 0x97,
	0x19, 0xb5, 0xcc, 0xb7, 0xe2, 0x8f, 0xfb, 0xb6, 0x10, 0x78, 0xf3, 0x4b, 0x01, 0xca, 0xe6, 0xf4,
	0xcf, 0x61, 0x23, 0xcc, 0xda, 0xdd, 0xb6, 0x1a, 0xc5, 0x56, 0x6d, 0xf6, 0x9e, 0xaf, 0xb8, 0x09,
	0x78, 0xcc, 0xa0, 0x97, 0x50, 0xe1, 0xc6, 0x58, 0xbb, 0x90, 0xf2, 0x7b, 0xd3, 0xfc, 0x2a, 0xd3,
	0x71, 0x4e, 0x21, 0x17, 0x8a, 0x11, 0x0f, 0xcc, 0x39, 0x77, 0x96, 0x3a, 0x1f, 0x38, 0x86, 0xd2,
	0x4a, 0x74, 0x0c, 0xe5, 0xec, 0x57, 0x92, 0x9e, 0x68, 0x8e, 0x21, 0x49, 0xe2, 0x64, 0xb3, 0x63,
	0xc6, 0x88, 0xd1, 0x11, 0x14, 0x7d, 0x26, 0x4d, 0x53, 0x6f, 0xcf, 0x33, 0x3e, 0x93, 0xf9, 0x22,
	0x3e, 0x93, 0xe8, 0x19, 0x54, 0x73, 0x9b, 0x4d, 0xef, 0xfe, 0xb3, 0x3c, 0x03, 0x43, 0x4d, 0xe4,
	0x0f, 0x0e, 0x60, 0x33, 0x7b, 0xf9, 0x9a, 0x8b, 0x98, 0x28, 0xfd, 0x81, 0xe9, 0xe9, 0x2f, 0xfa,
	0xd5, 0xf0, 0xba, 0xbe, 0x86, 0x2a, 0x50, 0x7a, 0xd3, 0xef, 0xbe, 0xab, 0x5b, 0x67, 0xbb, 0xb0,
	0xe5, 0xf1, 0x78, 0xaa, 0xea, 0x59, 0x2d, 0xe3, 0x52, 0xf5, 0xfb, 0x92, 0x7e, 0x75, 0x55, 0x4e,
	0x7f, 0x06, 0x8f, 0xbf, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd4, 0x2c, 0xd9, 0x7b, 0x07, 0x00,
	0x00,
}