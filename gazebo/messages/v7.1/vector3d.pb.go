// Code generated by protoc-gen-go.
// source: vector3d.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Vector3D struct {
	X                *float64 `protobuf:"fixed64,2,req,name=x" json:"x,omitempty"`
	Y                *float64 `protobuf:"fixed64,3,req,name=y" json:"y,omitempty"`
	Z                *float64 `protobuf:"fixed64,4,req,name=z" json:"z,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Vector3D) Reset()                    { *m = Vector3D{} }
func (m *Vector3D) String() string            { return proto.CompactTextString(m) }
func (*Vector3D) ProtoMessage()               {}
func (*Vector3D) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{0} }

func (m *Vector3D) GetX() float64 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Vector3D) GetY() float64 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Vector3D) GetZ() float64 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vector3D)(nil), "gazebo.msgs.Vector3d")
}

func init() { proto.RegisterFile("vector3d.proto", fileDescriptor117) }

var fileDescriptor117 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4b, 0x4d, 0x2e,
	0xc9, 0x2f, 0x32, 0x4e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac, 0x4a,
	0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x56, 0xd2, 0xe6, 0xe2, 0x08, 0x83, 0x4a, 0x0b, 0x71,
	0x72, 0x31, 0x56, 0x48, 0x30, 0x29, 0x30, 0x69, 0x30, 0x82, 0x98, 0x95, 0x12, 0xcc, 0x30, 0x66,
	0x95, 0x04, 0x0b, 0x88, 0x09, 0x08, 0x00, 0x00, 0xff, 0xff, 0x99, 0xe2, 0x7f, 0x04, 0x4a, 0x00,
	0x00, 0x00,
}
