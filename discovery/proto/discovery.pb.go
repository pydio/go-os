// Code generated by protoc-gen-go.
// source: github.com/pydio/go-os/discovery/proto/discovery.proto
// DO NOT EDIT!

/*
Package go_micro_os_discovery is a generated protocol buffer package.

It is generated from these files:
	github.com/pydio/go-os/discovery/proto/discovery.proto

It has these top-level messages:
	Heartbeat
	Service
	Result
	Node
	Endpoint
	Value
*/
package go_micro_os_discovery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Heartbeat struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Service   *Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Timestamp int64    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Interval  int64    `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	Ttl       int64    `protobuf:"varint,5,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Heartbeat) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Service struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,3,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Endpoints []*Endpoint       `protobuf:"bytes,4,rep,name=endpoints" json:"endpoints,omitempty"`
	Nodes     []*Node           `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Result struct {
	Action    string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Service   *Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Timestamp int64    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Result) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Node struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port     int64             `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Endpoint struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Request  *Value            `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
	Response *Value            `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Endpoint) GetRequest() *Value {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Endpoint) GetResponse() *Value {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Endpoint) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Value struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type   string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Values []*Value `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Value) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Heartbeat)(nil), "go.micro.os.discovery.Heartbeat")
	proto.RegisterType((*Service)(nil), "go.micro.os.discovery.Service")
	proto.RegisterType((*Result)(nil), "go.micro.os.discovery.Result")
	proto.RegisterType((*Node)(nil), "go.micro.os.discovery.Node")
	proto.RegisterType((*Endpoint)(nil), "go.micro.os.discovery.Endpoint")
	proto.RegisterType((*Value)(nil), "go.micro.os.discovery.Value")
}

func init() {
	proto.RegisterFile("github.com/pydio/go-os/discovery/proto/discovery.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd2, 0xcf, 0x0c, 0x02, 0xa1, 0x11, 0xa0, 0x68, 0x59, 0x01, 0xea, 0x09, 0x24, 0x36,
	0x15, 0x1f, 0x5a, 0xad, 0x40, 0x1c, 0x2b, 0x2d, 0x07, 0x38, 0x18, 0x89, 0xbb, 0x9b, 0x8c, 0x4a,
	0x44, 0x13, 0x07, 0xdb, 0x5d, 0xd1, 0x7f, 0xc2, 0x95, 0x7f, 0xc0, 0xcf, 0xe0, 0x67, 0x31, 0x76,
	0x9c, 0x16, 0xb4, 0x6d, 0xf7, 0x80, 0x7a, 0x9b, 0xb1, 0xdf, 0x9b, 0xbe, 0xf7, 0xc6, 0x0d, 0x9c,
	0x2f, 0x4a, 0xfb, 0x65, 0x35, 0xcf, 0x72, 0x55, 0x4d, 0xab, 0x32, 0xd7, 0x6a, 0xba, 0x50, 0x67,
	0xca, 0x4c, 0x8b, 0xd2, 0xe4, 0xea, 0x8a, 0xf4, 0x7a, 0xda, 0x68, 0x65, 0xd5, 0xb6, 0xcf, 0x7c,
	0x8f, 0xf7, 0x17, 0x2a, 0xf3, 0xf8, 0x4c, 0x99, 0x6c, 0x73, 0x39, 0xf9, 0x19, 0x41, 0x72, 0x49,
	0x52, 0xdb, 0x39, 0x49, 0x8b, 0x77, 0x20, 0x2e, 0x8b, 0x34, 0x7a, 0x12, 0x3d, 0x4d, 0x04, 0x57,
	0x78, 0x01, 0x23, 0x43, 0xfa, 0xaa, 0xcc, 0x29, 0x8d, 0xf9, 0xf0, 0xd6, 0xcb, 0x47, 0xd9, 0xce,
	0x31, 0xd9, 0xa7, 0x16, 0x25, 0x3a, 0x38, 0x9e, 0x42, 0x62, 0xcb, 0x8a, 0x8c, 0x95, 0x55, 0x93,
	0xf6, 0x98, 0xdb, 0x13, 0xdb, 0x03, 0x3c, 0x81, 0x71, 0x59, 0x5b, 0x86, 0xca, 0x65, 0xda, 0xf7,
	0x97, 0x9b, 0x1e, 0xef, 0x42, 0xcf, 0xda, 0x65, 0x3a, 0xf0, 0xc7, 0xae, 0x9c, 0xfc, 0x8a, 0x61,
	0x14, 0x7e, 0x00, 0x11, 0xfa, 0xb5, 0xac, 0x28, 0x68, 0xf4, 0x35, 0xa6, 0x30, 0x62, 0x11, 0xa6,
	0x54, 0xb5, 0x57, 0x99, 0x88, 0xae, 0xc5, 0x4b, 0x18, 0x57, 0x64, 0x65, 0x21, 0xad, 0x64, 0x11,
	0x3d, 0x36, 0xf0, 0xfc, 0xb0, 0x81, 0xec, 0x43, 0x80, 0xcf, 0x6a, 0xab, 0xd7, 0x62, 0xc3, 0xc6,
	0x77, 0x90, 0x50, 0x5d, 0x34, 0x8a, 0x65, 0x1a, 0x96, 0xec, 0x46, 0x3d, 0xde, 0x33, 0x6a, 0x16,
	0x70, 0x62, 0xcb, 0xc0, 0x17, 0x30, 0xa8, 0x55, 0x41, 0x86, 0x6d, 0x39, 0xea, 0xc3, 0x3d, 0xd4,
	0x8f, 0x8c, 0x11, 0x2d, 0xf2, 0xe4, 0x2d, 0xdc, 0xfe, 0x47, 0x8c, 0x0b, 0xe6, 0x2b, 0xad, 0x83,
	0x73, 0x57, 0xe2, 0x3d, 0x18, 0x70, 0x62, 0x2b, 0x0a, 0xb6, 0xdb, 0xe6, 0x4d, 0x7c, 0x11, 0x4d,
	0xbe, 0xc3, 0x50, 0x90, 0x59, 0x2d, 0x2d, 0x3e, 0x80, 0xa1, 0xcc, 0xad, 0xcb, 0xa6, 0x25, 0x86,
	0xee, 0x58, 0xab, 0x9d, 0xfc, 0x8e, 0xa0, 0xef, 0x6c, 0x5c, 0x7b, 0x4b, 0xbc, 0x25, 0x59, 0x14,
	0x9a, 0x8c, 0xe9, 0xb6, 0x14, 0x5a, 0xb7, 0xd3, 0x46, 0x69, 0x1b, 0x66, 0xf9, 0x1a, 0x67, 0x7f,
	0x6d, 0xae, 0x8d, 0xfb, 0xd9, 0x81, 0xcc, 0xf6, 0xad, 0xed, 0xff, 0x42, 0xfc, 0x11, 0xc3, 0xb8,
	0x5b, 0xe6, 0xce, 0x87, 0x77, 0x0e, 0x23, 0x4d, 0xdf, 0x56, 0xec, 0x3c, 0x64, 0x78, 0xba, 0x47,
	0xe3, 0x67, 0x37, 0x53, 0x74, 0x60, 0xce, 0x7e, 0xcc, 0xc6, 0x1b, 0x55, 0x1b, 0xf2, 0xa6, 0x6f,
	0x22, 0x6e, 0xd0, 0xf8, 0xfe, 0x5a, 0x2c, 0x67, 0x37, 0xbc, 0xc2, 0xe3, 0x44, 0x43, 0x30, 0xf0,
	0xd2, 0x76, 0xc6, 0xc2, 0x67, 0x76, 0xdd, 0x74, 0x2c, 0x5f, 0xe3, 0x6b, 0x18, 0x7a, 0xb6, 0x09,
	0xff, 0xc3, 0xc3, 0x86, 0x03, 0x76, 0x3e, 0xf4, 0xdf, 0xae, 0x57, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x58, 0xbf, 0xf0, 0xf5, 0x04, 0x00, 0x00,
}
