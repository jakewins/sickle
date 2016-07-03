// Code generated by protoc-gen-go.
// source: world_reset.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WorldReset struct {
	All              *bool  `protobuf:"varint,1,opt,name=all,def=1" json:"all,omitempty"`
	TimeOnly         *bool  `protobuf:"varint,2,opt,name=time_only,def=0" json:"time_only,omitempty"`
	ModelOnly        *bool  `protobuf:"varint,3,opt,name=model_only,def=0" json:"model_only,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WorldReset) Reset()                    { *m = WorldReset{} }
func (m *WorldReset) String() string            { return proto.CompactTextString(m) }
func (*WorldReset) ProtoMessage()               {}
func (*WorldReset) Descriptor() ([]byte, []int) { return fileDescriptor124, []int{0} }

const Default_WorldReset_All bool = true
const Default_WorldReset_TimeOnly bool = false
const Default_WorldReset_ModelOnly bool = false

func (m *WorldReset) GetAll() bool {
	if m != nil && m.All != nil {
		return *m.All
	}
	return Default_WorldReset_All
}

func (m *WorldReset) GetTimeOnly() bool {
	if m != nil && m.TimeOnly != nil {
		return *m.TimeOnly
	}
	return Default_WorldReset_TimeOnly
}

func (m *WorldReset) GetModelOnly() bool {
	if m != nil && m.ModelOnly != nil {
		return *m.ModelOnly
	}
	return Default_WorldReset_ModelOnly
}

func init() {
	proto.RegisterType((*WorldReset)(nil), "gazebo.msgs.WorldReset")
}

func init() { proto.RegisterFile("world_reset.proto", fileDescriptor124) }

var fileDescriptor124 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xcf, 0x2f, 0xca,
	0x49, 0x89, 0x2f, 0x4a, 0x2d, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e,
	0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x56, 0x0a, 0xe1, 0xe2, 0x0a, 0x07,
	0xa9, 0x08, 0x02, 0x29, 0x10, 0x12, 0xe4, 0x62, 0x4e, 0xcc, 0xc9, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0xb0, 0x62, 0x29, 0x29, 0x2a, 0x4d, 0x15, 0x92, 0xe0, 0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x8d,
	0xcf, 0xcf, 0xcb, 0xa9, 0x94, 0x60, 0x02, 0x4b, 0xb0, 0xa6, 0x25, 0xe6, 0x14, 0xa7, 0x0a, 0x49,
	0x72, 0x71, 0xe5, 0xe6, 0xa7, 0xa4, 0xe6, 0x40, 0xa4, 0x98, 0x91, 0xa4, 0x00, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0x88, 0xd6, 0x25, 0x76, 0x00, 0x00, 0x00,
}
