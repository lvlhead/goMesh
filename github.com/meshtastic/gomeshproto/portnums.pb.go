// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: meshtastic/portnums.proto

package gomeshproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// For any new 'apps' that run on the device or via sister apps on phones/PCs they should pick and use a
// unique 'portnum' for their application.
// If you are making a new app using meshtastic, please send in a pull request to add your 'portnum' to this
// master table.
// PortNums should be assigned in the following range:
// 0-63   Core Meshtastic use, do not use for third party apps
// 64-127 Registered 3rd party apps, send in a pull request that adds a new entry to portnums.proto to  register your application
// 256-511 Use one of these portnums for your private applications that you don't want to register publically
// All other values are reserved.
// Note: This was formerly a Type enum named 'typ' with the same id #
// We have change to this 'portnum' based scheme for specifying app handlers for particular payloads.
// This change is backwards compatible by treating the legacy OPAQUE/CLEAR_TEXT values identically.
type PortNum int32

const (
	// Deprecated: do not use in new code (formerly called OPAQUE)
	// A message sent from a device outside of the mesh, in a form the mesh does not understand
	// NOTE: This must be 0, because it is documented in IMeshService.aidl to be so
	// ENCODING: binary undefined
	PortNum_UNKNOWN_APP PortNum = 0
	// A simple UTF-8 text message, which even the little micros in the mesh
	// can understand and show on their screen eventually in some circumstances
	// even signal might send messages in this form (see below)
	// ENCODING: UTF-8 Plaintext (?)
	PortNum_TEXT_MESSAGE_APP PortNum = 1
	// Reserved for built-in GPIO/example app.
	// See remote_hardware.proto/HardwareMessage for details on the message sent/received to this port number
	// ENCODING: Protobuf
	PortNum_REMOTE_HARDWARE_APP PortNum = 2
	// The built-in position messaging app.
	// Payload is a Position message.
	// ENCODING: Protobuf
	PortNum_POSITION_APP PortNum = 3
	// The built-in user info app.
	// Payload is a User message.
	// ENCODING: Protobuf
	PortNum_NODEINFO_APP PortNum = 4
	// Protocol control packets for mesh protocol use.
	// Payload is a Routing message.
	// ENCODING: Protobuf
	PortNum_ROUTING_APP PortNum = 5
	// Admin control packets.
	// Payload is a AdminMessage message.
	// ENCODING: Protobuf
	PortNum_ADMIN_APP PortNum = 6
	// Compressed TEXT_MESSAGE payloads.
	// ENCODING: UTF-8 Plaintext (?) with Unishox2 Compression
	// NOTE: The Device Firmware converts a TEXT_MESSAGE_APP to TEXT_MESSAGE_COMPRESSED_APP if the compressed
	// payload is shorter. There's no need for app developers to do this themselves. Also the firmware will decompress
	// any incoming TEXT_MESSAGE_COMPRESSED_APP payload and convert to TEXT_MESSAGE_APP.
	PortNum_TEXT_MESSAGE_COMPRESSED_APP PortNum = 7
	// Waypoint payloads.
	// Payload is a Waypoint message.
	// ENCODING: Protobuf
	PortNum_WAYPOINT_APP PortNum = 8
	// Audio Payloads.
	// Encapsulated codec2 packets. On 2.4 GHZ Bandwidths only for now
	// ENCODING: codec2 audio frames
	// NOTE: audio frames contain a 3 byte header (0xc0 0xde 0xc2) and a one byte marker for the decompressed bitrate.
	// This marker comes from the 'moduleConfig.audio.bitrate' enum minus one.
	PortNum_AUDIO_APP PortNum = 9
	// Same as Text Message but originating from Detection Sensor Module.
	// NOTE: This portnum traffic is not sent to the public MQTT starting at firmware version 2.2.9
	PortNum_DETECTION_SENSOR_APP PortNum = 10
	// Provides a 'ping' service that replies to any packet it receives.
	// Also serves as a small example module.
	// ENCODING: ASCII Plaintext
	PortNum_REPLY_APP PortNum = 32
	// Used for the python IP tunnel feature
	// ENCODING: IP Packet. Handled by the python API, firmware ignores this one and pases on.
	PortNum_IP_TUNNEL_APP PortNum = 33
	// Paxcounter lib included in the firmware
	// ENCODING: protobuf
	PortNum_PAXCOUNTER_APP PortNum = 34
	// Provides a hardware serial interface to send and receive from the Meshtastic network.
	// Connect to the RX/TX pins of a device with 38400 8N1. Packets received from the Meshtastic
	// network is forwarded to the RX pin while sending a packet to TX will go out to the Mesh network.
	// Maximum packet size of 240 bytes.
	// Module is disabled by default can be turned on by setting SERIAL_MODULE_ENABLED = 1 in SerialPlugh.cpp.
	// ENCODING: binary undefined
	PortNum_SERIAL_APP PortNum = 64
	// STORE_FORWARD_APP (Work in Progress)
	// Maintained by Jm Casler (MC Hamster) : jm@casler.org
	// ENCODING: Protobuf
	PortNum_STORE_FORWARD_APP PortNum = 65
	// Optional port for messages for the range test module.
	// ENCODING: ASCII Plaintext
	// NOTE: This portnum traffic is not sent to the public MQTT starting at firmware version 2.2.9
	PortNum_RANGE_TEST_APP PortNum = 66
	// Provides a format to send and receive telemetry data from the Meshtastic network.
	// Maintained by Charles Crossan (crossan007) : crossan007@gmail.com
	// ENCODING: Protobuf
	PortNum_TELEMETRY_APP PortNum = 67
	// Experimental tools for estimating node position without a GPS
	// Maintained by Github user a-f-G-U-C (a Meshtastic contributor)
	// Project files at https://github.com/a-f-G-U-C/Meshtastic-ZPS
	// ENCODING: arrays of int64 fields
	PortNum_ZPS_APP PortNum = 68
	// Used to let multiple instances of Linux native applications communicate
	// as if they did using their LoRa chip.
	// Maintained by GitHub user GUVWAF.
	// Project files at https://github.com/GUVWAF/Meshtasticator
	// ENCODING: Protobuf (?)
	PortNum_SIMULATOR_APP PortNum = 69
	// Provides a traceroute functionality to show the route a packet towards
	// a certain destination would take on the mesh. Contains a RouteDiscovery message as payload.
	// ENCODING: Protobuf
	PortNum_TRACEROUTE_APP PortNum = 70
	// Aggregates edge info for the network by sending out a list of each node's neighbors
	// ENCODING: Protobuf
	PortNum_NEIGHBORINFO_APP PortNum = 71
	// ATAK Plugin
	// Portnum for payloads from the official Meshtastic ATAK plugin
	PortNum_ATAK_PLUGIN PortNum = 72
	// Provides unencrypted information about a node for consumption by a map via MQTT
	PortNum_MAP_REPORT_APP PortNum = 73
	// PowerStress based monitoring support (for automated power consumption testing)
	PortNum_POWERSTRESS_APP PortNum = 74
	// Private applications should use portnums >= 256.
	// To simplify initial development and testing you can use "PRIVATE_APP"
	// in your code without needing to rebuild protobuf files (via [regen-protos.sh](https://github.com/meshtastic/firmware/blob/master/bin/regen-protos.sh))
	PortNum_PRIVATE_APP PortNum = 256
	// ATAK Forwarder Module https://github.com/paulmandal/atak-forwarder
	// ENCODING: libcotshrink
	PortNum_ATAK_FORWARDER PortNum = 257
	// Currently we limit port nums to no higher than this value
	PortNum_MAX PortNum = 511
)

