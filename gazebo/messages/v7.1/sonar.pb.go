// Code generated by protoc-gen-go.
// source: sonar.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Sonar struct {
	Frame     *string  `protobuf:"bytes,1,req,name=frame" json:"frame,omitempty"`
	WorldPose *Pose    `protobuf:"bytes,2,req,name=world_pose" json:"world_pose,omitempty"`
	RangeMin  *float64 `protobuf:"fixed64,3,req,name=range_min" json:"range_min,omitempty"`
	RangeMax  *float64 `protobuf:"fixed64,4,req,name=range_max" json:"range_max,omitempty"`
	Radius    *float64 `protobuf:"fixed64,5,req,name=radius" json:"radius,omitempty"`
	Range     *float64 `protobuf:"fixed64,6,req,name=range" json:"range,omitempty"`
	// / Location of the contact in the world frame.
	Contact          *Vector3D `protobuf:"bytes,7,opt,name=contact" json:"contact,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Sonar) Reset()                    { *m = Sonar{} }
func (m *Sonar) String() string            { return proto.CompactTextString(m) }
func (*Sonar) ProtoMessage()               {}
func (*Sonar) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{0} }

func (m *Sonar) GetFrame() string {
	if m != nil && m.Frame != nil {
		return *m.Frame
	}
	return ""
}

func (m *Sonar) GetWorldPose() *Pose {
	if m != nil {
		return m.WorldPose
	}
	return nil
}

func (m *Sonar) GetRangeMin() float64 {
	if m != nil && m.RangeMin != nil {
		return *m.RangeMin
	}
	return 0
}

func (m *Sonar) GetRangeMax() float64 {
	if m != nil && m.RangeMax != nil {
		return *m.RangeMax
	}
	return 0
}

func (m *Sonar) GetRadius() float64 {
	if m != nil && m.Radius != nil {
		return *m.Radius
	}
	return 0
}

func (m *Sonar) GetRange() float64 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *Sonar) GetContact() *Vector3D {
	if m != nil {
		return m.Contact
	}
	return nil
}

func init() {
	proto.RegisterType((*Sonar)(nil), "gazebo.msgs.Sonar")
}

func init() { proto.RegisterFile("sonar.proto", fileDescriptor102) }

var fileDescriptor102 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8e, 0x5d, 0xcb, 0x82, 0x30,
	0x14, 0xc7, 0x99, 0xcf, 0xa3, 0xe2, 0x91, 0x04, 0x07, 0xc1, 0xf0, 0x4a, 0x82, 0xa2, 0xab, 0x5d,
	0xd4, 0x17, 0x09, 0x82, 0x6e, 0x65, 0xe9, 0x12, 0x21, 0x77, 0x64, 0x5b, 0x2f, 0xf4, 0x91, 0xfa,
	0x94, 0xcd, 0xd9, 0x85, 0x5d, 0xfe, 0xdf, 0xce, 0xef, 0x40, 0x6a, 0x50, 0x09, 0xcd, 0x07, 0x8d,
	0x16, 0x69, 0xda, 0x8a, 0x97, 0x3c, 0x23, 0xef, 0x4d, 0x6b, 0x0a, 0x18, 0xd0, 0xc8, 0x29, 0x28,
	0xb2, 0xbb, 0xac, 0x2d, 0xea, 0x7d, 0x33, 0xe9, 0xd5, 0x9b, 0x40, 0x78, 0x1c, 0x87, 0x74, 0x01,
	0xe1, 0x45, 0x8b, 0x5e, 0x32, 0x52, 0x06, 0xdb, 0x84, 0xae, 0x01, 0x1e, 0xa8, 0xaf, 0x4d, 0x35,
	0x8e, 0x59, 0xe0, 0xbc, 0x74, 0x97, 0xf3, 0xd9, 0x59, 0x7e, 0x70, 0x01, 0xcd, 0x21, 0xd1, 0x42,
	0xb5, 0xb2, 0xea, 0x3b, 0xc5, 0xfe, 0x5c, 0x8b, 0xcc, 0x2c, 0xf1, 0x64, 0xff, 0xde, 0xca, 0x20,
	0xd2, 0xa2, 0xe9, 0x6e, 0x86, 0x85, 0x5e, 0x3b, 0x96, 0xaf, 0xb0, 0xc8, 0xcb, 0x0d, 0xc4, 0x35,
	0x2a, 0x2b, 0x6a, 0xcb, 0xe2, 0x92, 0x38, 0xd0, 0xf2, 0x07, 0x74, 0xfa, 0xbe, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x0e, 0x73, 0x0d, 0xe3, 0x00, 0x00, 0x00,
}
