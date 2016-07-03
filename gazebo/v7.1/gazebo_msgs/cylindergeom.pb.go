// Code generated by protoc-gen-go.
// source: cylindergeom.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CylinderGeom struct {
	Radius           *float64 `protobuf:"fixed64,1,req,name=radius" json:"radius,omitempty"`
	Length           *float64 `protobuf:"fixed64,2,req,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CylinderGeom) Reset()                    { *m = CylinderGeom{} }
func (m *CylinderGeom) String() string            { return proto.CompactTextString(m) }
func (*CylinderGeom) ProtoMessage()               {}
func (*CylinderGeom) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *CylinderGeom) GetRadius() float64 {
	if m != nil && m.Radius != nil {
		return *m.Radius
	}
	return 0
}

func (m *CylinderGeom) GetLength() float64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*CylinderGeom)(nil), "gazebo.msgs.CylinderGeom")
}

func init() { proto.RegisterFile("cylindergeom.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xae, 0xcc, 0xc9,
	0xcc, 0x4b, 0x49, 0x2d, 0x4a, 0x4f, 0xcd, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x56, 0xd2, 0xe3, 0xe2, 0x71,
	0x86, 0x2a, 0x71, 0x07, 0x2a, 0x11, 0xe2, 0xe3, 0x62, 0x2b, 0x4a, 0x4c, 0xc9, 0x2c, 0x2d, 0x96,
	0x60, 0x54, 0x60, 0xd2, 0x60, 0x04, 0xf1, 0x73, 0x52, 0xf3, 0xd2, 0x4b, 0x32, 0x24, 0x98, 0x40,
	0x7c, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x45, 0x53, 0xba, 0x51, 0x00, 0x00, 0x00,
}