// Code generated by protoc-gen-go.
// source: altimeter.proto
// DO NOT EDIT!

/*
Package gazebo_msgs is a generated protocol buffer package.

It is generated from these files:
	altimeter.proto
	any.proto
	atmosphere.proto
	axis.proto
	battery.proto
	boxgeom.proto
	camera_cmd.proto
	camera_lens.proto
	camerasensor.proto
	cessna.proto
	collision.proto
	color.proto
	contact.proto
	contactsensor.proto
	contacts.proto
	cylindergeom.proto
	density.proto
	diagnostics.proto
	distortion.proto
	empty.proto
	factory.proto
	fluid.proto
	fog.proto
	friction.proto
	geometry.proto
	gps.proto
	gps_sensor.proto
	gui_camera.proto
	gui.proto
	gz_string.proto
	gz_string_v.proto
	header.proto
	heightmapgeom.proto
	hydra.proto
	imagegeom.proto
	image.proto
	images_stamped.proto
	image_stamped.proto
	imu.proto
	imu_sensor.proto
	inertial.proto
	int.proto
	joint_animation.proto
	joint_cmd.proto
	joint.proto
	joint_wrench.proto
	joint_wrench_stamped.proto
	joystick.proto
	laserscan.proto
	laserscan_stamped.proto
	light.proto
	link_data.proto
	link.proto
	log_control.proto
	logical_camera_image.proto
	logical_camera_sensor.proto
	log_playback_control.proto
	log_playback_stats.proto
	log_status.proto
	magnetometer.proto
	material.proto
	meshgeom.proto
	model_configuration.proto
	model.proto
	model_v.proto
	packet.proto
	param.proto
	param_v.proto
	physics.proto
	pid.proto
	planegeom.proto
	plugin.proto
	pointcloud.proto
	polylinegeom.proto
	pose_animation.proto
	pose.proto
	poses_stamped.proto
	pose_stamped.proto
	pose_trajectory.proto
	pose_v.proto
	projector.proto
	propagation_grid.proto
	propagation_particle.proto
	publishers.proto
	publish.proto
	quaternion.proto
	raysensor.proto
	request.proto
	response.proto
	rest_login.proto
	rest_logout.proto
	rest_post.proto
	rest_response.proto
	road.proto
	scene.proto
	selection.proto
	sensor_noise.proto
	sensor.proto
	server_control.proto
	shadows.proto
	sim_event.proto
	sky.proto
	sonar.proto
	sonar_stamped.proto
	spheregeom.proto
	spherical_coordinates.proto
	subscribe.proto
	surface.proto
	tactile.proto
	test.proto
	time.proto
	topic_info.proto
	track_visual.proto
	undo_redo.proto
	user_cmd.proto
	user_cmd_stats.proto
	vector2d.proto
	vector3d.proto
	visual.proto
	wind.proto
	wireless_node.proto
	wireless_nodes.proto
	world_control.proto
	world_modify.proto
	world_reset.proto
	world_stats.proto
	wrench.proto
	wrench_stamped.proto

It has these top-level messages:
	Altimeter
	Any
	Atmosphere
	Axis
	Battery
	BoxGeom
	CameraCmd
	CameraLens
	CameraSensor
	Cessna
	Collision
	Color
	Contact
	ContactSensor
	Contacts
	CylinderGeom
	Density
	Diagnostics
	Distortion
	Empty
	Factory
	Fluid
	Fog
	Friction
	Geometry
	GPS
	GPSSensor
	GUICamera
	GUI
	GzString
	GzString_V
	Header
	HeightmapGeom
	Hydra
	ImageGeom
	Image
	ImagesStamped
	ImageStamped
	IMU
	IMUSensor
	Inertial
	Int
	JointAnimation
	JointCmd
	Joint
	JointWrench
	ForceTorque
	Joystick
	LaserScan
	LaserScanStamped
	Light
	LinkData
	Link
	LogControl
	LogicalCameraImage
	LogicalCameraSensor
	LogPlaybackControl
	LogPlaybackStatistics
	LogStatus
	Magnetometer
	Material
	MeshGeom
	ModelConfiguration
	Model
	Model_V
	Packet
	Param
	Param_V
	Physics
	PID
	PlaneGeom
	Plugin
	PointCloud
	Polyline
	PoseAnimation
	Pose
	PosesStamped
	PoseStamped
	PoseTrajectory
	Pose_V
	Projector
	PropagationGrid
	PropagationParticle
	Publishers
	Publish
	Quaternion
	RaySensor
	Request
	Response
	RestLogin
	RestLogout
	RestPost
	RestResponse
	Road
	Scene
	Selection
	SensorNoise
	Sensor
	ServerControl
	Shadows
	SimEvent
	Sky
	Sonar
	SonarStamped
	SphereGeom
	SphericalCoordinates
	Subscribe
	Surface
	Tactile
	Test
	Time
	TopicInfo
	TrackVisual
	UndoRedo
	UserCmd
	UserCmdStats
	Vector2D
	Vector3D
	Visual
	Wind
	WirelessNode
	WirelessNodes
	WorldControl
	WorldModify
	WorldReset
	WorldStatistics
	Wrench
	WrenchStamped
*/
package gazebo_msgs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / \brief Altimeter sensor data
type Altimeter struct {
	// / \brief Timestamp of the altimeter data
	Time *Time `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	// / \brief Vertical position data, in meters.
	VerticalPosition *float64 `protobuf:"fixed64,2,req,name=vertical_position" json:"vertical_position,omitempty"`
	// / \brief Vertical velocity data, in meters/second.
	VerticalVelocity *float64 `protobuf:"fixed64,3,req,name=vertical_velocity" json:"vertical_velocity,omitempty"`
	// / \brief Vertical reference.
	VerticalReference *float64 `protobuf:"fixed64,4,req,name=vertical_reference" json:"vertical_reference,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *Altimeter) Reset()                    { *m = Altimeter{} }
func (m *Altimeter) String() string            { return proto.CompactTextString(m) }
func (*Altimeter) ProtoMessage()               {}
func (*Altimeter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Altimeter) GetTime() *Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Altimeter) GetVerticalPosition() float64 {
	if m != nil && m.VerticalPosition != nil {
		return *m.VerticalPosition
	}
	return 0
}

