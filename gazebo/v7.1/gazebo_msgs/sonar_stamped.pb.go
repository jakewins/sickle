// Code generated by protoc-gen-go.
// source: sonar_stamped.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SonarStamped struct {
	// Time when the data was captured
	Time             *Time  `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	Sonar            *Sonar `protobuf:"bytes,2,req,name=sonar" json:"sonar,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SonarStamped) Reset()                    { *m = SonarStamped{} }
func (m *SonarStamped) String() string            { return proto.CompactTextString(m) }
func (*SonarStamped) ProtoMessage()               {}
func (*SonarStamped) Descriptor() ([]byte, []int) { return fileDescriptor103, []int{0} }

func (m *SonarStamped) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *SonarStamped) GetSonar() *Sonar {
	if m != nil {
		return m.Sonar
	}
	return nil
}

func init() {
	proto.RegisterType((*SonarStamped)(nil), "gazebo.msgs.SonarStamped")
}

func init() { proto.RegisterFile("sonar_stamped.proto", fileDescriptor103) }

var fileDescriptor103 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcf, 0x4b,
	0x2c, 0x8a, 0x2f, 0x2e, 0x49, 0xcc, 0x2d, 0x48, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0xe2, 0x2a, 0xc9,
	0xcc, 0x4d, 0x85, 0x48, 0x48, 0x71, 0x83, 0x55, 0x43, 0x38, 0x4a, 0x41, 0x5c, 0x3c, 0xc1, 0x20,
	0x6e, 0x30, 0x44, 0xaf, 0x90, 0x3c, 0x17, 0x0b, 0x48, 0xa9, 0x04, 0xa3, 0x02, 0x93, 0x06, 0xb7,
	0x91, 0xa0, 0x1e, 0x92, 0x21, 0x7a, 0x21, 0x40, 0x09, 0x21, 0x45, 0x2e, 0x56, 0xb0, 0x7e, 0x09,
	0x26, 0xb0, 0x0a, 0x21, 0x14, 0x15, 0x60, 0xa3, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x8f,
	0xd4, 0x4d, 0x8f, 0x00, 0x00, 0x00,
}