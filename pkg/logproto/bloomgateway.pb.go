// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/logproto/bloomgateway.proto

package logproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_grafana_loki_pkg_logql_syntax "github.com/grafana/loki/pkg/logql/syntax"
	github_com_prometheus_common_model "github.com/prometheus/common/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type FilterChunkRefRequest struct {
	From    github_com_prometheus_common_model.Time               `protobuf:"varint,1,opt,name=from,proto3,customtype=github.com/prometheus/common/model.Time" json:"from"`
	Through github_com_prometheus_common_model.Time               `protobuf:"varint,2,opt,name=through,proto3,customtype=github.com/prometheus/common/model.Time" json:"through"`
	Refs    []*GroupedChunkRefs                                   `protobuf:"bytes,3,rep,name=refs,proto3" json:"refs,omitempty"`
	Filters []github_com_grafana_loki_pkg_logql_syntax.LineFilter `protobuf:"bytes,4,rep,name=filters,proto3,customtype=github.com/grafana/loki/pkg/logql/syntax.LineFilter" json:"filters"`
}

func (m *FilterChunkRefRequest) Reset()      { *m = FilterChunkRefRequest{} }
func (*FilterChunkRefRequest) ProtoMessage() {}
func (*FilterChunkRefRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50b5dd1dbcd1415, []int{0}
}
func (m *FilterChunkRefRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterChunkRefRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterChunkRefRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterChunkRefRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChunkRefRequest.Merge(m, src)
}
func (m *FilterChunkRefRequest) XXX_Size() int {
	return m.Size()
}
func (m *FilterChunkRefRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChunkRefRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChunkRefRequest proto.InternalMessageInfo

func (m *FilterChunkRefRequest) GetRefs() []*GroupedChunkRefs {
	if m != nil {
		return m.Refs
	}
	return nil
}

type FilterChunkRefResponse struct {
	ChunkRefs []*GroupedChunkRefs `protobuf:"bytes,1,rep,name=chunkRefs,proto3" json:"chunkRefs,omitempty"`
}

func (m *FilterChunkRefResponse) Reset()      { *m = FilterChunkRefResponse{} }
func (*FilterChunkRefResponse) ProtoMessage() {}
func (*FilterChunkRefResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50b5dd1dbcd1415, []int{1}
}
func (m *FilterChunkRefResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterChunkRefResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterChunkRefResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterChunkRefResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChunkRefResponse.Merge(m, src)
}
func (m *FilterChunkRefResponse) XXX_Size() int {
	return m.Size()
}
func (m *FilterChunkRefResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChunkRefResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChunkRefResponse proto.InternalMessageInfo

func (m *FilterChunkRefResponse) GetChunkRefs() []*GroupedChunkRefs {
	if m != nil {
		return m.ChunkRefs
	}
	return nil
}

type ShortRef struct {
	From     github_com_prometheus_common_model.Time `protobuf:"varint,1,opt,name=from,proto3,customtype=github.com/prometheus/common/model.Time" json:"from"`
	Through  github_com_prometheus_common_model.Time `protobuf:"varint,2,opt,name=through,proto3,customtype=github.com/prometheus/common/model.Time" json:"through"`
	Checksum uint32                                  `protobuf:"varint,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *ShortRef) Reset()      { *m = ShortRef{} }
func (*ShortRef) ProtoMessage() {}
func (*ShortRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50b5dd1dbcd1415, []int{2}
}
func (m *ShortRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShortRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShortRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShortRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShortRef.Merge(m, src)
}
func (m *ShortRef) XXX_Size() int {
	return m.Size()
}
func (m *ShortRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ShortRef.DiscardUnknown(m)
}

var xxx_messageInfo_ShortRef proto.InternalMessageInfo

func (m *ShortRef) GetChecksum() uint32 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

type GroupedChunkRefs struct {
	Fingerprint uint64      `protobuf:"varint,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Tenant      string      `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Refs        []*ShortRef `protobuf:"bytes,3,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (m *GroupedChunkRefs) Reset()      { *m = GroupedChunkRefs{} }
func (*GroupedChunkRefs) ProtoMessage() {}
func (*GroupedChunkRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a50b5dd1dbcd1415, []int{3}
}
func (m *GroupedChunkRefs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupedChunkRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupedChunkRefs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupedChunkRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupedChunkRefs.Merge(m, src)
}
func (m *GroupedChunkRefs) XXX_Size() int {
	return m.Size()
}
func (m *GroupedChunkRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupedChunkRefs.DiscardUnknown(m)
}

var xxx_messageInfo_GroupedChunkRefs proto.InternalMessageInfo

func (m *GroupedChunkRefs) GetFingerprint() uint64 {
	if m != nil {
		return m.Fingerprint
	}
	return 0
}

func (m *GroupedChunkRefs) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *GroupedChunkRefs) GetRefs() []*ShortRef {
	if m != nil {
		return m.Refs
	}
	return nil
}

func init() {
	proto.RegisterType((*FilterChunkRefRequest)(nil), "logproto.FilterChunkRefRequest")
	proto.RegisterType((*FilterChunkRefResponse)(nil), "logproto.FilterChunkRefResponse")
	proto.RegisterType((*ShortRef)(nil), "logproto.ShortRef")
	proto.RegisterType((*GroupedChunkRefs)(nil), "logproto.GroupedChunkRefs")
}

func init() { proto.RegisterFile("pkg/logproto/bloomgateway.proto", fileDescriptor_a50b5dd1dbcd1415) }

var fileDescriptor_a50b5dd1dbcd1415 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xbd, 0x6e, 0xd4, 0x30,
	0x1c, 0x8f, 0x7b, 0xa7, 0xf6, 0xea, 0x82, 0x40, 0x56, 0xa9, 0xa2, 0x20, 0xf9, 0xa2, 0x08, 0xc1,
	0x4d, 0x89, 0xd4, 0x2e, 0x48, 0x6c, 0x57, 0x89, 0x0a, 0x89, 0xc9, 0x20, 0x86, 0x6e, 0xb9, 0xd4,
	0xf9, 0x50, 0x12, 0xff, 0x53, 0xdb, 0x11, 0x74, 0xe3, 0x11, 0x78, 0x0c, 0x9e, 0x80, 0x27, 0x60,
	0xe8, 0x78, 0x63, 0xc5, 0x50, 0x71, 0xb9, 0x85, 0xb1, 0x8f, 0x80, 0xea, 0x5c, 0x7a, 0x77, 0x15,
	0xe8, 0x24, 0x26, 0x26, 0x7f, 0xfc, 0xff, 0x3f, 0xfb, 0xf7, 0x61, 0xe3, 0x61, 0x95, 0x27, 0x41,
	0x01, 0x49, 0x25, 0x41, 0x43, 0x30, 0x29, 0x00, 0xca, 0x24, 0xd4, 0xfc, 0x63, 0x78, 0xe1, 0x9b,
	0x2d, 0x32, 0xe8, 0x8a, 0xce, 0x7e, 0x02, 0x09, 0xb4, 0x7d, 0xb7, 0xb3, 0xb6, 0xee, 0x3c, 0x5d,
	0x3b, 0xa0, 0x9b, 0xb4, 0x45, 0xef, 0xfb, 0x16, 0x7e, 0xf2, 0x3a, 0x2b, 0x34, 0x97, 0xc7, 0x69,
	0x2d, 0x72, 0xc6, 0x63, 0xc6, 0xcf, 0x6b, 0xae, 0x34, 0x39, 0xc6, 0xfd, 0x58, 0x42, 0x69, 0x23,
	0x17, 0x8d, 0x7a, 0xe3, 0xe0, 0xf2, 0x7a, 0x68, 0xfd, 0xb8, 0x1e, 0xbe, 0x48, 0x32, 0x9d, 0xd6,
	0x13, 0x3f, 0x82, 0x32, 0xa8, 0x24, 0x94, 0x5c, 0xa7, 0xbc, 0x56, 0x41, 0x04, 0x65, 0x09, 0x22,
	0x28, 0xe1, 0x8c, 0x17, 0xfe, 0xfb, 0xac, 0xe4, 0xcc, 0x80, 0xc9, 0x1b, 0xbc, 0xa3, 0x53, 0x09,
	0x75, 0x92, 0xda, 0x5b, 0xff, 0x76, 0x4e, 0x87, 0x27, 0x3e, 0xee, 0x4b, 0x1e, 0x2b, 0xbb, 0xe7,
	0xf6, 0x46, 0x7b, 0x87, 0x8e, 0x7f, 0x27, 0xe4, 0x44, 0x42, 0x5d, 0xf1, 0xb3, 0x8e, 0xbf, 0x62,
	0xa6, 0x8f, 0xe4, 0x78, 0x27, 0x36, 0xc2, 0x94, 0xdd, 0x37, 0x90, 0xfd, 0x25, 0xe4, 0x6d, 0x26,
	0x78, 0xab, 0x7a, 0xfc, 0x6a, 0x41, 0xe8, 0x68, 0x85, 0x50, 0x22, 0xc3, 0x38, 0x14, 0x61, 0x50,
	0x40, 0x9e, 0x05, 0x0b, 0xf7, 0xce, 0x8b, 0x40, 0x5d, 0x08, 0x1d, 0x7e, 0x5a, 0x01, 0xb3, 0xee,
	0x06, 0x8f, 0xe1, 0x83, 0xfb, 0x2e, 0xaa, 0x0a, 0x84, 0xe2, 0xe4, 0x25, 0xde, 0x8d, 0x3a, 0x66,
	0x36, 0xda, 0xc8, 0x7d, 0xd9, 0xec, 0x7d, 0x43, 0x78, 0xf0, 0x2e, 0x05, 0xa9, 0x19, 0x8f, 0xff,
	0xbb, 0x34, 0x1c, 0x3c, 0x88, 0x52, 0x1e, 0xe5, 0xaa, 0x2e, 0xed, 0x9e, 0x8b, 0x46, 0x0f, 0xd9,
	0xdd, 0xda, 0xd3, 0xf8, 0xf1, 0x7d, 0x5d, 0xc4, 0xc5, 0x7b, 0x71, 0x26, 0x12, 0x2e, 0x2b, 0x99,
	0x09, 0x6d, 0x64, 0xf4, 0xd9, 0xea, 0x16, 0x39, 0xc0, 0xdb, 0x9a, 0x8b, 0x50, 0x68, 0xc3, 0x6d,
	0x97, 0x2d, 0x56, 0xe4, 0xf9, 0x5a, 0xee, 0x64, 0xe9, 0x5d, 0xe7, 0x4d, 0x9b, 0xf7, 0x61, 0x8c,
	0x1f, 0x8c, 0x6f, 0x3f, 0xc7, 0x49, 0xfb, 0x39, 0xc8, 0x07, 0xfc, 0x68, 0x3d, 0x12, 0x45, 0x86,
	0x4b, 0xf0, 0x1f, 0xdf, 0xbc, 0xe3, 0xfe, 0xbd, 0xa1, 0x8d, 0xd3, 0xb3, 0xc6, 0xa7, 0xd3, 0x19,
	0xb5, 0xae, 0x66, 0xd4, 0xba, 0x99, 0x51, 0xf4, 0xb9, 0xa1, 0xe8, 0x6b, 0x43, 0xd1, 0x65, 0x43,
	0xd1, 0xb4, 0xa1, 0xe8, 0x67, 0x43, 0xd1, 0xaf, 0x86, 0x5a, 0x37, 0x0d, 0x45, 0x5f, 0xe6, 0xd4,
	0x9a, 0xce, 0xa9, 0x75, 0x35, 0xa7, 0xd6, 0xe9, 0xb3, 0x0d, 0xcf, 0xcb, 0x5c, 0x3a, 0xd9, 0x36,
	0xc3, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x30, 0x9d, 0x8e, 0xf4, 0x03, 0x00, 0x00,
}

func (this *FilterChunkRefRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterChunkRefRequest)
	if !ok {
		that2, ok := that.(FilterChunkRefRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.From.Equal(that1.From) {
		return false
	}
	if !this.Through.Equal(that1.Through) {
		return false
	}
	if len(this.Refs) != len(that1.Refs) {
		return false
	}
	for i := range this.Refs {
		if !this.Refs[i].Equal(that1.Refs[i]) {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	return true
}
func (this *FilterChunkRefResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterChunkRefResponse)
	if !ok {
		that2, ok := that.(FilterChunkRefResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ChunkRefs) != len(that1.ChunkRefs) {
		return false
	}
	for i := range this.ChunkRefs {
		if !this.ChunkRefs[i].Equal(that1.ChunkRefs[i]) {
			return false
		}
	}
	return true
}
func (this *ShortRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ShortRef)
	if !ok {
		that2, ok := that.(ShortRef)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.From.Equal(that1.From) {
		return false
	}
	if !this.Through.Equal(that1.Through) {
		return false
	}
	if this.Checksum != that1.Checksum {
		return false
	}
	return true
}
func (this *GroupedChunkRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GroupedChunkRefs)
	if !ok {
		that2, ok := that.(GroupedChunkRefs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Fingerprint != that1.Fingerprint {
		return false
	}
	if this.Tenant != that1.Tenant {
		return false
	}
	if len(this.Refs) != len(that1.Refs) {
		return false
	}
	for i := range this.Refs {
		if !this.Refs[i].Equal(that1.Refs[i]) {
			return false
		}
	}
	return true
}
func (this *FilterChunkRefRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&logproto.FilterChunkRefRequest{")
	s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",\n")
	s = append(s, "Through: "+fmt.Sprintf("%#v", this.Through)+",\n")
	if this.Refs != nil {
		s = append(s, "Refs: "+fmt.Sprintf("%#v", this.Refs)+",\n")
	}
	s = append(s, "Filters: "+fmt.Sprintf("%#v", this.Filters)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FilterChunkRefResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&logproto.FilterChunkRefResponse{")
	if this.ChunkRefs != nil {
		s = append(s, "ChunkRefs: "+fmt.Sprintf("%#v", this.ChunkRefs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ShortRef) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&logproto.ShortRef{")
	s = append(s, "From: "+fmt.Sprintf("%#v", this.From)+",\n")
	s = append(s, "Through: "+fmt.Sprintf("%#v", this.Through)+",\n")
	s = append(s, "Checksum: "+fmt.Sprintf("%#v", this.Checksum)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GroupedChunkRefs) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&logproto.GroupedChunkRefs{")
	s = append(s, "Fingerprint: "+fmt.Sprintf("%#v", this.Fingerprint)+",\n")
	s = append(s, "Tenant: "+fmt.Sprintf("%#v", this.Tenant)+",\n")
	if this.Refs != nil {
		s = append(s, "Refs: "+fmt.Sprintf("%#v", this.Refs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBloomgateway(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BloomGatewayClient is the client API for BloomGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BloomGatewayClient interface {
	FilterChunkRefs(ctx context.Context, in *FilterChunkRefRequest, opts ...grpc.CallOption) (*FilterChunkRefResponse, error)
}

type bloomGatewayClient struct {
	cc *grpc.ClientConn
}

func NewBloomGatewayClient(cc *grpc.ClientConn) BloomGatewayClient {
	return &bloomGatewayClient{cc}
}

func (c *bloomGatewayClient) FilterChunkRefs(ctx context.Context, in *FilterChunkRefRequest, opts ...grpc.CallOption) (*FilterChunkRefResponse, error) {
	out := new(FilterChunkRefResponse)
	err := c.cc.Invoke(ctx, "/logproto.BloomGateway/FilterChunkRefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BloomGatewayServer is the server API for BloomGateway service.
type BloomGatewayServer interface {
	FilterChunkRefs(context.Context, *FilterChunkRefRequest) (*FilterChunkRefResponse, error)
}

// UnimplementedBloomGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedBloomGatewayServer struct {
}

func (*UnimplementedBloomGatewayServer) FilterChunkRefs(ctx context.Context, req *FilterChunkRefRequest) (*FilterChunkRefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterChunkRefs not implemented")
}

func RegisterBloomGatewayServer(s *grpc.Server, srv BloomGatewayServer) {
	s.RegisterService(&_BloomGateway_serviceDesc, srv)
}

func _BloomGateway_FilterChunkRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterChunkRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloomGatewayServer).FilterChunkRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logproto.BloomGateway/FilterChunkRefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloomGatewayServer).FilterChunkRefs(ctx, req.(*FilterChunkRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BloomGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logproto.BloomGateway",
	HandlerType: (*BloomGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilterChunkRefs",
			Handler:    _BloomGateway_FilterChunkRefs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/logproto/bloomgateway.proto",
}

func (m *FilterChunkRefRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterChunkRefRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterChunkRefRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Filters) > 0 {
		for iNdEx := len(m.Filters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Filters[iNdEx].Size()
				i -= size
				if _, err := m.Filters[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintBloomgateway(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Refs) > 0 {
		for iNdEx := len(m.Refs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Refs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBloomgateway(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Through != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.Through))
		i--
		dAtA[i] = 0x10
	}
	if m.From != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.From))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FilterChunkRefResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterChunkRefResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterChunkRefResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChunkRefs) > 0 {
		for iNdEx := len(m.ChunkRefs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChunkRefs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBloomgateway(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ShortRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShortRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShortRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checksum != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.Checksum))
		i--
		dAtA[i] = 0x18
	}
	if m.Through != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.Through))
		i--
		dAtA[i] = 0x10
	}
	if m.From != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.From))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GroupedChunkRefs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupedChunkRefs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupedChunkRefs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Refs) > 0 {
		for iNdEx := len(m.Refs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Refs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBloomgateway(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintBloomgateway(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0x12
	}
	if m.Fingerprint != 0 {
		i = encodeVarintBloomgateway(dAtA, i, uint64(m.Fingerprint))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBloomgateway(dAtA []byte, offset int, v uint64) int {
	offset -= sovBloomgateway(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilterChunkRefRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != 0 {
		n += 1 + sovBloomgateway(uint64(m.From))
	}
	if m.Through != 0 {
		n += 1 + sovBloomgateway(uint64(m.Through))
	}
	if len(m.Refs) > 0 {
		for _, e := range m.Refs {
			l = e.Size()
			n += 1 + l + sovBloomgateway(uint64(l))
		}
	}
	if len(m.Filters) > 0 {
		for _, e := range m.Filters {
			l = e.Size()
			n += 1 + l + sovBloomgateway(uint64(l))
		}
	}
	return n
}

func (m *FilterChunkRefResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChunkRefs) > 0 {
		for _, e := range m.ChunkRefs {
			l = e.Size()
			n += 1 + l + sovBloomgateway(uint64(l))
		}
	}
	return n
}

func (m *ShortRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != 0 {
		n += 1 + sovBloomgateway(uint64(m.From))
	}
	if m.Through != 0 {
		n += 1 + sovBloomgateway(uint64(m.Through))
	}
	if m.Checksum != 0 {
		n += 1 + sovBloomgateway(uint64(m.Checksum))
	}
	return n
}

func (m *GroupedChunkRefs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fingerprint != 0 {
		n += 1 + sovBloomgateway(uint64(m.Fingerprint))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovBloomgateway(uint64(l))
	}
	if len(m.Refs) > 0 {
		for _, e := range m.Refs {
			l = e.Size()
			n += 1 + l + sovBloomgateway(uint64(l))
		}
	}
	return n
}

func sovBloomgateway(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBloomgateway(x uint64) (n int) {
	return sovBloomgateway(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FilterChunkRefRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForRefs := "[]*GroupedChunkRefs{"
	for _, f := range this.Refs {
		repeatedStringForRefs += strings.Replace(f.String(), "GroupedChunkRefs", "GroupedChunkRefs", 1) + ","
	}
	repeatedStringForRefs += "}"
	s := strings.Join([]string{`&FilterChunkRefRequest{`,
		`From:` + fmt.Sprintf("%v", this.From) + `,`,
		`Through:` + fmt.Sprintf("%v", this.Through) + `,`,
		`Refs:` + repeatedStringForRefs + `,`,
		`Filters:` + fmt.Sprintf("%v", this.Filters) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FilterChunkRefResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForChunkRefs := "[]*GroupedChunkRefs{"
	for _, f := range this.ChunkRefs {
		repeatedStringForChunkRefs += strings.Replace(f.String(), "GroupedChunkRefs", "GroupedChunkRefs", 1) + ","
	}
	repeatedStringForChunkRefs += "}"
	s := strings.Join([]string{`&FilterChunkRefResponse{`,
		`ChunkRefs:` + repeatedStringForChunkRefs + `,`,
		`}`,
	}, "")
	return s
}
func (this *ShortRef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ShortRef{`,
		`From:` + fmt.Sprintf("%v", this.From) + `,`,
		`Through:` + fmt.Sprintf("%v", this.Through) + `,`,
		`Checksum:` + fmt.Sprintf("%v", this.Checksum) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GroupedChunkRefs) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForRefs := "[]*ShortRef{"
	for _, f := range this.Refs {
		repeatedStringForRefs += strings.Replace(f.String(), "ShortRef", "ShortRef", 1) + ","
	}
	repeatedStringForRefs += "}"
	s := strings.Join([]string{`&GroupedChunkRefs{`,
		`Fingerprint:` + fmt.Sprintf("%v", this.Fingerprint) + `,`,
		`Tenant:` + fmt.Sprintf("%v", this.Tenant) + `,`,
		`Refs:` + repeatedStringForRefs + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBloomgateway(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FilterChunkRefRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBloomgateway
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
			return fmt.Errorf("proto: FilterChunkRefRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterChunkRefRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= github_com_prometheus_common_model.Time(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Through", wireType)
			}
			m.Through = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Through |= github_com_prometheus_common_model.Time(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
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
				return ErrInvalidLengthBloomgateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Refs = append(m.Refs, &GroupedChunkRefs{})
			if err := m.Refs[len(m.Refs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
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
				return ErrInvalidLengthBloomgateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, github_com_grafana_loki_pkg_logql_syntax.LineFilter{})
			if err := m.Filters[len(m.Filters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBloomgateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBloomgateway
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
func (m *FilterChunkRefResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBloomgateway
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
			return fmt.Errorf("proto: FilterChunkRefResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterChunkRefResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
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
				return ErrInvalidLengthBloomgateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChunkRefs = append(m.ChunkRefs, &GroupedChunkRefs{})
			if err := m.ChunkRefs[len(m.ChunkRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBloomgateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBloomgateway
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
func (m *ShortRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBloomgateway
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
			return fmt.Errorf("proto: ShortRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShortRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= github_com_prometheus_common_model.Time(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Through", wireType)
			}
			m.Through = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Through |= github_com_prometheus_common_model.Time(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			m.Checksum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Checksum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBloomgateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBloomgateway
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
func (m *GroupedChunkRefs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBloomgateway
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
			return fmt.Errorf("proto: GroupedChunkRefs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupedChunkRefs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fingerprint", wireType)
			}
			m.Fingerprint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fingerprint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
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
				return ErrInvalidLengthBloomgateway
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBloomgateway
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
				return ErrInvalidLengthBloomgateway
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Refs = append(m.Refs, &ShortRef{})
			if err := m.Refs[len(m.Refs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBloomgateway(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBloomgateway
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBloomgateway
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
func skipBloomgateway(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBloomgateway
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
					return 0, ErrIntOverflowBloomgateway
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBloomgateway
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
				return 0, ErrInvalidLengthBloomgateway
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBloomgateway
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBloomgateway
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBloomgateway(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBloomgateway
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBloomgateway = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBloomgateway   = fmt.Errorf("proto: integer overflow")
)
