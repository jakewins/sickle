// Code generated by protoc-gen-go.
// source: camera_lens.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CameraLens struct {
	// / \brief Type of projection of the lens
	// /        possible values are "gnomonical", "stereographic", "equidistant",
	// /        "equisolid_angle", "stereographic", "custom".
	// /        If you set this value to "custom" you need to specify at least one
	// /        of the `c1`, `c2`, `c3`, `f` or `fun`.
	Type *string `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	// / \brief Linear image scaling factor
	C1 *float64 `protobuf:"fixed64,2,opt,name=c1" json:"c1,omitempty"`
	// / \brief Angle scaling factor
	C2 *float64 `protobuf:"fixed64,3,opt,name=c2" json:"c2,omitempty"`
	// / \brief Angle offset factor
	C3 *float64 `protobuf:"fixed64,4,opt,name=c3" json:"c3,omitempty"`
	// / \brief Linear scaling factor, unlike `c1`, will be adjusted to match hfov
	// /        if scale_to_fov is set to `true`.
	F *float64 `protobuf:"fixed64,5,opt,name=f" json:"f,omitempty"`
	// / \brief Angle modification function
	//         possible values are "tan", "sin" and "id".
	Fun *string `protobuf:"bytes,6,opt,name=fun" json:"fun,omitempty"`
	// / \brief Scale image to fit horizontal FOV
	ScaleToHfov *bool `protobuf:"varint,7,opt,name=scale_to_hfov" json:"scale_to_hfov,omitempty"`
	// / \brief Everything outside of this angle will be hidden,
	// /        the angle is counted from camera's X (forward) axis.
	CutoffAngle *float64 `protobuf:"fixed64,8,opt,name=cutoff_angle" json:"cutoff_angle,omitempty"`
	// / \brief Horizontal field of view in radians.
	Hfov *float64 `protobuf:"fixed64,9,opt,name=hfov" json:"hfov,omitempty"`
	// / \brief Size of cube map texture,
	// /        used to store intermediate rendering result.
	EnvTextureSize   *int32 `protobuf:"varint,10,opt,name=env_texture_size" json:"env_texture_size,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CameraLens) Reset()                    { *m = CameraLens{} }
func (m *CameraLens) String() string            { return proto.CompactTextString(m) }
func (*CameraLens) ProtoMessage()               {}
func (*CameraLens) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *CameraLens) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *CameraLens) GetC1() float64 {
	if m != nil && m.C1 != nil {
		return *m.C1
	}
	return 0
}

func (m *CameraLens) GetC2() float64 {
	if m != nil && m.C2 != nil {
		return *m.C2
	}
	return 0
}

func (m *CameraLens) GetC3() float64 {
	if m != nil && m.C3 != nil {
		return *m.C3
	}
	return 0
}

func (m *CameraLens) GetF() float64 {
	if m != nil && m.F != nil {
		return *m.F
	}
	return 0
}

func (m *CameraLens) GetFun() string {
	if m != nil && m.Fun != nil {
		return *m.Fun
	}
	return ""
}

func (m *CameraLens) GetScaleToHfov() bool {
	if m != nil && m.ScaleToHfov != nil {
		return *m.ScaleToHfov
	}
	return false
}

func (m *CameraLens) GetCutoffAngle() float64 {
	if m != nil && m.CutoffAngle != nil {
		return *m.CutoffAngle
	}
	return 0
}

func (m *CameraLens) GetHfov() float64 {
	if m != nil && m.Hfov != nil {
		return *m.Hfov
	}
	return 0
}

func (m *CameraLens) GetEnvTextureSize() int32 {
	if m != nil && m.EnvTextureSize != nil {
		return *m.EnvTextureSize
	}
	return 0
}

func init() {
	proto.RegisterType((*CameraLens)(nil), "gazebo.msgs.CameraLens")
}

func init() { proto.RegisterFile("camera_lens.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8d, 0x4b, 0x6a, 0xc3, 0x30,
	0x10, 0x40, 0x91, 0x3f, 0xad, 0x35, 0x76, 0xa1, 0x15, 0x2d, 0xcc, 0xd2, 0x74, 0xe5, 0x95, 0xa1,
	0xcd, 0x11, 0xb2, 0xcd, 0x1d, 0x84, 0x22, 0x46, 0x4e, 0xc0, 0x96, 0x8c, 0x25, 0x9b, 0xc4, 0x57,
	0xca, 0x25, 0x23, 0xb4, 0xc8, 0xee, 0xbd, 0x99, 0x37, 0x0c, 0x7c, 0x69, 0x35, 0xd1, 0xa2, 0xe4,
	0x48, 0xd6, 0xf7, 0xf3, 0xe2, 0x82, 0x13, 0xf5, 0xa0, 0x76, 0x3a, 0xbb, 0x7e, 0xf2, 0x83, 0xff,
	0x7d, 0x30, 0x80, 0x63, 0x4a, 0x4e, 0xb1, 0x10, 0x0d, 0x14, 0xe1, 0x3e, 0x13, 0xb2, 0x36, 0xeb,
	0xb8, 0x00, 0xc8, 0xf4, 0x1f, 0x66, 0x2d, 0xeb, 0x58, 0xe2, 0x7f, 0xcc, 0x5f, 0x7c, 0xc0, 0x22,
	0x31, 0x07, 0x66, 0xb0, 0x4c, 0x58, 0x43, 0x6e, 0x56, 0x8b, 0x6f, 0x51, 0xb8, 0xf8, 0x81, 0x0f,
	0xaf, 0xd5, 0x48, 0x32, 0x38, 0x79, 0x31, 0x6e, 0xc3, 0xf7, 0x38, 0xae, 0xc4, 0x37, 0x34, 0x7a,
	0x0d, 0xce, 0x18, 0xa9, 0xec, 0x30, 0x12, 0x56, 0xe9, 0x32, 0xbe, 0x4d, 0x0d, 0x4f, 0x86, 0xf0,
	0x49, 0x76, 0x93, 0x81, 0x6e, 0x61, 0x5d, 0x48, 0xfa, 0xeb, 0x4e, 0x08, 0x71, 0x53, 0x3e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x6d, 0xaa, 0xbf, 0x00, 0xce, 0x00, 0x00, 0x00,
}
