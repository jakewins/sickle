// Code generated by protoc-gen-go.
// source: gps.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GPS struct {
	Time             *Time    `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	LinkName         *string  `protobuf:"bytes,2,req,name=link_name" json:"link_name,omitempty"`
	LatitudeDeg      *float64 `protobuf:"fixed64,3,req,name=latitude_deg" json:"latitude_deg,omitempty"`
	LongitudeDeg     *float64 `protobuf:"fixed64,4,req,name=longitude_deg" json:"longitude_deg,omitempty"`
	Altitude         *float64 `protobuf:"fixed64,5,req,name=altitude" json:"altitude,omitempty"`
	VelocityEast     *float64 `protobuf:"fixed64,6,opt,name=velocity_east" json:"velocity_east,omitempty"`
	VelocityNorth    *float64 `protobuf:"fixed64,7,opt,name=velocity_north" json:"velocity_north,omitempty"`
	VelocityUp       *float64 `protobuf:"fixed64,8,opt,name=velocity_up" json:"velocity_up,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GPS) Reset()                    { *m = GPS{} }
func (m *GPS) String() string            { return proto.CompactTextString(m) }
func (*GPS) ProtoMessage()               {}
func (*GPS) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *GPS) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *GPS) GetLinkName() string {
	if m != nil && m.LinkName != nil {
		return *m.LinkName
	}
	return ""
}

func (m *GPS) GetLatitudeDeg() float64 {
	if m != nil && m.LatitudeDeg != nil {
		return *m.LatitudeDeg
	}
	return 0
}

func (m *GPS) GetLongitudeDeg() float64 {
	if m != nil && m.LongitudeDeg != nil {
		return *m.LongitudeDeg
	}
	return 0
}

func (m *GPS) GetAltitude() float64 {
	if m != nil && m.Altitude != nil {
		return *m.Altitude
	}
	return 0
}

func (m *GPS) GetVelocityEast() float64 {
	if m != nil && m.VelocityEast != nil {
		return *m.VelocityEast
	}
	return 0
}

func (m *GPS) GetVelocityNorth() float64 {
	if m != nil && m.VelocityNorth != nil {
		return *m.VelocityNorth
	}
	return 0
}

func (m *GPS) GetVelocityUp() float64 {
	if m != nil && m.VelocityUp != nil {
		return *m.VelocityUp
	}
	return 0
}

func init() {
	proto.RegisterType((*GPS)(nil), "gazebo.msgs.GPS")
}

func init() { proto.RegisterFile("gps.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xca, 0xc1, 0x6e, 0x82, 0x30,
	0x1c, 0x80, 0xf1, 0x14, 0xd8, 0x06, 0x7f, 0xb6, 0x65, 0x74, 0x6a, 0x1a, 0x2f, 0x12, 0x4f, 0x9e,
	0x7a, 0xf0, 0x25, 0xbc, 0x9a, 0xe8, 0x9d, 0x54, 0x69, 0x6a, 0x63, 0x69, 0x09, 0x2d, 0x26, 0xfa,
	0x6c, 0x3e, 0x9c, 0x05, 0x22, 0x7a, 0xfd, 0x7d, 0x1f, 0x24, 0xa2, 0xb6, 0xb4, 0x6e, 0x8c, 0x33,
	0x38, 0x15, 0xec, 0xc6, 0x0f, 0x86, 0x56, 0x56, 0xd8, 0x39, 0x38, 0x59, 0xf1, 0x21, 0x2c, 0xef,
	0x08, 0xc2, 0xcd, 0x76, 0x87, 0x17, 0x10, 0x75, 0x4a, 0x50, 0x1e, 0xac, 0xd2, 0x75, 0x46, 0xdf,
	0x7e, 0xba, 0xf7, 0x01, 0x67, 0x90, 0x28, 0xa9, 0xcf, 0x85, 0x66, 0xfe, 0x0a, 0xfc, 0x95, 0xe0,
	0x09, 0x7c, 0x2b, 0xe6, 0xa4, 0x6b, 0x4b, 0x5e, 0x94, 0x5c, 0x90, 0xd0, 0x2b, 0xc2, 0x53, 0xf8,
	0x51, 0x46, 0x8b, 0x17, 0x47, 0x3d, 0xff, 0x41, 0xcc, 0xd4, 0x30, 0x93, 0x8f, 0xe7, 0x78, 0xe1,
	0xca, 0x1c, 0xa5, 0xbb, 0x16, 0x9c, 0x59, 0x47, 0x3e, 0x73, 0xe4, 0x79, 0x06, 0xbf, 0x23, 0x6b,
	0xd3, 0xb8, 0x13, 0xf9, 0xea, 0xfd, 0x1f, 0xd2, 0xd1, 0xdb, 0x9a, 0xc4, 0x1d, 0x3e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0xe9, 0x37, 0x39, 0xe3, 0x00, 0x00, 0x00,
}