// Enum value maps for PortNum.
var (
	PortNum_name = map[int32]string{
		0:   "UNKNOWN_APP",
		1:   "TEXT_MESSAGE_APP",
		2:   "REMOTE_HARDWARE_APP",
		3:   "POSITION_APP",
		4:   "NODEINFO_APP",
		5:   "ROUTING_APP",
		6:   "ADMIN_APP",
		7:   "TEXT_MESSAGE_COMPRESSED_APP",
		8:   "WAYPOINT_APP",
		9:   "AUDIO_APP",
		10:  "DETECTION_SENSOR_APP",
		32:  "REPLY_APP",
		33:  "IP_TUNNEL_APP",
		34:  "PAXCOUNTER_APP",
		64:  "SERIAL_APP",
		65:  "STORE_FORWARD_APP",
		66:  "RANGE_TEST_APP",
		67:  "TELEMETRY_APP",
		68:  "ZPS_APP",
		69:  "SIMULATOR_APP",
		70:  "TRACEROUTE_APP",
		71:  "NEIGHBORINFO_APP",
		72:  "ATAK_PLUGIN",
		73:  "MAP_REPORT_APP",
		74:  "POWERSTRESS_APP",
		256: "PRIVATE_APP",
		257: "ATAK_FORWARDER",
		511: "MAX",
	}
	PortNum_value = map[string]int32{
		"UNKNOWN_APP":                 0,
		"TEXT_MESSAGE_APP":            1,
		"REMOTE_HARDWARE_APP":         2,
		"POSITION_APP":                3,
		"NODEINFO_APP":                4,
		"ROUTING_APP":                 5,
		"ADMIN_APP":                   6,
		"TEXT_MESSAGE_COMPRESSED_APP": 7,
		"WAYPOINT_APP":                8,
		"AUDIO_APP":                   9,
		"DETECTION_SENSOR_APP":        10,
		"REPLY_APP":                   32,
		"IP_TUNNEL_APP":               33,
		"PAXCOUNTER_APP":              34,
		"SERIAL_APP":                  64,
		"STORE_FORWARD_APP":           65,
		"RANGE_TEST_APP":              66,
		"TELEMETRY_APP":               67,
		"ZPS_APP":                     68,
		"SIMULATOR_APP":               69,
		"TRACEROUTE_APP":              70,
		"NEIGHBORINFO_APP":            71,
		"ATAK_PLUGIN":                 72,
		"MAP_REPORT_APP":              73,
		"POWERSTRESS_APP":             74,
		"PRIVATE_APP":                 256,
		"ATAK_FORWARDER":              257,
		"MAX":                         511,
	}
)

