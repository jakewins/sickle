// Code generated by protoc-gen-go.
// source: gui_camera.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GUICamera struct {
	Name           *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	ViewController *string      `protobuf:"bytes,2,opt,name=view_controller" json:"view_controller,omitempty"`
	Pose           *Pose        `protobuf:"bytes,3,opt,name=pose" json:"pose,omitempty"`
	Track          *TrackVisual `protobuf:"bytes,4,opt,name=track" json:"track,omitempty"`
	// / \brief Type of projection: "perspective" or "orthographic".
	ProjectionType   *string `protobuf:"bytes,5,opt,name=projection_type" json:"projection_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GUICamera) Reset()                    { *m = GUICamera{} }
func (m *GUICamera) String() string            { return proto.CompactTextString(m) }
func (*GUICamera) ProtoMessage()               {}
func (*GUICamera) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

func (m *GUICamera) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GUICamera) GetViewController() string {
	if m != nil && m.ViewController != nil {
		return *m.ViewController
	}
	return ""
}

func (m *GUICamera) GetPose() *Pose {
	if m != nil {
		return m.Pose
	}
	return nil
}

func (m *GUICamera) GetTrack() *TrackVisual {
	if m != nil {
		return m.Track
	}
	return nil
}

func (m *GUICamera) GetProjectionType() string {
	if m != nil && m.ProjectionType != nil {
		return *m.ProjectionType
	}
	return ""
}

func init() {
	proto.RegisterType((*GUICamera)(nil), "gazebo.msgs.GUICamera")
}

func init() { proto.RegisterFile("gui_camera.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0xcd, 0x8c,
	0x4f, 0x4e, 0xcc, 0x4d, 0x2d, 0x4a, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f,
	0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0xe2, 0x2a, 0xc8, 0x2f, 0x4e, 0x85,
	0x48, 0x48, 0x09, 0x95, 0x14, 0x25, 0x26, 0x67, 0xc7, 0x97, 0x65, 0x16, 0x97, 0x26, 0xe6, 0x40,
	0xc4, 0x94, 0xa6, 0x32, 0x72, 0x71, 0xba, 0x87, 0x7a, 0x3a, 0x83, 0x0d, 0x10, 0xe2, 0xe1, 0x62,
	0xc9, 0x03, 0xb2, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x85, 0xc4, 0xb9, 0xf8, 0xcb, 0x32, 0x53,
	0xcb, 0xe3, 0x93, 0xf3, 0xf3, 0x4a, 0x8a, 0xf2, 0x73, 0x72, 0x52, 0x8b, 0x24, 0x98, 0x14, 0x18,
	0x81, 0x12, 0xf2, 0x5c, 0x2c, 0x20, 0x63, 0x25, 0x98, 0x81, 0x3c, 0x6e, 0x23, 0x41, 0x3d, 0x24,
	0x0b, 0xf5, 0x02, 0x80, 0x12, 0x42, 0xea, 0x5c, 0xac, 0x60, 0xbb, 0x24, 0x58, 0xc0, 0x2a, 0x24,
	0x50, 0x54, 0x84, 0x80, 0x64, 0xc2, 0xc0, 0x8e, 0x00, 0x59, 0x01, 0x74, 0x47, 0x56, 0x6a, 0x72,
	0x49, 0x66, 0x7e, 0x5e, 0x7c, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x2b, 0xc8, 0x0a, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0c, 0x44, 0x9d, 0xef, 0xd7, 0x00, 0x00, 0x00,
}
