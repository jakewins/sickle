// Code generated by protoc-gen-go.
// source: model_configuration.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ModelConfiguration struct {
	Time             *Time     `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	JointNames       []string  `protobuf:"bytes,2,rep,name=joint_names" json:"joint_names,omitempty"`
	JointPositions   []float64 `protobuf:"fixed64,3,rep,name=joint_positions" json:"joint_positions,omitempty"`
	Pose             *Pose     `protobuf:"bytes,4,opt,name=pose" json:"pose,omitempty"`
	LinkName         *string   `protobuf:"bytes,5,opt,name=link_name" json:"link_name,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ModelConfiguration) Reset()                    { *m = ModelConfiguration{} }
func (m *ModelConfiguration) String() string            { return proto.CompactTextString(m) }
func (*ModelConfiguration) ProtoMessage()               {}
func (*ModelConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor62, []int{0} }

func (m *ModelConfiguration) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *ModelConfiguration) GetJointNames() []string {
	if m != nil {
		return m.JointNames
	}
	return nil
}

func (m *ModelConfiguration) GetJointPositions() []float64 {
	if m != nil {
		return m.JointPositions
	}
	return nil
}

func (m *ModelConfiguration) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *ModelConfiguration) GetLinkName() string {
	if m != nil && m.LinkName != nil {
		return *m.LinkName
	}
	return ""
}

func init() {
	proto.RegisterType((*ModelConfiguration)(nil), "gazebo.msgs.ModelConfiguration")
}

func init() { proto.RegisterFile("model_configuration.proto", fileDescriptor62) }

var fileDescriptor62 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0x89, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d,
	0x4e, 0x2f, 0x96, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d, 0x85, 0x48, 0x48, 0x71, 0x15, 0xe4, 0x17, 0x43,
	0xd9, 0x4a, 0xd3, 0x19, 0xb9, 0x84, 0x7c, 0x41, 0x46, 0x38, 0x23, 0x9b, 0x20, 0x24, 0xcf, 0xc5,
	0x02, 0xd2, 0x20, 0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0x6d, 0x24, 0xa8, 0x87, 0x64, 0x94, 0x5e, 0x08,
	0x50, 0x42, 0x48, 0x98, 0x8b, 0x3b, 0x2b, 0x3f, 0x33, 0xaf, 0x24, 0x3e, 0x2f, 0x31, 0x37, 0xb5,
	0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x53, 0x48, 0x9c, 0x8b, 0x1f, 0x22, 0x08, 0xb4, 0x20, 0x13,
	0x64, 0x4e, 0xb1, 0x04, 0x33, 0x50, 0x82, 0x11, 0x64, 0x1c, 0xc8, 0x4e, 0x09, 0x16, 0x05, 0x46,
	0x0c, 0xe3, 0x02, 0x80, 0x12, 0x42, 0x82, 0x5c, 0x9c, 0x39, 0x99, 0x79, 0xd9, 0x60, 0xd3, 0x24,
	0x58, 0x81, 0xaa, 0x38, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x76, 0xc5, 0x62, 0xda, 0x00,
	0x00, 0x00,
}
