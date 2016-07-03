// Code generated by protoc-gen-go.
// source: gps_sensor.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GPSSensor struct {
	// / \brief Position sensing. Consists of horizontal and vertical noise
	// / properties
	Position *GPSSensor_Sensing `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	// / \brief Velocity sensing. Consists of horizontal and vertical noise
	// / properties
	Velocity         *GPSSensor_Sensing `protobuf:"bytes,2,opt,name=velocity" json:"velocity,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *GPSSensor) Reset()                    { *m = GPSSensor{} }
func (m *GPSSensor) String() string            { return proto.CompactTextString(m) }
func (*GPSSensor) ProtoMessage()               {}
func (*GPSSensor) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *GPSSensor) GetPosition() *GPSSensor_Sensing {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *GPSSensor) GetVelocity() *GPSSensor_Sensing {
	if m != nil {
		return m.Velocity
	}
	return nil
}

// / \brief Sensing information
type GPSSensor_Sensing struct {
	// / \brief Horizontal noise
	HorizontalNoise *SensorNoise `protobuf:"bytes,1,opt,name=horizontal_noise" json:"horizontal_noise,omitempty"`
	// / \brief Vertical noise
	VerticalNoise    *SensorNoise `protobuf:"bytes,2,opt,name=vertical_noise" json:"vertical_noise,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GPSSensor_Sensing) Reset()                    { *m = GPSSensor_Sensing{} }
func (m *GPSSensor_Sensing) String() string            { return proto.CompactTextString(m) }
func (*GPSSensor_Sensing) ProtoMessage()               {}
func (*GPSSensor_Sensing) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0, 0} }

func (m *GPSSensor_Sensing) GetHorizontalNoise() *SensorNoise {
	if m != nil {
		return m.HorizontalNoise
	}
	return nil
}

func (m *GPSSensor_Sensing) GetVerticalNoise() *SensorNoise {
	if m != nil {
		return m.VerticalNoise
	}
	return nil
}

func init() {
	proto.RegisterType((*GPSSensor)(nil), "gazebo.msgs.GPSSensor")
	proto.RegisterType((*GPSSensor_Sensing)(nil), "gazebo.msgs.GPSSensor.Sensing")
}

func init() { proto.RegisterFile("gps_sensor.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0x28, 0x8e,
	0x2f, 0x4e, 0xcd, 0x2b, 0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f,
	0xac, 0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0x12, 0x82, 0x48, 0xc5, 0xe7, 0xe5,
	0x67, 0x16, 0xa7, 0x42, 0x14, 0x28, 0x3d, 0x60, 0xe4, 0xe2, 0x74, 0x0f, 0x08, 0x0e, 0x06, 0xcb,
	0x08, 0x19, 0x70, 0x71, 0x14, 0xe4, 0x17, 0x67, 0x96, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x1b, 0xc9, 0xe9, 0x21, 0x99, 0xa0, 0x07, 0x57, 0xa9, 0x07, 0xa2, 0x32, 0xf3, 0xd2,
	0x41, 0x3a, 0xca, 0x52, 0x73, 0xf2, 0x93, 0x33, 0x4b, 0x2a, 0x25, 0x98, 0x88, 0xd1, 0x21, 0x95,
	0xcf, 0xc5, 0x0e, 0xd3, 0x6c, 0xc4, 0x25, 0x90, 0x91, 0x5f, 0x94, 0x59, 0x95, 0x9f, 0x57, 0x92,
	0x98, 0x03, 0x71, 0x16, 0xd4, 0x5a, 0x09, 0x14, 0x43, 0x20, 0x26, 0xf8, 0x81, 0xe4, 0x81, 0x16,
	0xf2, 0x95, 0xa5, 0x16, 0x95, 0x64, 0x26, 0xc3, 0x75, 0x30, 0xe1, 0xd7, 0x01, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x40, 0xcd, 0x8b, 0xbd, 0x16, 0x01, 0x00, 0x00,
}