func (x PortNum) Enum() *PortNum {
	p := new(PortNum)
	*p = x
	return p
}

func (x PortNum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortNum) Descriptor() protoreflect.EnumDescriptor {
	return file_meshtastic_portnums_proto_enumTypes[0].Descriptor()
}

func (PortNum) Type() protoreflect.EnumType {
	return &file_meshtastic_portnums_proto_enumTypes[0]
}

func (x PortNum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortNum.Descriptor instead.
func (PortNum) EnumDescriptor() ([]byte, []int) {
	return file_meshtastic_portnums_proto_rawDescGZIP(), []int{0}
}

var File_meshtastic_portnums_proto protoreflect.FileDescriptor

var file_meshtastic_portnums_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2a, 0xa2, 0x04, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x41,
	0x50, 0x50, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x41, 0x50,
	0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x50, 0x50, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x44, 0x45, 0x49, 0x4e, 0x46,
	0x4f, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x45, 0x58, 0x54, 0x5f,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53,
	0x45, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x41, 0x59, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x54,
	0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x41, 0x50,
	0x50, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x41, 0x50, 0x50,
	0x10, 0x20, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x50, 0x5f, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x41, 0x50, 0x50, 0x10, 0x21, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x58, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x22, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52,
	0x49, 0x41, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x40, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x41,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x41,
	0x50, 0x50, 0x10, 0x42, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x54, 0x52,
	0x59, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x43, 0x12, 0x0b, 0x0a, 0x07, 0x5a, 0x50, 0x53, 0x5f, 0x41,
	0x50, 0x50, 0x10, 0x44, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x49, 0x4d, 0x55, 0x4c, 0x41, 0x54, 0x4f,
	0x52, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x45, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x43, 0x45,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x46, 0x12, 0x14, 0x0a, 0x10, 0x4e,
	0x45, 0x49, 0x47, 0x48, 0x42, 0x4f, 0x52, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x41, 0x50, 0x50, 0x10,
	0x47, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x54, 0x41, 0x4b, 0x5f, 0x50, 0x4c, 0x55, 0x47, 0x49, 0x4e,
	0x10, 0x48, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x50, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x49, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4f, 0x57, 0x45, 0x52, 0x53,
	0x54, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x4a, 0x12, 0x10, 0x0a, 0x0b, 0x50,
	0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x80, 0x02, 0x12, 0x13, 0x0a,
	0x0e, 0x41, 0x54, 0x41, 0x4b, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x45, 0x52, 0x10,
	0x81, 0x02, 0x12, 0x08, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10, 0xff, 0x03, 0x42, 0x5d, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x42, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x6e, 0x75, 0x6d, 0x73, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0xaa, 0x02, 0x14, 0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_portnums_proto_rawDescOnce sync.Once
	file_meshtastic_portnums_proto_rawDescData = file_meshtastic_portnums_proto_rawDesc
)

func file_meshtastic_portnums_proto_rawDescGZIP() []byte {
	file_meshtastic_portnums_proto_rawDescOnce.Do(func() {
		file_meshtastic_portnums_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_portnums_proto_rawDescData)
	})
	return file_meshtastic_portnums_proto_rawDescData
}

var file_meshtastic_portnums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meshtastic_portnums_proto_goTypes = []interface{}{
	(PortNum)(0), // 0: meshtastic.PortNum
}
var file_meshtastic_portnums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meshtastic_portnums_proto_init() }
func file_meshtastic_portnums_proto_init() {
	if File_meshtastic_portnums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meshtastic_portnums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_portnums_proto_goTypes,
		DependencyIndexes: file_meshtastic_portnums_proto_depIdxs,
		EnumInfos:         file_meshtastic_portnums_proto_enumTypes,
	}.Build()
	File_meshtastic_portnums_proto = out.File
	file_meshtastic_portnums_proto_rawDesc = nil
	file_meshtastic_portnums_proto_goTypes = nil
	file_meshtastic_portnums_proto_depIdxs = nil
}
