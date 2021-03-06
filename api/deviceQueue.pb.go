// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deviceQueue.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnqueueDeviceQueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort used (must be >0)
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EnqueueDeviceQueueItemRequest) Reset()                    { *m = EnqueueDeviceQueueItemRequest{} }
func (m *EnqueueDeviceQueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemRequest) ProtoMessage()               {}
func (*EnqueueDeviceQueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EnqueueDeviceQueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *EnqueueDeviceQueueItemRequest) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *EnqueueDeviceQueueItemRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *EnqueueDeviceQueueItemRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EnqueueDeviceQueueItemResponse struct {
}

func (m *EnqueueDeviceQueueItemResponse) Reset()                    { *m = EnqueueDeviceQueueItemResponse{} }
func (m *EnqueueDeviceQueueItemResponse) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDeviceQueueItemResponse) ProtoMessage()               {}
func (*EnqueueDeviceQueueItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type DeleteDeviceQueueItemRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// ID of the queue item to delete.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteDeviceQueueItemRequest) Reset()                    { *m = DeleteDeviceQueueItemRequest{} }
func (m *DeleteDeviceQueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDeviceQueueItemRequest) ProtoMessage()               {}
func (*DeleteDeviceQueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DeleteDeviceQueueItemRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DeleteDeviceQueueItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteDeviceQueueItemResponse struct {
}

func (m *DeleteDeviceQueueItemResponse) Reset()                    { *m = DeleteDeviceQueueItemResponse{} }
func (m *DeleteDeviceQueueItemResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDeviceQueueItemResponse) ProtoMessage()               {}
func (*DeleteDeviceQueueItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type DeviceQueueItem struct {
	// ID of the queue item.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// Random reference (used on ack notification).
	Reference string `protobuf:"bytes,3,opt,name=reference" json:"reference,omitempty"`
	// Is an ACK required from the node.
	Confirmed bool `protobuf:"varint,4,opt,name=confirmed" json:"confirmed,omitempty"`
	// Transmission is pending (waiting for an ack).
	Pending bool `protobuf:"varint,5,opt,name=pending" json:"pending,omitempty"`
	// FPort used (must be >0).
	FPort uint32 `protobuf:"varint,6,opt,name=fPort" json:"fPort,omitempty"`
	// Base64 encoded data.
	Data []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DeviceQueueItem) Reset()                    { *m = DeviceQueueItem{} }
func (m *DeviceQueueItem) String() string            { return proto.CompactTextString(m) }
func (*DeviceQueueItem) ProtoMessage()               {}
func (*DeviceQueueItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *DeviceQueueItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeviceQueueItem) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *DeviceQueueItem) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DeviceQueueItem) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *DeviceQueueItem) GetPending() bool {
	if m != nil {
		return m.Pending
	}
	return false
}

func (m *DeviceQueueItem) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *DeviceQueueItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListDeviceQueueItemsRequest struct {
	// Hex encoded DevEUI of the node.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *ListDeviceQueueItemsRequest) Reset()                    { *m = ListDeviceQueueItemsRequest{} }
func (m *ListDeviceQueueItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsRequest) ProtoMessage()               {}
func (*ListDeviceQueueItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListDeviceQueueItemsRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

type ListDeviceQueueItemsResponse struct {
	Items []*DeviceQueueItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListDeviceQueueItemsResponse) Reset()                    { *m = ListDeviceQueueItemsResponse{} }
func (m *ListDeviceQueueItemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceQueueItemsResponse) ProtoMessage()               {}
func (*ListDeviceQueueItemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ListDeviceQueueItemsResponse) GetItems() []*DeviceQueueItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EnqueueDeviceQueueItemRequest)(nil), "api.EnqueueDeviceQueueItemRequest")
	proto.RegisterType((*EnqueueDeviceQueueItemResponse)(nil), "api.EnqueueDeviceQueueItemResponse")
	proto.RegisterType((*DeleteDeviceQueueItemRequest)(nil), "api.DeleteDeviceQueueItemRequest")
	proto.RegisterType((*DeleteDeviceQueueItemResponse)(nil), "api.DeleteDeviceQueueItemResponse")
	proto.RegisterType((*DeviceQueueItem)(nil), "api.DeviceQueueItem")
	proto.RegisterType((*ListDeviceQueueItemsRequest)(nil), "api.ListDeviceQueueItemsRequest")
	proto.RegisterType((*ListDeviceQueueItemsResponse)(nil), "api.ListDeviceQueueItemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceQueue service

type DeviceQueueClient interface {
	// Enqueue adds the given item to the queue. When the node operates in
	// Class-C mode, the data will be pushed directly to the network-server.
	Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(ctx context.Context, in *DeleteDeviceQueueItemRequest, opts ...grpc.CallOption) (*DeleteDeviceQueueItemResponse, error)
	// List lists the items in the queue for the given node.
	List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error)
}

