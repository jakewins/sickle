// Code generated by protoc-gen-go.
// source: material.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Material_ShaderType int32

const (
	Material_VERTEX                   Material_ShaderType = 1
	Material_PIXEL                    Material_ShaderType = 2
	Material_NORMAL_MAP_OBJECT_SPACE  Material_ShaderType = 3
	Material_NORMAL_MAP_TANGENT_SPACE Material_ShaderType = 4
)

var Material_ShaderType_name = map[int32]string{
	1: "VERTEX",
	2: "PIXEL",
	3: "NORMAL_MAP_OBJECT_SPACE",
	4: "NORMAL_MAP_TANGENT_SPACE",
}
var Material_ShaderType_value = map[string]int32{
	"VERTEX":                   1,
	"PIXEL":                    2,
	"NORMAL_MAP_OBJECT_SPACE":  3,
	"NORMAL_MAP_TANGENT_SPACE": 4,
}

func (x Material_ShaderType) Enum() *Material_ShaderType {
	p := new(Material_ShaderType)
	*p = x
	return p
}
func (x Material_ShaderType) String() string {
	return proto.EnumName(Material_ShaderType_name, int32(x))
}
func (x *Material_ShaderType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Material_ShaderType_value, data, "Material_ShaderType")
	if err != nil {
		return err
	}
	*x = Material_ShaderType(value)
	return nil
}
func (Material_ShaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor60, []int{0, 0} }

