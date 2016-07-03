// Code generated by protoc-gen-go.
// source: world_stats.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WorldStatistics struct {
	SimTime          *Time                  `protobuf:"bytes,2,req,name=sim_time" json:"sim_time,omitempty"`
	PauseTime        *Time                  `protobuf:"bytes,3,req,name=pause_time" json:"pause_time,omitempty"`
	RealTime         *Time                  `protobuf:"bytes,4,req,name=real_time" json:"real_time,omitempty"`
	Paused           *bool                  `protobuf:"varint,5,req,name=paused" json:"paused,omitempty"`
	Iterations       *uint64                `protobuf:"varint,6,req,name=iterations" json:"iterations,omitempty"`
	ModelCount       *int32                 `protobuf:"varint,7,opt,name=model_count" json:"model_count,omitempty"`
	LogPlaybackStats *LogPlaybackStatistics `protobuf:"bytes,8,opt,name=log_playback_stats" json:"log_playback_stats,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *WorldStatistics) Reset()                    { *m = WorldStatistics{} }
func (m *WorldStatistics) String() string            { return proto.CompactTextString(m) }
func (*WorldStatistics) ProtoMessage()               {}
func (*WorldStatistics) Descriptor() ([]byte, []int) { return fileDescriptor125, []int{0} }

func (m *WorldStatistics) GetSimTime() *Time {
	if m != nil {
		return m.SimTime
	}
	return nil
}

func (m *WorldStatistics) GetPauseTime() *Time {
	if m != nil {
		return m.PauseTime
	}
	return nil
}

func (m *WorldStatistics) GetRealTime() *Time {
	if m != nil {
		return m.RealTime
	}
	return nil
}

func (m *WorldStatistics) GetPaused() bool {
	if m != nil && m.Paused != nil {
		return *m.Paused
	}
	return false
}

func (m *WorldStatistics) GetIterations() uint64 {
	if m != nil && m.Iterations != nil {
		return *m.Iterations
	}
	return 0
}

func (m *WorldStatistics) GetModelCount() int32 {
	if m != nil && m.ModelCount != nil {
		return *m.ModelCount
	}
	return 0
}

func (m *WorldStatistics) GetLogPlaybackStats() *LogPlaybackStatistics {
	if m != nil {
		return m.LogPlaybackStats
	}
	return nil
}

func init() {
	proto.RegisterType((*WorldStatistics)(nil), "gazebo.msgs.WorldStatistics")
}

func init() { proto.RegisterFile("world_stats.proto", fileDescriptor125) }

var fileDescriptor125 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0xcf, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0x5a, 0xdb, 0xba, 0x4e, 0x40, 0xe9, 0x78, 0x09, 0x3d, 0x95, 0xaa, 0xd0, 0x53,
	0x0e, 0x7d, 0x00, 0x9f, 0xc0, 0x83, 0xa0, 0xe0, 0x31, 0xa4, 0xbb, 0x21, 0x04, 0x93, 0xce, 0x92,
	0xc9, 0x22, 0xfa, 0xec, 0x1e, 0xcc, 0xba, 0x82, 0x2b, 0xb2, 0xb7, 0x90, 0xff, 0xe3, 0x4f, 0x7e,
	0x58, 0xbf, 0x51, 0x0a, 0x8d, 0xe6, 0x6c, 0x32, 0xab, 0x36, 0x51, 0x26, 0x14, 0xce, 0x7c, 0xd8,
	0x23, 0xa9, 0xc8, 0x8e, 0x37, 0x32, 0x90, 0xd3, 0x6d, 0x30, 0xef, 0x47, 0x53, 0xbf, 0x8e, 0xd9,
	0x06, 0xb2, 0x8f, 0x76, 0x38, 0xef, 0x3e, 0x67, 0x70, 0xf5, 0xd2, 0x17, 0x3d, 0x15, 0xe0, 0x39,
	0xfb, 0x9a, 0xf1, 0x06, 0x2a, 0xf6, 0x51, 0xf7, 0x4a, 0xce, 0xb7, 0xf3, 0xbd, 0x38, 0xac, 0xd5,
	0xa8, 0x59, 0x3d, 0x97, 0x00, 0xef, 0x00, 0x5a, 0xd3, 0xb1, 0x1d, 0xd8, 0xd9, 0x14, 0xbb, 0x85,
	0x8b, 0x64, 0x4d, 0x18, 0xd4, 0x62, 0x4a, 0x5d, 0xc2, 0xea, 0xbb, 0xac, 0x91, 0xcb, 0x42, 0x2a,
	0x44, 0x00, 0x9f, 0x6d, 0x2a, 0x3f, 0xa2, 0x13, 0xcb, 0x55, 0xb9, 0x5b, 0xe0, 0x35, 0x88, 0x48,
	0x8d, 0x0d, 0xba, 0xa6, 0xee, 0x94, 0xe5, 0xf9, 0x76, 0xb6, 0x5f, 0xe2, 0x3d, 0xe0, 0xff, 0x99,
	0xb2, 0x2a, 0x99, 0x38, 0xec, 0xfe, 0xbc, 0xf3, 0x40, 0xee, 0xf1, 0x47, 0xfd, 0x4e, 0xfd, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xea, 0xc8, 0x32, 0xd8, 0x45, 0x01, 0x00, 0x00,
}
