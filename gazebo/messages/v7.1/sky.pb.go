// Code generated by protoc-gen-go.
// source: sky.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Sky struct {
	Time             *float64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Sunrise          *float64 `protobuf:"fixed64,2,opt,name=sunrise" json:"sunrise,omitempty"`
	Sunset           *float64 `protobuf:"fixed64,3,opt,name=sunset" json:"sunset,omitempty"`
	WindSpeed        *float64 `protobuf:"fixed64,4,opt,name=wind_speed" json:"wind_speed,omitempty"`
	WindDirection    *float64 `protobuf:"fixed64,5,opt,name=wind_direction" json:"wind_direction,omitempty"`
	CloudAmbient     *Color   `protobuf:"bytes,6,opt,name=cloud_ambient" json:"cloud_ambient,omitempty"`
	Humidity         *float64 `protobuf:"fixed64,7,opt,name=humidity" json:"humidity,omitempty"`
	MeanCloudSize    *float64 `protobuf:"fixed64,8,opt,name=mean_cloud_size" json:"mean_cloud_size,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Sky) Reset()                    { *m = Sky{} }
func (m *Sky) String() string            { return proto.CompactTextString(m) }
func (*Sky) ProtoMessage()               {}
func (*Sky) Descriptor() ([]byte, []int) { return fileDescriptor101, []int{0} }

func (m *Sky) GetTime() float64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Sky) GetSunrise() float64 {
	if m != nil && m.Sunrise != nil {
		return *m.Sunrise
	}
	return 0
}

func (m *Sky) GetSunset() float64 {
	if m != nil && m.Sunset != nil {
		return *m.Sunset
	}
	return 0
}

func (m *Sky) GetWindSpeed() float64 {
	if m != nil && m.WindSpeed != nil {
		return *m.WindSpeed
	}
	return 0
}

func (m *Sky) GetWindDirection() float64 {
	if m != nil && m.WindDirection != nil {
		return *m.WindDirection
	}
	return 0
}

func (m *Sky) GetCloudAmbient() *Color {
	if m != nil {
		return m.CloudAmbient
	}
	return nil
}

func (m *Sky) GetHumidity() float64 {
	if m != nil && m.Humidity != nil {
		return *m.Humidity
	}
	return 0
}

func (m *Sky) GetMeanCloudSize() float64 {
	if m != nil && m.MeanCloudSize != nil {
		return *m.MeanCloudSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Sky)(nil), "gazebo.msgs.Sky")
}

func init() { proto.RegisterFile("sky.proto", fileDescriptor101) }

var fileDescriptor101 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xce, 0x41, 0x4e, 0xc3, 0x30,
	0x10, 0x05, 0x50, 0x99, 0x96, 0xb6, 0x4c, 0xa0, 0x45, 0x5e, 0xc0, 0x88, 0x15, 0x62, 0x05, 0x1b,
	0x2f, 0xb8, 0x02, 0x47, 0xe0, 0x00, 0x51, 0x12, 0x8f, 0xc2, 0x28, 0xb1, 0x1d, 0x79, 0x1c, 0xa1,
	0xe4, 0x60, 0x9c, 0x8f, 0xc4, 0xd9, 0x74, 0x39, 0xef, 0x6b, 0xbe, 0x3e, 0xdc, 0x49, 0x37, 0x99,
	0x21, 0x86, 0x14, 0x74, 0xd1, 0x56, 0x33, 0xd5, 0xc1, 0x38, 0x69, 0xe5, 0xa5, 0x68, 0x42, 0x1f,
	0xe2, 0x96, 0xbc, 0xfd, 0x29, 0xd8, 0x7d, 0x77, 0x93, 0xbe, 0x87, 0x7d, 0x62, 0x47, 0xa8, 0x5e,
	0xd5, 0xbb, 0xd2, 0x17, 0x38, 0xca, 0xe8, 0x23, 0x0b, 0xe1, 0x4d, 0x86, 0x33, 0x1c, 0x16, 0x10,
	0x4a, 0xb8, 0xcb, 0xb7, 0x06, 0xf8, 0x65, 0x6f, 0x4b, 0x19, 0x88, 0x2c, 0xee, 0xb3, 0x3d, 0xc1,
	0x39, 0x9b, 0xe5, 0x48, 0x4d, 0xe2, 0xe0, 0xf1, 0x36, 0xfb, 0x07, 0x3c, 0x34, 0x7d, 0x18, 0x6d,
	0x59, 0xb9, 0x9a, 0xc9, 0x27, 0x3c, 0x2c, 0x5c, 0x7c, 0x6a, 0x73, 0x35, 0xca, 0x7c, 0xad, 0x9b,
	0xf4, 0x23, 0x9c, 0x7e, 0x46, 0xc7, 0x96, 0xd3, 0x84, 0xc7, 0xfc, 0xfc, 0x0c, 0x17, 0x47, 0x95,
	0x2f, 0xb7, 0x06, 0xe1, 0x99, 0xf0, 0xb4, 0x06, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0x3e,
	0xa1, 0x2c, 0xde, 0x00, 0x00, 0x00,
}