// Code generated by protoc-gen-go.
// source: sim_event.proto
// DO NOT EDIT!

package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SimEvent struct {
	// / \brief ID of this event message
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// / \brief Type of sim event
	Type *string `protobuf:"bytes,2,req,name=type" json:"type,omitempty"`
	// / \brief Name of sim event
	Name *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	// / \brief Statistics of the world
	WorldStatistics *WorldStatistics `protobuf:"bytes,4,req,name=world_statistics" json:"world_statistics,omitempty"`
	// / \brief Data describing the sim event
	Data             *string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SimEvent) Reset()                    { *m = SimEvent{} }
func (m *SimEvent) String() string            { return proto.CompactTextString(m) }
func (*SimEvent) ProtoMessage()               {}
func (*SimEvent) Descriptor() ([]byte, []int) { return fileDescriptor100, []int{0} }

func (m *SimEvent) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *SimEvent) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *SimEvent) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SimEvent) GetWorldStatistics() *WorldStatistics {
	if m != nil {
		return m.WorldStatistics
	}
	return nil
}

func (m *SimEvent) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*SimEvent)(nil), "gazebo.msgs.SimEvent")
}

func init() { proto.RegisterFile("sim_event.proto", fileDescriptor100) }

var fileDescriptor100 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0xcc, 0x8d,
	0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac,
	0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0x12, 0x2c, 0xcf, 0x2f, 0xca, 0x49, 0x89,
	0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x86, 0xc8, 0x2b, 0x55, 0x70, 0x71, 0x04, 0x67, 0xe6, 0xba, 0x82,
	0x74, 0x08, 0x71, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x0a, 0xf1, 0x70,
	0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x69, 0x70, 0x82, 0x78, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0xcc, 0x60, 0x9e, 0x19, 0x97, 0x00, 0xc2, 0xa0, 0xcc, 0xe2, 0x92, 0xcc, 0xe4, 0x62,
	0x09, 0x16, 0xa0, 0x0c, 0xb7, 0x91, 0x8c, 0x1e, 0x92, 0x75, 0x7a, 0xe1, 0x20, 0x45, 0xc1, 0x70,
	0x35, 0x20, 0x53, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x58, 0x81, 0x36, 0x70, 0x02, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x8b, 0xd6, 0x83, 0x1b, 0xab, 0x00, 0x00, 0x00,
}