// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	Response
	GetRequest
	GetResponse
*/
package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetResponse struct {
	Consignments []*Consignment `protobuf:"bytes,1,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetResponse) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.consignment.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *GetResponse) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xea, 0x30,
	0x10, 0x85, 0x6f, 0xf8, 0xcf, 0x04, 0x5d, 0x74, 0xbd, 0xb8, 0xb5, 0xe8, 0xa2, 0x51, 0x28, 0x12,
	0xab, 0x54, 0xa2, 0x8f, 0x90, 0x05, 0x62, 0x6b, 0x36, 0xdd, 0x15, 0xea, 0x8c, 0x82, 0xa5, 0xc6,
	0x4e, 0x6d, 0x43, 0x9f, 0xad, 0x0f, 0xd2, 0xf7, 0xa9, 0x70, 0x48, 0x71, 0x55, 0xd1, 0xb2, 0xf3,
	0x99, 0x99, 0xcf, 0xe7, 0x8c, 0x65, 0x98, 0x54, 0x5a, 0x59, 0x75, 0xc7, 0x95, 0x34, 0xa2, 0x90,
	0x25, 0x4a, 0xeb, 0x9f, 0x53, 0xd7, 0x25, 0xb4, 0x50, 0x69, 0x29, 0xb8, 0x56, 0xa9, 0xd1, 0xfb,
	0xd4, 0xeb, 0x27, 0x6f, 0x01, 0x44, 0xd9, 0x49, 0x93, 0xbf, 0xd0, 0x12, 0x39, 0x0d, 0xe2, 0x60,
	0x16, 0xb2, 0x96, 0xc8, 0x49, 0x0c, 0x51, 0x8e, 0x86, 0x6b, 0x51, 0x59, 0xa1, 0x24, 0x6d, 0xb9,
	0x86, 0x5f, 0x22, 0xff, 0xa1, 0xf7, 0x8a, 0xa2, 0xd8, 0x5a, 0xda, 0x8e, 0x83, 0x59, 0x97, 0x1d,
	0x15, 0xc9, 0x00, 0xb8, 0x92, 0x76, 0x23, 0x24, 0x6a, 0x43, 0x3b, 0x71, 0x7b, 0x16, 0xcd, 0x27,
	0xe9, 0xb9, 0x20, 0x69, 0xd6, 0xcc, 0x32, 0x0f, 0x23, 0xd7, 0x10, 0xee, 0xd1, 0x18, 0x7c, 0x7e,
	0x14, 0x39, 0xed, 0x3a, 0xf3, 0x41, 0x5d, 0x58, 0xe6, 0x49, 0x09, 0xe1, 0x27, 0xf5, 0x2d, 0xf8,
	0x0d, 0x44, 0x7c, 0x67, 0xac, 0x2a, 0x51, 0x1f, 0xd8, 0x3a, 0x38, 0x34, 0xa5, 0x65, 0x7e, 0xc8,
	0xad, 0xb4, 0x28, 0x84, 0x74, 0xb9, 0x43, 0x76, 0x54, 0xe4, 0x0a, 0xfa, 0x3b, 0x53, 0x43, 0x9d,
	0xba, 0x71, 0x90, 0xce, 0x6e, 0xc0, 0xd0, 0x54, 0x4a, 0x1a, 0x24, 0x14, 0xfa, 0x5c, 0xe3, 0xc6,
	0x62, 0x6d, 0x39, 0x60, 0x8d, 0x24, 0x0b, 0x88, 0xbc, 0xb5, 0x9c, 0x6f, 0x34, 0x9f, 0xfe, 0xb8,
	0x77, 0x73, 0x66, 0x3e, 0x99, 0x0c, 0x01, 0x16, 0x68, 0x19, 0xbe, 0xec, 0xd0, 0xd8, 0xe4, 0x01,
	0x22, 0xa7, 0x8e, 0xfe, 0x4b, 0x18, 0x7a, 0xb3, 0x86, 0x06, 0xee, 0x79, 0x2f, 0xb4, 0xf9, 0x82,
	0xce, 0xdf, 0x03, 0x18, 0xad, 0xb6, 0xa2, 0xaa, 0x84, 0x2c, 0x56, 0xa8, 0xf7, 0x82, 0x23, 0x59,
	0xc3, 0xbf, 0xcc, 0xed, 0xe3, 0x7f, 0x8d, 0xcb, 0x6e, 0x1f, 0x27, 0xe7, 0xc7, 0x9a, 0xf8, 0xc9,
	0x1f, 0xb2, 0x86, 0xd1, 0x02, 0xad, 0xc7, 0x19, 0x72, 0x7b, 0x1e, 0x3c, 0x3d, 0xc4, 0x78, 0xfa,
	0xcb, 0x54, 0xe3, 0xf0, 0xd4, 0x73, 0x5f, 0xff, 0xfe, 0x23, 0x00, 0x00, 0xff, 0xff, 0xae, 0xb2,
	0x23, 0xd1, 0x21, 0x03, 0x00, 0x00,
}
