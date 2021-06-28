// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: resource_usage_agent.proto

package resource_usage_agent

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

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

type CPUTimeRecord struct {
	ResourceGroupTag []byte `protobuf:"bytes,1,opt,name=resource_group_tag,json=resourceGroupTag,proto3" json:"resource_group_tag,omitempty"`
	// UNIX timestamp in second.
	RecordListTimestampSec []uint64 `protobuf:"varint,2,rep,packed,name=record_list_timestamp_sec,json=recordListTimestampSec" json:"record_list_timestamp_sec,omitempty"`
	// The value can be greater than 1000ms if the requests are running parallelly.
	RecordListCpuTimeMs  []uint32 `protobuf:"varint,3,rep,packed,name=record_list_cpu_time_ms,json=recordListCpuTimeMs" json:"record_list_cpu_time_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPUTimeRecord) Reset()         { *m = CPUTimeRecord{} }
func (m *CPUTimeRecord) String() string { return proto.CompactTextString(m) }
func (*CPUTimeRecord) ProtoMessage()    {}
func (*CPUTimeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_623affeb03bbc261, []int{0}
}
func (m *CPUTimeRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CPUTimeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CPUTimeRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CPUTimeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUTimeRecord.Merge(dst, src)
}
func (m *CPUTimeRecord) XXX_Size() int {
	return m.Size()
}
func (m *CPUTimeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUTimeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CPUTimeRecord proto.InternalMessageInfo

func (m *CPUTimeRecord) GetResourceGroupTag() []byte {
	if m != nil {
		return m.ResourceGroupTag
	}
	return nil
}

func (m *CPUTimeRecord) GetRecordListTimestampSec() []uint64 {
	if m != nil {
		return m.RecordListTimestampSec
	}
	return nil
}

func (m *CPUTimeRecord) GetRecordListCpuTimeMs() []uint32 {
	if m != nil {
		return m.RecordListCpuTimeMs
	}
	return nil
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_usage_agent_623affeb03bbc261, []int{1}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(dst, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return m.Size()
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CPUTimeRecord)(nil), "resource_usage_agent.CPUTimeRecord")
	proto.RegisterType((*EmptyResponse)(nil), "resource_usage_agent.EmptyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceUsageAgent service

type ResourceUsageAgentClient interface {
	// Report the CPU time records. By default, the records with the same
	// resource group tag will be batched by minute.
	ReportCPUTime(ctx context.Context, opts ...grpc.CallOption) (ResourceUsageAgent_ReportCPUTimeClient, error)
}

type resourceUsageAgentClient struct {
	cc *grpc.ClientConn
}

func NewResourceUsageAgentClient(cc *grpc.ClientConn) ResourceUsageAgentClient {
	return &resourceUsageAgentClient{cc}
}

func (c *resourceUsageAgentClient) ReportCPUTime(ctx context.Context, opts ...grpc.CallOption) (ResourceUsageAgent_ReportCPUTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResourceUsageAgent_serviceDesc.Streams[0], "/resource_usage_agent.ResourceUsageAgent/ReportCPUTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceUsageAgentReportCPUTimeClient{stream}
	return x, nil
}

type ResourceUsageAgent_ReportCPUTimeClient interface {
	Send(*CPUTimeRecord) error
	CloseAndRecv() (*EmptyResponse, error)
	grpc.ClientStream
}

type resourceUsageAgentReportCPUTimeClient struct {
	grpc.ClientStream
}

func (x *resourceUsageAgentReportCPUTimeClient) Send(m *CPUTimeRecord) error {
	return x.ClientStream.SendMsg(m)
}

func (x *resourceUsageAgentReportCPUTimeClient) CloseAndRecv() (*EmptyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ResourceUsageAgent service

type ResourceUsageAgentServer interface {
	// Report the CPU time records. By default, the records with the same
	// resource group tag will be batched by minute.
	ReportCPUTime(ResourceUsageAgent_ReportCPUTimeServer) error
}

func RegisterResourceUsageAgentServer(s *grpc.Server, srv ResourceUsageAgentServer) {
	s.RegisterService(&_ResourceUsageAgent_serviceDesc, srv)
}

func _ResourceUsageAgent_ReportCPUTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResourceUsageAgentServer).ReportCPUTime(&resourceUsageAgentReportCPUTimeServer{stream})
}

type ResourceUsageAgent_ReportCPUTimeServer interface {
	SendAndClose(*EmptyResponse) error
	Recv() (*CPUTimeRecord, error)
	grpc.ServerStream
}

type resourceUsageAgentReportCPUTimeServer struct {
	grpc.ServerStream
}

func (x *resourceUsageAgentReportCPUTimeServer) SendAndClose(m *EmptyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *resourceUsageAgentReportCPUTimeServer) Recv() (*CPUTimeRecord, error) {
	m := new(CPUTimeRecord)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ResourceUsageAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource_usage_agent.ResourceUsageAgent",
	HandlerType: (*ResourceUsageAgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportCPUTime",
			Handler:       _ResourceUsageAgent_ReportCPUTime_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "resource_usage_agent.proto",
}

func (m *CPUTimeRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPUTimeRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ResourceGroupTag) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(len(m.ResourceGroupTag)))
		i += copy(dAtA[i:], m.ResourceGroupTag)
	}
	if len(m.RecordListTimestampSec) > 0 {
		dAtA2 := make([]byte, len(m.RecordListTimestampSec)*10)
		var j1 int
		for _, num := range m.RecordListTimestampSec {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.RecordListCpuTimeMs) > 0 {
		dAtA4 := make([]byte, len(m.RecordListCpuTimeMs)*10)
		var j3 int
		for _, num := range m.RecordListCpuTimeMs {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintResourceUsageAgent(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintResourceUsageAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CPUTimeRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.ResourceGroupTag)
	if l > 0 {
		n += 1 + l + sovResourceUsageAgent(uint64(l))
	}
	if len(m.RecordListTimestampSec) > 0 {
		l = 0
		for _, e := range m.RecordListTimestampSec {
			l += sovResourceUsageAgent(uint64(e))
		}
		n += 1 + sovResourceUsageAgent(uint64(l)) + l
	}
	if len(m.RecordListCpuTimeMs) > 0 {
		l = 0
		for _, e := range m.RecordListCpuTimeMs {
			l += sovResourceUsageAgent(uint64(e))
		}
		n += 1 + sovResourceUsageAgent(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EmptyResponse) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovResourceUsageAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozResourceUsageAgent(x uint64) (n int) {
	return sovResourceUsageAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CPUTimeRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CPUTimeRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPUTimeRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceGroupTag", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceUsageAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceGroupTag = append(m.ResourceGroupTag[:0], dAtA[iNdEx:postIndex]...)
			if m.ResourceGroupTag == nil {
				m.ResourceGroupTag = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowResourceUsageAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RecordListTimestampSec = append(m.RecordListTimestampSec, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowResourceUsageAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthResourceUsageAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowResourceUsageAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RecordListTimestampSec = append(m.RecordListTimestampSec, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordListTimestampSec", wireType)
			}
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowResourceUsageAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RecordListCpuTimeMs = append(m.RecordListCpuTimeMs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowResourceUsageAgent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthResourceUsageAgent
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowResourceUsageAgent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RecordListCpuTimeMs = append(m.RecordListCpuTimeMs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordListCpuTimeMs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceUsageAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipResourceUsageAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthResourceUsageAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResourceUsageAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceUsageAgent
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
					return 0, ErrIntOverflowResourceUsageAgent
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
					return 0, ErrIntOverflowResourceUsageAgent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthResourceUsageAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowResourceUsageAgent
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
				next, err := skipResourceUsageAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthResourceUsageAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceUsageAgent   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("resource_usage_agent.proto", fileDescriptor_resource_usage_agent_623affeb03bbc261)
}

var fileDescriptor_resource_usage_agent_623affeb03bbc261 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x9b, 0xbf, 0x3f, 0x22, 0xc1, 0xa1, 0x25, 0x16, 0xad, 0xb3, 0x18, 0x4a, 0x05, 0x99,
	0x85, 0x8c, 0xa0, 0x6e, 0x5c, 0x6a, 0x11, 0x37, 0x0a, 0x12, 0xa7, 0x4b, 0x09, 0xe3, 0x78, 0x09,
	0x43, 0x9d, 0x26, 0xe4, 0x26, 0x05, 0xdf, 0xc4, 0x47, 0x70, 0xe5, 0x73, 0xb8, 0x74, 0xe9, 0x52,
	0xc6, 0x17, 0x91, 0xc9, 0xb4, 0x6a, 0xa1, 0xae, 0x72, 0xc9, 0x39, 0x5f, 0x72, 0xcf, 0xa1, 0xa1,
	0x01, 0x54, 0xce, 0xe4, 0x20, 0x1c, 0x66, 0x12, 0x44, 0x26, 0x61, 0x6a, 0x13, 0x6d, 0x94, 0x55,
	0xac, 0xb7, 0x4a, 0x0b, 0x7b, 0x52, 0x49, 0xe5, 0x0d, 0x07, 0xf5, 0xd4, 0x78, 0xc3, 0x8e, 0x71,
	0x68, 0xfd, 0xd8, 0x5c, 0x0c, 0x5f, 0x08, 0x0d, 0x46, 0xd7, 0xe3, 0xb4, 0x28, 0x81, 0x43, 0xae,
	0xcc, 0x3d, 0xdb, 0xa7, 0xec, 0xfb, 0x41, 0x69, 0x94, 0xd3, 0xc2, 0x66, 0xb2, 0x4f, 0x06, 0x24,
	0xde, 0xe0, 0xdd, 0x85, 0x72, 0x51, 0x0b, 0x69, 0x26, 0xd9, 0x09, 0xdd, 0x31, 0x9e, 0x13, 0x0f,
	0x05, 0x5a, 0x61, 0x8b, 0x12, 0xd0, 0x66, 0xa5, 0x16, 0x08, 0x79, 0xff, 0xdf, 0xa0, 0x1d, 0xff,
	0xe7, 0x5b, 0x8d, 0xe1, 0xb2, 0x40, 0x9b, 0x2e, 0xe4, 0x1b, 0xc8, 0xd9, 0x31, 0xdd, 0xfe, 0x8d,
	0xe6, 0xda, 0x79, 0x5c, 0x94, 0xd8, 0x6f, 0x0f, 0xda, 0x71, 0xc0, 0x37, 0x7f, 0xc0, 0x91, 0x76,
	0x35, 0x7b, 0x85, 0xc3, 0x0e, 0x0d, 0xce, 0x4b, 0x6d, 0x1f, 0x39, 0xa0, 0x56, 0x53, 0x84, 0x43,
	0xa4, 0x8c, 0xcf, 0xb7, 0x1a, 0xd7, 0xf9, 0x4f, 0xeb, 0xf8, 0xec, 0x96, 0x06, 0x1c, 0xb4, 0x32,
	0x76, 0x1e, 0x8e, 0xed, 0x26, 0x2b, 0x2b, 0x5c, 0xca, 0x1e, 0xfe, 0x61, 0x5a, 0xfa, 0x70, 0xd8,
	0x8a, 0xc9, 0xd9, 0xde, 0xfb, 0xf3, 0x3a, 0x79, 0xad, 0x22, 0xf2, 0x56, 0x45, 0xe4, 0xa3, 0x8a,
	0xc8, 0xd3, 0x67, 0xd4, 0xa2, 0x5d, 0x65, 0x64, 0x62, 0x8b, 0xc9, 0x2c, 0x99, 0xcc, 0x7c, 0xbd,
	0x77, 0x6b, 0xfe, 0x38, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x07, 0x93, 0xf9, 0x74, 0xc0, 0x01,
	0x00, 0x00,
}
