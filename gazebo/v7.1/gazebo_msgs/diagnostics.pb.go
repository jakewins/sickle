// Code generated by protoc-gen-go.
// source: diagnostics.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Diagnostics struct {
	Time             []*Diagnostics_DiagTime `protobuf:"bytes,1,rep,name=time" json:"time,omitempty"`
	RealTime         *Time                   `protobuf:"bytes,2,req,name=real_time" json:"real_time,omitempty"`
	SimTime          *Time                   `protobuf:"bytes,3,req,name=sim_time" json:"sim_time,omitempty"`
	RealTimeFactor   *float64                `protobuf:"fixed64,4,req,name=real_time_factor" json:"real_time_factor,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Diagnostics) Reset()                    { *m = Diagnostics{} }
func (m *Diagnostics) String() string            { return proto.CompactTextString(m) }
func (*Diagnostics) ProtoMessage()               {}
func (*Diagnostics) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *Diagnostics) GetTime() []*Diagnostics_DiagTime {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Diagnostics) GetRealTime() *Time {
	if m != nil {
		return m.RealTime
	}
	return nil
}

func (m *Diagnostics) GetSimTime() *Time {
	if m != nil {
		return m.SimTime
	}
	return nil
}

func (m *Diagnostics) GetRealTimeFactor() float64 {
	if m != nil && m.RealTimeFactor != nil {
		return *m.RealTimeFactor
	}
	return 0
}

type Diagnostics_DiagTime struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Elapsed          *Time   `protobuf:"bytes,2,req,name=elapsed" json:"elapsed,omitempty"`
	Wall             *Time   `protobuf:"bytes,3,req,name=wall" json:"wall,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Diagnostics_DiagTime) Reset()                    { *m = Diagnostics_DiagTime{} }
func (m *Diagnostics_DiagTime) String() string            { return proto.CompactTextString(m) }
func (*Diagnostics_DiagTime) ProtoMessage()               {}
func (*Diagnostics_DiagTime) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0, 0} }

func (m *Diagnostics_DiagTime) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Diagnostics_DiagTime) GetElapsed() *Time {
	if m != nil {
		return m.Elapsed
	}
	return nil
}

func (m *Diagnostics_DiagTime) GetWall() *Time {
	if m != nil {
		return m.Wall
	}
	return nil
}

func init() {
	proto.RegisterType((*Diagnostics)(nil), "gazebo.msgs.Diagnostics")
	proto.RegisterType((*Diagnostics_DiagTime)(nil), "gazebo.msgs.Diagnostics.DiagTime")
}

func init() { proto.RegisterFile("diagnostics.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xc9, 0x4c, 0x4c,
	0xcf, 0xcb, 0x2f, 0x2e, 0xc9, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e,
	0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0x85, 0x48, 0x28, 0x35, 0x31, 0x71, 0x71, 0xbb, 0x20, 0x94, 0x0b, 0xe9, 0x73, 0xb1, 0x80, 0x64,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x14, 0xf5, 0x90, 0xf4, 0xe9, 0x21, 0xa9, 0x03, 0xb3,
	0x43, 0x80, 0x0a, 0x85, 0x54, 0xb8, 0x38, 0x8b, 0x52, 0x13, 0x73, 0xe2, 0xc1, 0xba, 0x98, 0x14,
	0x98, 0x80, 0xba, 0x04, 0x51, 0x74, 0x81, 0x55, 0x29, 0x73, 0x71, 0x14, 0x67, 0xe6, 0x42, 0x14,
	0x31, 0xe3, 0x52, 0x24, 0xc1, 0x25, 0x00, 0x37, 0x2a, 0x3e, 0x2d, 0x31, 0xb9, 0x24, 0xbf, 0x48,
	0x82, 0x05, 0xa8, 0x98, 0x51, 0x2a, 0x96, 0x8b, 0x03, 0x6e, 0x21, 0x0f, 0x17, 0x4b, 0x5e, 0x22,
	0xd8, 0x85, 0x4c, 0x1a, 0x9c, 0x42, 0x4a, 0x5c, 0xec, 0xa9, 0x39, 0x89, 0x05, 0xc5, 0xa9, 0x29,
	0xb8, 0x2d, 0x97, 0xe7, 0x62, 0x29, 0x4f, 0xcc, 0xc9, 0xc1, 0x69, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xca, 0x86, 0x43, 0x2a, 0x31, 0x01, 0x00, 0x00,
}