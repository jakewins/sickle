// Code generated by protoc-gen-go.
// source: log_playback_control.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogPlaybackControl struct {
	// / \brief Pause/play the log file.
	Pause *bool `protobuf:"varint,1,opt,name=pause" json:"pause,omitempty"`
	// / \brief Make a relative jump. The value indicates the number of
	// /        iterations that will be executed at once. If a negative
	// /        value is specified, the playback will jump backwards.
	MultiStep *int32 `protobuf:"zigzag32,2,opt,name=multi_step" json:"multi_step,omitempty"`
	// / \brief Jump to the beginning of the log file.
	Rewind *bool `protobuf:"varint,3,opt,name=rewind" json:"rewind,omitempty"`
	// / \brief Jump to the end of the log file.
	Forward *bool `protobuf:"varint,4,opt,name=forward" json:"forward,omitempty"`
	// / \brief Jump to a specific simulation time in the log file. The
	// /        playback service will load the frame with the closest
	// /        simulation time bigger than the "seek" value.
	Seek             *Time  `protobuf:"bytes,5,opt,name=seek" json:"seek,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogPlaybackControl) Reset()                    { *m = LogPlaybackControl{} }
func (m *LogPlaybackControl) String() string            { return proto.CompactTextString(m) }
func (*LogPlaybackControl) ProtoMessage()               {}
func (*LogPlaybackControl) Descriptor() ([]byte, []int) { return fileDescriptor56, []int{0} }

func (m *LogPlaybackControl) GetPause() bool {
	if m != nil && m.Pause != nil {
		return *m.Pause
	}
	return false
}

func (m *LogPlaybackControl) GetMultiStep() int32 {
	if m != nil && m.MultiStep != nil {
		return *m.MultiStep
	}
	return 0
}

func (m *LogPlaybackControl) GetRewind() bool {
	if m != nil && m.Rewind != nil {
		return *m.Rewind
	}
	return false
}

func (m *LogPlaybackControl) GetForward() bool {
	if m != nil && m.Forward != nil {
		return *m.Forward
	}
	return false
}

func (m *LogPlaybackControl) GetSeek() *Time {
	if m != nil {
		return m.Seek
	}
	return nil
}

func init() {
	proto.RegisterType((*LogPlaybackControl)(nil), "gazebo.msgs.LogPlaybackControl")
}

func init() { proto.RegisterFile("log_playback_control.proto", fileDescriptor56) }

var fileDescriptor56 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x0f, 0x82, 0x30,
	0x10, 0x46, 0x83, 0x82, 0x9a, 0x23, 0x6a, 0xe8, 0xd4, 0xb0, 0x68, 0x9c, 0x9c, 0x3a, 0xf8, 0x17,
	0x5c, 0x1d, 0x1c, 0xdc, 0x49, 0x81, 0xb3, 0x21, 0x14, 0xae, 0x69, 0x4b, 0x08, 0xfe, 0x7a, 0x1b,
	0x70, 0x70, 0xbc, 0x77, 0xdf, 0x4b, 0x1e, 0xe4, 0x9a, 0x54, 0x61, 0xb4, 0x9c, 0x4a, 0x59, 0xb5,
	0x45, 0x45, 0xbd, 0xb7, 0xa4, 0x85, 0xb1, 0xe4, 0x89, 0xa5, 0x4a, 0x7e, 0xb0, 0x24, 0xd1, 0x39,
	0xe5, 0x72, 0xf0, 0x4d, 0x87, 0xcb, 0xe3, 0x32, 0x01, 0x7b, 0x90, 0x7a, 0xfe, 0xac, 0xfb, 0x22,
	0xb1, 0x3d, 0x24, 0x46, 0x0e, 0x0e, 0x79, 0x74, 0x8e, 0xae, 0x3b, 0xc6, 0x00, 0xba, 0x41, 0xfb,
	0xa6, 0x70, 0x1e, 0x0d, 0x5f, 0x05, 0x96, 0xb1, 0x03, 0x6c, 0x2c, 0x8e, 0x4d, 0x5f, 0xf3, 0xf5,
	0xbc, 0x39, 0xc2, 0xf6, 0x4d, 0x76, 0x94, 0xb6, 0xe6, 0xf1, 0x0c, 0x4e, 0x10, 0x3b, 0xc4, 0x96,
	0x27, 0xe1, 0x4a, 0x6f, 0x99, 0xf8, 0x2b, 0x10, 0xaf, 0x10, 0xf0, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x19, 0xcb, 0xd7, 0xea, 0xb0, 0x00, 0x00, 0x00,
}