type deviceQueueClient struct {
	cc *grpc.ClientConn
}

func NewDeviceQueueClient(cc *grpc.ClientConn) DeviceQueueClient {
	return &deviceQueueClient{cc}
}

func (c *deviceQueueClient) Enqueue(ctx context.Context, in *EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDeviceQueueItemResponse, error) {
	out := new(EnqueueDeviceQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/Enqueue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) Delete(ctx context.Context, in *DeleteDeviceQueueItemRequest, opts ...grpc.CallOption) (*DeleteDeviceQueueItemResponse, error) {
	out := new(DeleteDeviceQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceQueueClient) List(ctx context.Context, in *ListDeviceQueueItemsRequest, opts ...grpc.CallOption) (*ListDeviceQueueItemsResponse, error) {
	out := new(ListDeviceQueueItemsResponse)
	err := grpc.Invoke(ctx, "/api.DeviceQueue/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceQueue service

type DeviceQueueServer interface {
	// Enqueue adds the given item to the queue. When the node operates in
	// Class-C mode, the data will be pushed directly to the network-server.
	Enqueue(context.Context, *EnqueueDeviceQueueItemRequest) (*EnqueueDeviceQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(context.Context, *DeleteDeviceQueueItemRequest) (*DeleteDeviceQueueItemResponse, error)
	// List lists the items in the queue for the given node.
	List(context.Context, *ListDeviceQueueItemsRequest) (*ListDeviceQueueItemsResponse, error)
}

func RegisterDeviceQueueServer(s *grpc.Server, srv DeviceQueueServer) {
	s.RegisterService(&_DeviceQueue_serviceDesc, srv)
}

func _DeviceQueue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Enqueue(ctx, req.(*EnqueueDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).Delete(ctx, req.(*DeleteDeviceQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceQueue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceQueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceQueue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceQueueServer).List(ctx, req.(*ListDeviceQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceQueue",
	HandlerType: (*DeviceQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DeviceQueue_Enqueue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceQueue_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceQueue_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceQueue.proto",
}

func init() { proto.RegisterFile("deviceQueue.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0x66, 0x92, 0x6c, 0xb6, 0x3d, 0xf5, 0x07, 0x87, 0x22, 0x21, 0xcd, 0xea, 0x74, 0x0a, 0x12,
	0xf6, 0x62, 0x03, 0x15, 0x6f, 0xbc, 0x6e, 0x85, 0x8a, 0x17, 0x1a, 0xf0, 0x01, 0xe2, 0xce, 0xd9,
	0x65, 0xa0, 0x9d, 0x49, 0x93, 0xd9, 0x82, 0x96, 0xde, 0xf8, 0x0a, 0x5e, 0xfb, 0x20, 0x3e, 0x87,
	0xaf, 0xe0, 0x33, 0x78, 0x2d, 0x99, 0x44, 0xd2, 0x86, 0xcd, 0xec, 0xdd, 0x9c, 0xdf, 0xef, 0x9c,
	0xef, 0x3b, 0x03, 0xcf, 0x04, 0xde, 0xc8, 0x25, 0x7e, 0xda, 0xe0, 0x06, 0x17, 0x65, 0xa5, 0x8d,
	0xa6, 0x7e, 0x51, 0xca, 0x38, 0x59, 0x6b, 0xbd, 0xbe, 0xc4, 0xac, 0x28, 0x65, 0x56, 0x28, 0xa5,
	0x4d, 0x61, 0xa4, 0x56, 0x75, 0x9b, 0xc2, 0x7f, 0x12, 0x98, 0x9d, 0xab, 0xeb, 0xa6, 0xe8, 0xac,
	0xaf, 0xbf, 0x30, 0x78, 0x95, 0xe3, 0xf5, 0x06, 0x6b, 0x43, 0x9f, 0x43, 0x28, 0xf0, 0xe6, 0xfc,
	0xf3, 0x45, 0x44, 0x18, 0x49, 0xf7, 0xf3, 0xce, 0xa2, 0x09, 0xec, 0x57, 0xb8, 0xc2, 0x0a, 0xd5,
	0x12, 0x23, 0xcf, 0x86, 0x7a, 0x47, 0x13, 0x5d, 0x6a, 0xb5, 0x92, 0xd5, 0x15, 0x8a, 0xc8, 0x67,
	0x24, 0xdd, 0xcb, 0x7b, 0x07, 0x3d, 0x84, 0xc9, 0xea, 0xa3, 0xae, 0x4c, 0x14, 0x30, 0x92, 0x3e,
	0xce, 0x5b, 0x83, 0x52, 0x08, 0x44, 0x61, 0x8a, 0x68, 0xc2, 0x48, 0xfa, 0x28, 0xb7, 0x6f, 0xce,
	0xe0, 0xc5, 0xd8, 0x78, 0x75, 0xa9, 0x55, 0x8d, 0xfc, 0x1d, 0x24, 0x67, 0x78, 0x89, 0x66, 0xf7,
	0xfc, 0xde, 0x83, 0xf9, 0x9f, 0x80, 0x27, 0x85, 0xdd, 0xc9, 0xcf, 0x3d, 0x29, 0xf8, 0x4b, 0x98,
	0x8d, 0xf4, 0xe9, 0x80, 0x7e, 0x11, 0x78, 0x3a, 0x88, 0x0d, 0x9b, 0x8c, 0x82, 0x3d, 0x20, 0xcb,
	0x77, 0x92, 0x15, 0x0c, 0xc9, 0x8a, 0x60, 0x5a, 0xa2, 0x12, 0x52, 0xad, 0x2d, 0x33, 0x7b, 0xf9,
	0x7f, 0xb3, 0xa7, 0x31, 0xdc, 0x46, 0xe3, 0xf4, 0x1e, 0x8d, 0x6f, 0xe0, 0xe8, 0x83, 0xac, 0xcd,
	0x60, 0xfc, 0x7a, 0x87, 0xc6, 0xfc, 0x3d, 0x24, 0xdb, 0xcb, 0x5a, 0x4a, 0xe8, 0x1c, 0x26, 0xb2,
	0x71, 0x44, 0x84, 0xf9, 0xe9, 0xc1, 0xe9, 0xe1, 0xa2, 0x28, 0xe5, 0x62, 0xc8, 0x5f, 0x9b, 0x72,
	0xfa, 0xd7, 0x83, 0x83, 0x7b, 0x21, 0xfa, 0x0d, 0xa6, 0x9d, 0xb2, 0x94, 0xdb, 0x3a, 0xe7, 0x19,
	0xc6, 0x27, 0xce, 0x9c, 0x4e, 0xa2, 0x57, 0xdf, 0x7f, 0xff, 0xf9, 0xe1, 0x31, 0x7e, 0x64, 0xaf,
	0xbd, 0xfd, 0x10, 0x75, 0x76, 0xdb, 0x6e, 0x73, 0x97, 0xd9, 0xda, 0xb7, 0x64, 0x4e, 0xbf, 0x42,
	0xd8, 0x6a, 0x4d, 0x8f, 0xbb, 0x91, 0xc7, 0x0f, 0x28, 0xe6, 0xae, 0x94, 0x0e, 0x38, 0xb5, 0xc0,
	0x7c, 0xce, 0x1c, 0xc0, 0xd9, 0xad, 0x14, 0x77, 0xb4, 0x84, 0xa0, 0xa1, 0x94, 0x32, 0xdb, 0xd5,
	0x21, 0x4a, 0x7c, 0xec, 0xc8, 0xe8, 0x60, 0x4f, 0x2c, 0xec, 0x8c, 0xba, 0xf6, 0xfd, 0x12, 0xda,
	0x9f, 0xfe, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x84, 0x7b, 0xe0, 0x21, 0x04, 0x00,
	0x00,
}