type Material struct {
	Script           *Material_Script     `protobuf:"bytes,1,opt,name=script" json:"script,omitempty"`
	ShaderType       *Material_ShaderType `protobuf:"varint,2,opt,name=shader_type,enum=gazebo.msgs.Material_ShaderType" json:"shader_type,omitempty"`
	NormalMap        *string              `protobuf:"bytes,3,opt,name=normal_map" json:"normal_map,omitempty"`
	Ambient          *Color               `protobuf:"bytes,4,opt,name=ambient" json:"ambient,omitempty"`
	Diffuse          *Color               `protobuf:"bytes,5,opt,name=diffuse" json:"diffuse,omitempty"`
	Specular         *Color               `protobuf:"bytes,6,opt,name=specular" json:"specular,omitempty"`
	Emissive         *Color               `protobuf:"bytes,7,opt,name=emissive" json:"emissive,omitempty"`
	Lighting         *bool                `protobuf:"varint,8,opt,name=lighting" json:"lighting,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Material) Reset()                    { *m = Material{} }
func (m *Material) String() string            { return proto.CompactTextString(m) }
func (*Material) ProtoMessage()               {}
func (*Material) Descriptor() ([]byte, []int) { return fileDescriptor60, []int{0} }

func (m *Material) GetScript() *Material_Script {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *Material) GetShaderType() Material_ShaderType {
	if m != nil && m.ShaderType != nil {
		return *m.ShaderType
	}
	return Material_VERTEX
}

func (m *Material) GetNormalMap() string {
	if m != nil && m.NormalMap != nil {
		return *m.NormalMap
	}
	return ""
}

func (m *Material) GetAmbient() *Color {
	if m != nil {
		return m.Ambient
	}
	return nil
}

func (m *Material) GetDiffuse() *Color {
	if m != nil {
		return m.Diffuse
	}
	return nil
}

func (m *Material) GetSpecular() *Color {
	if m != nil {
		return m.Specular
	}
	return nil
}

func (m *Material) GetEmissive() *Color {
	if m != nil {
		return m.Emissive
	}
	return nil
}

func (m *Material) GetLighting() bool {
	if m != nil && m.Lighting != nil {
		return *m.Lighting
	}
	return false
}

type Material_Script struct {
	Uri              []string `protobuf:"bytes,1,rep,name=uri" json:"uri,omitempty"`
	Name             *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Material_Script) Reset()                    { *m = Material_Script{} }
func (m *Material_Script) String() string            { return proto.CompactTextString(m) }
func (*Material_Script) ProtoMessage()               {}
func (*Material_Script) Descriptor() ([]byte, []int) { return fileDescriptor60, []int{0, 0} }

func (m *Material_Script) GetUri() []string {
	if m != nil {
		return m.Uri
	}
	return nil
}

func (m *Material_Script) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Material)(nil), "gazebo.msgs.Material")
	proto.RegisterType((*Material_Script)(nil), "gazebo.msgs.Material.Script")
	proto.RegisterEnum("gazebo.msgs.Material_ShaderType", Material_ShaderType_name, Material_ShaderType_value)
}

func init() { proto.RegisterFile("material.proto", fileDescriptor60) }

var fileDescriptor60 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0xd0, 0x4f, 0x6b, 0xc2, 0x30,
	0x18, 0x06, 0x70, 0xb4, 0x5a, 0xeb, 0xdb, 0x21, 0x25, 0x97, 0x05, 0xe7, 0x41, 0x74, 0x07, 0x0f,
	0xa3, 0x07, 0x61, 0x1f, 0xa0, 0x93, 0x30, 0x36, 0xfc, 0x87, 0x96, 0xe1, 0x69, 0x25, 0x6a, 0xac,
	0x81, 0xa6, 0x29, 0x49, 0x1c, 0x6c, 0x9f, 0x6a, 0x1f, 0x71, 0xd1, 0x29, 0xb8, 0x43, 0x77, 0x7e,
	0x7e, 0x4f, 0xf2, 0xbe, 0x2f, 0xb4, 0x04, 0x35, 0x4c, 0x71, 0x9a, 0x85, 0x85, 0x92, 0x46, 0x22,
	0x3f, 0xa5, 0x5f, 0x6c, 0x2d, 0x43, 0xa1, 0x53, 0xdd, 0xf6, 0x37, 0x32, 0x93, 0xea, 0x37, 0xe9,
	0x7d, 0x3b, 0xe0, 0x4d, 0xce, 0x18, 0x3d, 0x80, 0xab, 0x37, 0x8a, 0x17, 0x06, 0x57, 0xba, 0x95,
	0x81, 0x3f, 0xec, 0x84, 0x57, 0xbd, 0xf0, 0xc2, 0xc2, 0xe5, 0xc9, 0xa0, 0x47, 0xf0, 0xf5, 0x9e,
	0x6e, 0x99, 0x4a, 0xcc, 0x67, 0xc1, 0x70, 0xd5, 0x56, 0x5a, 0xc3, 0x6e, 0x49, 0xe5, 0x04, 0x63,
	0xeb, 0x10, 0x02, 0xc8, 0xa5, 0x12, 0x34, 0x4b, 0x04, 0x2d, 0xb0, 0x63, 0x5b, 0x4d, 0xd4, 0x87,
	0x06, 0x15, 0x6b, 0xce, 0x72, 0x83, 0x6b, 0xa7, 0x9f, 0xd1, 0x9f, 0x67, 0x46, 0xc7, 0x81, 0x8f,
	0x68, 0xcb, 0x77, 0xbb, 0x83, 0x66, 0xb8, 0x5e, 0x8a, 0xee, 0xc1, 0xd3, 0x05, 0xdb, 0x1c, 0x32,
	0xaa, 0xb0, 0xfb, 0x9f, 0x62, 0x82, 0x6b, 0xcd, 0x3f, 0x18, 0x6e, 0x94, 0xaa, 0x00, 0xbc, 0x8c,
	0xa7, 0x7b, 0xc3, 0xf3, 0x14, 0x7b, 0x56, 0x79, 0xed, 0x3e, 0xb8, 0xe7, 0xe5, 0x7d, 0x70, 0x0e,
	0x8a, 0xdb, 0x3b, 0x39, 0x76, 0xfc, 0x1b, 0xa8, 0xe5, 0x54, 0x1c, 0x4f, 0x50, 0x1d, 0x34, 0x7b,
	0xef, 0x00, 0x57, 0xeb, 0x02, 0xb8, 0x6f, 0x64, 0x11, 0x93, 0x55, 0x50, 0x41, 0x4d, 0xa8, 0xcf,
	0x5f, 0x56, 0x64, 0x1c, 0x54, 0xd1, 0x1d, 0xdc, 0x4e, 0x67, 0x8b, 0x49, 0x34, 0x4e, 0x26, 0xd1,
	0x3c, 0x99, 0x3d, 0xbd, 0x92, 0x51, 0x9c, 0x2c, 0xe7, 0xd1, 0x88, 0x04, 0x0e, 0xea, 0x00, 0xbe,
	0x0a, 0xe3, 0x68, 0xfa, 0x4c, 0xa6, 0x97, 0xb4, 0xf6, 0x13, 0x00, 0x00, 0xff, 0xff, 0x93, 0x29,
	0x2d, 0x0b, 0xdd, 0x01, 0x00, 0x00,
}
