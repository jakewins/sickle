// Code generated by protoc-gen-go.
// source: cessna.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Cessna struct {
	// / \brief Current RPM of the propeller.
	PropellerSpeed *float32 `protobuf:"fixed32,1,opt,name=propeller_speed" json:"propeller_speed,omitempty"`
	// / \brief Current left aileron angle in rads.
	LeftAileron *float32 `protobuf:"fixed32,2,opt,name=left_aileron" json:"left_aileron,omitempty"`
	// / \brief Current left flap angle in rads.
	LeftFlap *float32 `protobuf:"fixed32,3,opt,name=left_flap" json:"left_flap,omitempty"`
	// / \brief Current right aileron angle in rads.
	RightAileron *float32 `protobuf:"fixed32,4,opt,name=right_aileron" json:"right_aileron,omitempty"`
	// / \brief Current right flap angle in rads.
	RightFlap *float32 `protobuf:"fixed32,5,opt,name=right_flap" json:"right_flap,omitempty"`
	// / \brief Current elevators angle in rads.
	Elevators *float32 `protobuf:"fixed32,6,opt,name=elevators" json:"elevators,omitempty"`
	// / \brief Current ruddle angle in rads.
	Rudder *float32 `protobuf:"fixed32,7,opt,name=rudder" json:"rudder,omitempty"`
	// / \brief Target RPM of the propeller.
	CmdPropellerSpeed *float32 `protobuf:"fixed32,8,opt,name=cmd_propeller_speed" json:"cmd_propeller_speed,omitempty"`
	// / \brief Target left aileron angle in rads.
	CmdLeftAileron *float32 `protobuf:"fixed32,9,opt,name=cmd_left_aileron" json:"cmd_left_aileron,omitempty"`
	// / \brief Target left flap angle in rads.
	CmdLeftFlap *float32 `protobuf:"fixed32,10,opt,name=cmd_left_flap" json:"cmd_left_flap,omitempty"`
	// / \brief Target right aileron angle in rads.
	CmdRightAileron *float32 `protobuf:"fixed32,11,opt,name=cmd_right_aileron" json:"cmd_right_aileron,omitempty"`
	// / \brief Target right flap angle in rads.
	CmdRightFlap *float32 `protobuf:"fixed32,12,opt,name=cmd_right_flap" json:"cmd_right_flap,omitempty"`
	// / \brief Target elevators angle in rads.
	CmdElevators *float32 `protobuf:"fixed32,13,opt,name=cmd_elevators" json:"cmd_elevators,omitempty"`
	// / \brief Target ruddle angle in rads.
	CmdRudder        *float32 `protobuf:"fixed32,14,opt,name=cmd_rudder" json:"cmd_rudder,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Cessna) Reset()                    { *m = Cessna{} }
func (m *Cessna) String() string            { return proto.CompactTextString(m) }
func (*Cessna) ProtoMessage()               {}
func (*Cessna) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Cessna) GetPropellerSpeed() float32 {
	if m != nil && m.PropellerSpeed != nil {
		return *m.PropellerSpeed
	}
	return 0
}

func (m *Cessna) GetLeftAileron() float32 {
	if m != nil && m.LeftAileron != nil {
		return *m.LeftAileron
	}
	return 0
}

func (m *Cessna) GetLeftFlap() float32 {
	if m != nil && m.LeftFlap != nil {
		return *m.LeftFlap
	}
	return 0
}

func (m *Cessna) GetRightAileron() float32 {
	if m != nil && m.RightAileron != nil {
		return *m.RightAileron
	}
	return 0
}

func (m *Cessna) GetRightFlap() float32 {
	if m != nil && m.RightFlap != nil {
		return *m.RightFlap
	}
	return 0
}

func (m *Cessna) GetElevators() float32 {
	if m != nil && m.Elevators != nil {
		return *m.Elevators
	}
	return 0
}

func (m *Cessna) GetRudder() float32 {
	if m != nil && m.Rudder != nil {
		return *m.Rudder
	}
	return 0
}

func (m *Cessna) GetCmdPropellerSpeed() float32 {
	if m != nil && m.CmdPropellerSpeed != nil {
		return *m.CmdPropellerSpeed
	}
	return 0
}

func (m *Cessna) GetCmdLeftAileron() float32 {
	if m != nil && m.CmdLeftAileron != nil {
		return *m.CmdLeftAileron
	}
	return 0
}

func (m *Cessna) GetCmdLeftFlap() float32 {
	if m != nil && m.CmdLeftFlap != nil {
		return *m.CmdLeftFlap
	}
	return 0
}

func (m *Cessna) GetCmdRightAileron() float32 {
	if m != nil && m.CmdRightAileron != nil {
		return *m.CmdRightAileron
	}
	return 0
}

func (m *Cessna) GetCmdRightFlap() float32 {
	if m != nil && m.CmdRightFlap != nil {
		return *m.CmdRightFlap
	}
	return 0
}

func (m *Cessna) GetCmdElevators() float32 {
	if m != nil && m.CmdElevators != nil {
		return *m.CmdElevators
	}
	return 0
}

func (m *Cessna) GetCmdRudder() float32 {
	if m != nil && m.CmdRudder != nil {
		return *m.CmdRudder
	}
	return 0
}

func init() {
	proto.RegisterType((*Cessna)(nil), "gazebo.msgs.Cessna")
}

func init() { proto.RegisterFile("cessna.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0x4d, 0x6e, 0x83, 0x30,
	0x10, 0x46, 0x05, 0x6d, 0x69, 0x19, 0x7e, 0x5a, 0xdc, 0x3f, 0x57, 0xdd, 0x54, 0x5d, 0x65, 0xc5,
	0x25, 0x72, 0x10, 0xe4, 0xe0, 0x81, 0x20, 0x19, 0x6c, 0xd9, 0x24, 0x8b, 0x9c, 0x2f, 0x07, 0x0b,
	0x0c, 0x22, 0x84, 0x2c, 0x79, 0xef, 0x1b, 0x78, 0x40, 0x5c, 0xa2, 0x73, 0x9d, 0xc8, 0x8d, 0xd5,
	0xbd, 0x66, 0x51, 0x2d, 0x4e, 0xb8, 0xd3, 0x79, 0xeb, 0x6a, 0xf7, 0x7f, 0xf6, 0x21, 0xd8, 0x92,
	0x65, 0xdf, 0xf0, 0x3a, 0x0c, 0x0c, 0x2a, 0x85, 0xb6, 0x70, 0x06, 0x51, 0x72, 0xef, 0xcf, 0xdb,
	0xf8, 0xec, 0x03, 0x62, 0x85, 0x55, 0x5f, 0x88, 0x66, 0x30, 0xba, 0xe3, 0x3e, 0xd1, 0x0c, 0x42,
	0xa2, 0x95, 0x12, 0x86, 0x3f, 0x10, 0xfa, 0x84, 0xc4, 0x36, 0xf5, 0x7e, 0x59, 0x3e, 0x12, 0x66,
	0x00, 0x13, 0xa6, 0xe9, 0xd3, 0x7c, 0x8d, 0x0a, 0x8f, 0xa2, 0xd7, 0xd6, 0xf1, 0x80, 0x50, 0x0a,
	0x81, 0x3d, 0x48, 0x89, 0x96, 0x3f, 0xd3, 0xf3, 0x2f, 0xbc, 0x97, 0xad, 0x2c, 0xee, 0x9b, 0x5e,
	0x48, 0x72, 0x78, 0x1b, 0xe5, 0xaa, 0x2b, 0x9c, 0x23, 0xae, 0x86, 0x3e, 0x08, 0x84, 0x7f, 0x20,
	0x1b, 0xf1, 0xba, 0x2f, 0x22, 0xf5, 0x05, 0xe9, 0xa2, 0xe8, 0x24, 0xbe, 0x7d, 0xd3, 0xd2, 0x99,
	0xcc, 0xbf, 0x43, 0xf3, 0xa9, 0x35, 0x1d, 0xd9, 0x25, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xe4, 0x60,
	0x56, 0x62, 0x01, 0x00, 0x00,
}