func (m *Altimeter) GetVerticalVelocity() float64 {
	if m != nil && m.VerticalVelocity != nil {
		return *m.VerticalVelocity
	}
	return 0
}

func (m *Altimeter) GetVerticalReference() float64 {
	if m != nil && m.VerticalReference != nil {
		return *m.VerticalReference
	}
	return 0
}

func init() {
	proto.RegisterType((*Altimeter)(nil), "gazebo.msgs.Altimeter")
}

func init() { proto.RegisterFile("altimeter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcc, 0x29, 0xc9,
	0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4f, 0xac,
	0x4a, 0x4d, 0xca, 0xd7, 0xcb, 0x2d, 0x4e, 0x2f, 0x96, 0xe2, 0x02, 0xc9, 0x41, 0x24, 0x94, 0xea,
	0xb8, 0x38, 0x1d, 0x61, 0x6a, 0x85, 0xe4, 0xb9, 0x58, 0x40, 0x4c, 0x09, 0x46, 0x05, 0x26, 0x0d,
	0x6e, 0x23, 0x41, 0x3d, 0x24, 0x4d, 0x7a, 0x21, 0x40, 0x09, 0x21, 0x49, 0x2e, 0xc1, 0xb2, 0xd4,
	0xa2, 0x92, 0xcc, 0xe4, 0xc4, 0x9c, 0xf8, 0x82, 0xfc, 0xe2, 0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x09,
	0x26, 0xa0, 0x6a, 0x46, 0x14, 0xa9, 0xb2, 0xd4, 0x9c, 0xfc, 0xe4, 0xcc, 0x92, 0x4a, 0x09, 0x66,
	0xb0, 0x94, 0x14, 0x97, 0x10, 0x5c, 0xaa, 0x28, 0x35, 0x2d, 0xb5, 0x28, 0x35, 0x2f, 0x39, 0x55,
	0x82, 0x05, 0x24, 0x07, 0x08, 0x00, 0x00, 0xff, 0xff, 0x46, 0x3f, 0xc4, 0x2b, 0xaa, 0x00, 0x00,
	0x00,
}
