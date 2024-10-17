// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: meshtastic/remote_hardware.proto

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

// TODO: REPLACE
type HardwareMessage_Type int32

const (
	// Unset/unused
	HardwareMessage_UNSET HardwareMessage_Type = 0
	// Set gpio gpios based on gpio_mask/gpio_value
	HardwareMessage_WRITE_GPIOS HardwareMessage_Type = 1
	// We are now interested in watching the gpio_mask gpios.
	// If the selected gpios change, please broadcast GPIOS_CHANGED.
	// Will implicitly change the gpios requested to be INPUT gpios.
	HardwareMessage_WATCH_GPIOS HardwareMessage_Type = 2
	// The gpios listed in gpio_mask have changed, the new values are listed in gpio_value
	HardwareMessage_GPIOS_CHANGED HardwareMessage_Type = 3
	// Read the gpios specified in gpio_mask, send back a READ_GPIOS_REPLY reply with gpio_value populated
	HardwareMessage_READ_GPIOS HardwareMessage_Type = 4
	// A reply to READ_GPIOS. gpio_mask and gpio_value will be populated
	HardwareMessage_READ_GPIOS_REPLY HardwareMessage_Type = 5
)

// Enum value maps for HardwareMessage_Type.
var (
	HardwareMessage_Type_name = map[int32]string{
		0: "UNSET",
		1: "WRITE_GPIOS",
		2: "WATCH_GPIOS",
		3: "GPIOS_CHANGED",
		4: "READ_GPIOS",
		5: "READ_GPIOS_REPLY",
	}
	HardwareMessage_Type_value = map[string]int32{
		"UNSET":            0,
		"WRITE_GPIOS":      1,
		"WATCH_GPIOS":      2,
		"GPIOS_CHANGED":    3,
		"READ_GPIOS":       4,
		"READ_GPIOS_REPLY": 5,
	}
)

func (x HardwareMessage_Type) Enum() *HardwareMessage_Type {
	p := new(HardwareMessage_Type)
	*p = x
	return p
}

func (x HardwareMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HardwareMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_meshtastic_remote_hardware_proto_enumTypes[0].Descriptor()
}

func (HardwareMessage_Type) Type() protoreflect.EnumType {
	return &file_meshtastic_remote_hardware_proto_enumTypes[0]
}

func (x HardwareMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HardwareMessage_Type.Descriptor instead.
func (HardwareMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_meshtastic_remote_hardware_proto_rawDescGZIP(), []int{0, 0}
}

// An example app to show off the module system. This message is used for
// REMOTE_HARDWARE_APP PortNums.
// Also provides easy remote access to any GPIO.
// In the future other remote hardware operations can be added based on user interest
// (i.e. serial output, spi/i2c input/output).
// FIXME - currently this feature is turned on by default which is dangerous
// because no security yet (beyond the channel mechanism).
// It should be off by default and then protected based on some TBD mechanism
// (a special channel once multichannel support is included?)
type HardwareMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What type of HardwareMessage is this?
	Type HardwareMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=meshtastic.HardwareMessage_Type" json:"type,omitempty"`
	// What gpios are we changing. Not used for all MessageTypes, see MessageType for details
	GpioMask uint64 `protobuf:"varint,2,opt,name=gpio_mask,json=gpioMask,proto3" json:"gpio_mask,omitempty"`
	// For gpios that were listed in gpio_mask as valid, what are the signal levels for those gpios.
	// Not used for all MessageTypes, see MessageType for details
	GpioValue uint64 `protobuf:"varint,3,opt,name=gpio_value,json=gpioValue,proto3" json:"gpio_value,omitempty"`
}

func (x *HardwareMessage) Reset() {
	*x = HardwareMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_remote_hardware_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareMessage) ProtoMessage() {}

func (x *HardwareMessage) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_remote_hardware_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareMessage.ProtoReflect.Descriptor instead.
func (*HardwareMessage) Descriptor() ([]byte, []int) {
	return file_meshtastic_remote_hardware_proto_rawDescGZIP(), []int{0}
}

func (x *HardwareMessage) GetType() HardwareMessage_Type {
	if x != nil {
		return x.Type
	}
	return HardwareMessage_UNSET
}

func (x *HardwareMessage) GetGpioMask() uint64 {
	if x != nil {
		return x.GpioMask
	}
	return 0
}

func (x *HardwareMessage) GetGpioValue() uint64 {
	if x != nil {
		return x.GpioValue
	}
	return 0
}

var File_meshtastic_remote_hardware_proto protoreflect.FileDescriptor

var file_meshtastic_remote_hardware_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x22, 0xf1,
	0x01, 0x0a, 0x0f, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x48, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x70, 0x69, 0x6f,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x70, 0x69,
	0x6f, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x70, 0x69, 0x6f, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x67, 0x70, 0x69, 0x6f, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x6c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x52, 0x49, 0x54, 0x45,
	0x5f, 0x47, 0x50, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x47, 0x50, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x50, 0x49,
	0x4f, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a,
	0x52, 0x45, 0x41, 0x44, 0x5f, 0x47, 0x50, 0x49, 0x4f, 0x53, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x52, 0x45, 0x41, 0x44, 0x5f, 0x47, 0x50, 0x49, 0x4f, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x59,
	0x10, 0x05, 0x42, 0x63, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76,
	0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0xaa, 0x02, 0x14,
	0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_remote_hardware_proto_rawDescOnce sync.Once
	file_meshtastic_remote_hardware_proto_rawDescData = file_meshtastic_remote_hardware_proto_rawDesc
)

func file_meshtastic_remote_hardware_proto_rawDescGZIP() []byte {
	file_meshtastic_remote_hardware_proto_rawDescOnce.Do(func() {
		file_meshtastic_remote_hardware_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_remote_hardware_proto_rawDescData)
	})
	return file_meshtastic_remote_hardware_proto_rawDescData
}

var file_meshtastic_remote_hardware_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meshtastic_remote_hardware_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meshtastic_remote_hardware_proto_goTypes = []interface{}{
	(HardwareMessage_Type)(0), // 0: meshtastic.HardwareMessage.Type
	(*HardwareMessage)(nil),   // 1: meshtastic.HardwareMessage
}
var file_meshtastic_remote_hardware_proto_depIdxs = []int32{
	0, // 0: meshtastic.HardwareMessage.type:type_name -> meshtastic.HardwareMessage.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meshtastic_remote_hardware_proto_init() }
func file_meshtastic_remote_hardware_proto_init() {
	if File_meshtastic_remote_hardware_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_remote_hardware_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meshtastic_remote_hardware_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_remote_hardware_proto_goTypes,
		DependencyIndexes: file_meshtastic_remote_hardware_proto_depIdxs,
		EnumInfos:         file_meshtastic_remote_hardware_proto_enumTypes,
		MessageInfos:      file_meshtastic_remote_hardware_proto_msgTypes,
	}.Build()
	File_meshtastic_remote_hardware_proto = out.File
	file_meshtastic_remote_hardware_proto_rawDesc = nil
	file_meshtastic_remote_hardware_proto_goTypes = nil
	file_meshtastic_remote_hardware_proto_depIdxs = nil
}
