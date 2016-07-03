// Code generated by protoc-gen-go.
// source: logical_camera_image.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogicalCameraImage struct {
	// / \brief Pose of the logical camera.
	Pose *Pose `protobuf:"bytes,1,req,name=pose" json:"pose,omitempty"`
	// / \brief All the models seen by the LogicalCamera
	Model            []*LogicalCameraImage_Model `protobuf:"bytes,2,rep,name=model" json:"model,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *LogicalCameraImage) Reset()                    { *m = LogicalCameraImage{} }
func (m *LogicalCameraImage) String() string            { return proto.CompactTextString(m) }
func (*LogicalCameraImage) ProtoMessage()               {}
func (*LogicalCameraImage) Descriptor() ([]byte, []int) { return fileDescriptor54, []int{0} }

func (m *LogicalCameraImage) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *LogicalCameraImage) GetModel() []*LogicalCameraImage_Model {
	if m != nil {
		return m.Model
	}
	return nil
}

// / \brief Information about a model that is reported by a
// / LogicalCameraSensor
type LogicalCameraImage_Model struct {
	// / \brief Name of the detected model
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// / \brief Pose of the detected model. The pose is relative to the
	// / logical camera's pose.
	Pose             *Pose  `protobuf:"bytes,2,req,name=pose" json:"pose,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogicalCameraImage_Model) Reset()                    { *m = LogicalCameraImage_Model{} }
func (m *LogicalCameraImage_Model) String() string            { return proto.CompactTextString(m) }
func (*LogicalCameraImage_Model) ProtoMessage()               {}
func (*LogicalCameraImage_Model) Descriptor() ([]byte, []int) { return fileDescriptor54, []int{0, 0} }

func (m *LogicalCameraImage_Model) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LogicalCameraImage_Model) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func init() {
	proto.RegisterType((*LogicalCameraImage)(nil), "gazebo.msgs.LogicalCameraImage")
	proto.RegisterType((*LogicalCameraImage_Model)(nil), "gazebo.msgs.LogicalCameraImage.Model")
}

func init() { proto.RegisterFile("logical_camera_image.proto", fileDescriptor54) }

var fileDescriptor54 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xc9, 0x4f, 0xcf,
	0x4c, 0x4e, 0xcc, 0x89, 0x4f, 0x4e, 0xcc, 0x4d, 0x2d, 0x4a, 0x8c, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb,
	0x2d, 0x4e, 0x2f, 0x96, 0xe2, 0x2a, 0xc8, 0x2f, 0x86, 0x4a, 0x28, 0x2d, 0x66, 0xe4, 0x12, 0xf2,
	0x81, 0xe8, 0x73, 0x06, 0x6b, 0xf3, 0x04, 0xe9, 0x12, 0x92, 0xe7, 0x62, 0x01, 0x29, 0x92, 0x60,
	0x54, 0x60, 0xd2, 0xe0, 0x36, 0x12, 0xd4, 0x43, 0xd2, 0xae, 0x17, 0x00, 0x94, 0x10, 0x32, 0xe1,
	0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0x60, 0x52, 0x60, 0x06, 0xaa, 0x50, 0x45, 0x51, 0x81,
	0x69, 0xa0, 0x9e, 0x2f, 0x48, 0xb1, 0x94, 0x19, 0x17, 0x2b, 0x98, 0x21, 0xc4, 0xc3, 0xc5, 0x92,
	0x07, 0x94, 0x05, 0x9b, 0xcf, 0x09, 0xb7, 0x8d, 0x09, 0x87, 0x6d, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6a, 0x29, 0xcf, 0xbb, 0xdb, 0x00, 0x00, 0x00,
}
