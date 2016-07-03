// Code generated by protoc-gen-go.
// source: joint_wrench_stamped.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ForceTorque struct {
	Wrench           []*JointWrench `protobuf:"bytes,1,rep,name=wrench" json:"wrench,omitempty"`
	Time             *Time          `protobuf:"bytes,2,req,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ForceTorque) Reset()                    { *m = ForceTorque{} }
func (m *ForceTorque) String() string            { return proto.CompactTextString(m) }
func (*ForceTorque) ProtoMessage()               {}
func (*ForceTorque) Descriptor() ([]byte, []int) { return fileDescriptor46, []int{0} }

func (m *ForceTorque) GetWrench() []*JointWrench {
	if m != nil {
		return m.Wrench
	}
	return nil
}

func (m *ForceTorque) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*ForceTorque)(nil), "gazebo.msgs.ForceTorque")
}

func init() { proto.RegisterFile("joint_wrench_stamped.proto", fileDescriptor46) }

var fileDescriptor46 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xca, 0xcf, 0xcc,
	0x2b, 0x89, 0x2f, 0x2f, 0x4a, 0xcd, 0x4b, 0xce, 0x88, 0x2f, 0x2e, 0x49, 0xcc, 0x2d, 0x48, 0x4d,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb,
	0x2d, 0x4e, 0x2f, 0x96, 0x12, 0x42, 0x56, 0x08, 0x51, 0x20, 0xc5, 0x55, 0x92, 0x99, 0x9b, 0x0a,
	0x61, 0x2b, 0x45, 0x70, 0x71, 0xbb, 0xe5, 0x17, 0x25, 0xa7, 0x86, 0xe4, 0x17, 0x15, 0x96, 0xa6,
	0x0a, 0x69, 0x70, 0xb1, 0x41, 0x94, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xe8, 0x21,
	0x19, 0xa6, 0xe7, 0x05, 0x32, 0x2b, 0x1c, 0x2c, 0x2f, 0x24, 0xcf, 0xc5, 0x02, 0x32, 0x46, 0x82,
	0x49, 0x81, 0x09, 0xa8, 0x4e, 0x10, 0x45, 0x5d, 0x08, 0x50, 0x02, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0x8c, 0x7b, 0xf8, 0xa3, 0x00, 0x00, 0x00,
}
