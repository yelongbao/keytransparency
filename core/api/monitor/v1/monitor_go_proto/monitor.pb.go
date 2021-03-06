// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor/v1/monitor.proto

// Monitor Service
//
// The Key Transparency monitor server service consists of APIs to fetch
// monitor results queried using the mutations API.

package monitor_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	trillian "github.com/google/trillian"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// GetStateRequest requests the verification state of a keytransparency domain
// for a particular point in time.
type GetStateRequest struct {
	// kt_url is the URL of the keytransparency server for which the monitoring
	// result will be returned.
	KtUrl string `protobuf:"bytes,2,opt,name=kt_url,json=ktUrl,proto3" json:"kt_url,omitempty"`
	// domain_id identifies the merkle tree being monitored.
	DomainId string `protobuf:"bytes,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// epoch specifies the revision for which the monitoring results will
	// be returned (epochs start at 0).
	Epoch                int64    `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateRequest) Reset()         { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()    {}
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9cdd4901f6b9a2, []int{0}
}

func (m *GetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRequest.Unmarshal(m, b)
}
func (m *GetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRequest.Marshal(b, m, deterministic)
}
func (m *GetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRequest.Merge(m, src)
}
func (m *GetStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetStateRequest.Size(m)
}
func (m *GetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRequest proto.InternalMessageInfo

func (m *GetStateRequest) GetKtUrl() string {
	if m != nil {
		return m.KtUrl
	}
	return ""
}

func (m *GetStateRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetStateRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// State represents the monitor's evaluation of a Key Transparency domain
// at a particular epoch.
type State struct {
	// smr contains the map root for the sparse Merkle Tree signed with the
	// monitor's key on success. If the checks were not successful the
	// smr will be empty. The epochs are encoded into the smr map_revision.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,1,opt,name=smr,proto3" json:"smr,omitempty"`
	// seen_time contains the time when this particular signed map root was
	// retrieved and processed.
	SeenTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=seen_time,json=seenTime,proto3" json:"seen_time,omitempty"`
	// errors contains a list of errors representing the verification checks
	// that failed while monitoring the key-transparency server.
	Errors               []*status.Status `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9cdd4901f6b9a2, []int{1}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *State) GetSeenTime() *timestamp.Timestamp {
	if m != nil {
		return m.SeenTime
	}
	return nil
}

func (m *State) GetErrors() []*status.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStateRequest)(nil), "google.keytransparency.monitor.v1.GetStateRequest")
	proto.RegisterType((*State)(nil), "google.keytransparency.monitor.v1.State")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error)
}

type monitorClient struct {
	cc *grpc.ClientConn
}

func NewMonitorClient(cc *grpc.ClientConn) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorClient) GetStateByRevision(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServer is the server API for Monitor service.
type MonitorServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains extra data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetState(context.Context, *GetStateRequest) (*State, error)
	// GetSignedMapRootByRevision returns the monitor's result for a specific map
	// revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetStateByRevision(context.Context, *GetStateRequest) (*State, error)
}

func RegisterMonitorServer(s *grpc.Server, srv MonitorServer) {
	s.RegisterService(&_Monitor_serviceDesc, srv)
}

func _Monitor_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitor_GetStateByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServer).GetStateByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.monitor.v1.Monitor/GetStateByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServer).GetStateByRevision(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Monitor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.monitor.v1.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetState",
			Handler:    _Monitor_GetState_Handler,
		},
		{
			MethodName: "GetStateByRevision",
			Handler:    _Monitor_GetStateByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor/v1/monitor.proto",
}

func init() { proto.RegisterFile("monitor/v1/monitor.proto", fileDescriptor_6c9cdd4901f6b9a2) }

var fileDescriptor_6c9cdd4901f6b9a2 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0x96, 0x63, 0x25, 0x24, 0x1b, 0x09, 0xa4, 0x15, 0xe8, 0xac, 0x80, 0x44, 0x48, 0x15, 0x28,
	0x76, 0x75, 0xa6, 0x40, 0xa2, 0xe0, 0xe7, 0x8a, 0x3b, 0x28, 0xae, 0x71, 0xa0, 0x81, 0xc2, 0xda,
	0x73, 0x16, 0xdf, 0x2a, 0xf6, 0x8e, 0xd9, 0x1d, 0x5b, 0x8a, 0xa2, 0x34, 0xbc, 0x02, 0x05, 0xaf,
	0x41, 0xcf, 0x63, 0xf0, 0x0a, 0x3c, 0x05, 0x15, 0xf2, 0xae, 0x7d, 0x42, 0xa1, 0x00, 0x21, 0x5d,
	0x63, 0x7b, 0x3c, 0xdf, 0xb7, 0xfe, 0xe6, 0x9b, 0xcf, 0x24, 0x2a, 0x41, 0x2b, 0x04, 0xc3, 0x9b,
	0x63, 0xde, 0x3d, 0xb2, 0xca, 0x00, 0x02, 0x7d, 0x90, 0x03, 0xe4, 0x85, 0x64, 0x1b, 0xb9, 0x45,
	0x23, 0xb4, 0xad, 0x84, 0x91, 0x3a, 0xdb, 0xb2, 0x1e, 0xd5, 0x1c, 0xcf, 0xee, 0x79, 0x08, 0x17,
	0x95, 0xe2, 0x42, 0x6b, 0x40, 0x81, 0x0a, 0xb4, 0xf5, 0x07, 0xcc, 0xee, 0x77, 0x5d, 0x57, 0x5d,
	0xd4, 0x1f, 0x38, 0xaa, 0x52, 0x5a, 0x14, 0x65, 0xd5, 0x01, 0x8e, 0x3a, 0x80, 0xa9, 0x32, 0x6e,
	0x51, 0x60, 0xdd, 0x33, 0x6f, 0xa2, 0x51, 0x45, 0xa1, 0x84, 0xf6, 0xf5, 0xe2, 0x3d, 0xb9, 0x75,
	0x26, 0x71, 0x85, 0x02, 0x65, 0x22, 0x3f, 0xd6, 0xd2, 0x22, 0xbd, 0x43, 0x46, 0x1b, 0x4c, 0x6b,
	0x53, 0x44, 0x83, 0x79, 0xb0, 0x9c, 0x24, 0xc3, 0x0d, 0xbe, 0x35, 0x05, 0xbd, 0x4b, 0x26, 0x6b,
	0x28, 0x85, 0xd2, 0xa9, 0x5a, 0x47, 0xa1, 0xeb, 0x8c, 0xfd, 0x8b, 0xd7, 0x6b, 0x7a, 0x9b, 0x0c,
	0x65, 0x05, 0xd9, 0x65, 0x14, 0xcc, 0x83, 0x65, 0x98, 0xf8, 0x62, 0xf1, 0x25, 0x20, 0x43, 0x77,
	0x34, 0x7d, 0x48, 0x42, 0x5b, 0x1a, 0xd7, 0x9d, 0xc6, 0x47, 0xec, 0x4a, 0xc4, 0x4a, 0xe5, 0x5a,
	0xae, 0xcf, 0x45, 0x95, 0x00, 0x60, 0xd2, 0x62, 0xe8, 0x13, 0x32, 0xb1, 0x52, 0xea, 0xb4, 0x1d,
	0xc9, 0x29, 0x98, 0xc6, 0x33, 0xd6, 0x19, 0xd6, 0xcf, 0xcb, 0xde, 0xf4, 0xf3, 0x26, 0xe3, 0x16,
	0xdc, 0x96, 0xf4, 0x11, 0x19, 0x49, 0x63, 0xc0, 0xd8, 0x28, 0x9c, 0x87, 0xcb, 0x69, 0x4c, 0x7b,
	0x96, 0xa9, 0x32, 0xb6, 0x72, 0x26, 0x24, 0x1d, 0x22, 0xfe, 0x39, 0x20, 0x37, 0xce, 0xbd, 0xdb,
	0xf4, 0x6b, 0x40, 0xc6, 0xbd, 0x07, 0x34, 0x66, 0x7f, 0xdd, 0x0d, 0x3b, 0x30, 0x6c, 0xb6, 0xfc,
	0x07, 0x8e, 0x23, 0x2c, 0x4e, 0x3f, 0x7d, 0xff, 0xf1, 0x79, 0xf0, 0x82, 0x3e, 0xe3, 0xbf, 0x65,
	0xc3, 0x4a, 0xd3, 0x48, 0x63, 0xf9, 0xce, 0xbb, 0xbe, 0xe7, 0xde, 0x55, 0xcb, 0x77, 0x57, 0x7e,
	0xef, 0xdd, 0x12, 0xa5, 0x7d, 0x5a, 0xb4, 0x57, 0xa4, 0xdf, 0x02, 0x42, 0x7b, 0x15, 0x27, 0xdb,
	0x44, 0x36, 0xca, 0x2a, 0xd0, 0xd7, 0x2c, 0xfe, 0xcc, 0x89, 0x7f, 0x49, 0x9f, 0xff, 0xa7, 0x78,
	0xbe, 0x73, 0xa9, 0xd8, 0x9f, 0xbc, 0x7a, 0x77, 0x9a, 0x2b, 0xbc, 0xac, 0x2f, 0x58, 0x06, 0x25,
	0xef, 0x92, 0x7a, 0xf0, 0x79, 0x9e, 0x81, 0xf1, 0xe9, 0xff, 0xf3, 0x2f, 0x4a, 0x73, 0x48, 0x7d,
	0x12, 0x46, 0xee, 0xf6, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0xf0, 0x6d, 0xdf, 0x6b,
	0x03, 0x00, 0x00,
}
