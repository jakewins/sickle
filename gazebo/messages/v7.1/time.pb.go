// Code generated by protoc-gen-go.
// source: time.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Time struct {
	Sec              *int32 `protobuf:"varint,1,req,name=sec" json:"sec,omitempty"`
	Nsec             *int32 `protobuf:"varint,2,req,name=nsec" json:"nsec,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor110, []int{0} }

func (m *Time) GetSec() int32 {
	if m != nil && m.Sec != nil {
		return *m.Sec
	}
	return 0
}

func (m *Time) GetNsec() int32 {
	if m != nil && m.Nsec != nil {
		return *m.Nsec
	}
	return 0
}

func init() {
	proto.RegisterType((*Time)(nil), "gazebo.msgs.Time")
}

func init() { proto.RegisterFile("time.proto", fileDescriptor110) }

var fileDescriptor110 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb,
	0x2d, 0x4e, 0x2f, 0x56, 0x52, 0xe4, 0x62, 0x09, 0x01, 0x4a, 0x09, 0x71, 0x73, 0x31, 0x17, 0xa7,
	0x26, 0x4b, 0x30, 0x2a, 0x30, 0x69, 0xb0, 0x0a, 0xf1, 0x70, 0xb1, 0xe4, 0x81, 0x78, 0x4c, 0x20,
	0x1e, 0x20, 0x00, 0x00, 0xff, 0xff, 0x15, 0x1a, 0x66, 0xc5, 0x3c, 0x00, 0x00, 0x00,
}
