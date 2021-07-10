// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: proto/serialreader.proto

package serialreader_server

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetTimeSeriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTimeSeriesData) Reset() {
	*x = GetTimeSeriesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_serialreader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTimeSeriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTimeSeriesData) ProtoMessage() {}

func (x *GetTimeSeriesData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_serialreader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTimeSeriesData.ProtoReflect.Descriptor instead.
func (*GetTimeSeriesData) Descriptor() ([]byte, []int) {
	return file_proto_serialreader_proto_rawDescGZIP(), []int{0}
}

type SparkFunWeatherShieldTimeSeriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status                 bool                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	HumidityValue          float64              `protobuf:"fixed32,3,opt,name=humidityValue,proto3" json:"humidityValue,omitempty"`
	HumidityUnit           string               `protobuf:"bytes,4,opt,name=humidityUnit,proto3" json:"humidityUnit,omitempty"`
	TemperatureValue       float64              `protobuf:"fixed32,5,opt,name=temperatureValue,proto3" json:"temperatureValue,omitempty"`
	TemperatureUnit        string               `protobuf:"bytes,6,opt,name=temperatureUnit,proto3" json:"temperatureUnit,omitempty"`
	PressureValue          float64              `protobuf:"fixed32,7,opt,name=pressureValue,proto3" json:"pressureValue,omitempty"`
	PressureUnit           string               `protobuf:"bytes,8,opt,name=pressureUnit,proto3" json:"pressureUnit,omitempty"`
	TemperatureBackupValue float64              `protobuf:"fixed32,9,opt,name=temperatureBackupValue,proto3" json:"temperatureBackupValue,omitempty"`
	TemperatureBackupUnit  string               `protobuf:"bytes,10,opt,name=temperatureBackupUnit,proto3" json:"temperatureBackupUnit,omitempty"`
	AltitudeValue          float64              `protobuf:"fixed32,11,opt,name=altitudeValue,proto3" json:"altitudeValue,omitempty"`
	AltitudeUnit           string               `protobuf:"bytes,12,opt,name=altitudeUnit,proto3" json:"altitudeUnit,omitempty"`
	IlluminanceValue       float64              `protobuf:"fixed32,13,opt,name=illuminanceValue,proto3" json:"illuminanceValue,omitempty"`
	IlluminanceUnit        string               `protobuf:"bytes,14,opt,name=illuminanceUnit,proto3" json:"illuminanceUnit,omitempty"`
	SoilMoistureValue      float64              `protobuf:"fixed32,15,opt,name=soilMoistureValue,proto3" json:"soilMoistureValue,omitempty"`
	SoilMoistureUnit       string               `protobuf:"bytes,16,opt,name=soilMoistureUnit,proto3" json:"soilMoistureUnit,omitempty"`
}

func (x *SparkFunWeatherShieldTimeSeriesData) Reset() {
	*x = SparkFunWeatherShieldTimeSeriesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_serialreader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparkFunWeatherShieldTimeSeriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparkFunWeatherShieldTimeSeriesData) ProtoMessage() {}

func (x *SparkFunWeatherShieldTimeSeriesData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_serialreader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparkFunWeatherShieldTimeSeriesData.ProtoReflect.Descriptor instead.
func (*SparkFunWeatherShieldTimeSeriesData) Descriptor() ([]byte, []int) {
	return file_proto_serialreader_proto_rawDescGZIP(), []int{1}
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetHumidityValue() float64 {
	if x != nil {
		return x.HumidityValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetHumidityUnit() string {
	if x != nil {
		return x.HumidityUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetTemperatureValue() float64 {
	if x != nil {
		return x.TemperatureValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetTemperatureUnit() string {
	if x != nil {
		return x.TemperatureUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetPressureValue() float64 {
	if x != nil {
		return x.PressureValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetPressureUnit() string {
	if x != nil {
		return x.PressureUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetTemperatureBackupValue() float64 {
	if x != nil {
		return x.TemperatureBackupValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetTemperatureBackupUnit() string {
	if x != nil {
		return x.TemperatureBackupUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetAltitudeValue() float64 {
	if x != nil {
		return x.AltitudeValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetAltitudeUnit() string {
	if x != nil {
		return x.AltitudeUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetIlluminanceValue() float64 {
	if x != nil {
		return x.IlluminanceValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetIlluminanceUnit() string {
	if x != nil {
		return x.IlluminanceUnit
	}
	return ""
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetSoilMoistureValue() float64 {
	if x != nil {
		return x.SoilMoistureValue
	}
	return 0
}

func (x *SparkFunWeatherShieldTimeSeriesData) GetSoilMoistureUnit() string {
	if x != nil {
		return x.SoilMoistureUnit
	}
	return ""
}

var File_proto_serialreader_proto protoreflect.FileDescriptor

var file_proto_serialreader_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc9, 0x05, 0x0a, 0x23, 0x53, 0x70, 0x61, 0x72,
	0x6b, 0x46, 0x75, 0x6e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68,
	0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x16, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x69,
	0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x69, 0x6c, 0x6c, 0x75, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6c, 0x6c, 0x75, 0x6d, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x6f, 0x69,
	0x6c, 0x4d, 0x6f, 0x69, 0x73, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x73, 0x6f, 0x69, 0x6c, 0x4d, 0x6f, 0x69, 0x73, 0x74, 0x75,
	0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x6f, 0x69, 0x6c, 0x4d,
	0x6f, 0x69, 0x73, 0x74, 0x75, 0x72, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x6f, 0x69, 0x6c, 0x4d, 0x6f, 0x69, 0x73, 0x74, 0x75, 0x72, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x32, 0x76, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x46,
	0x75, 0x6e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x46, 0x75, 0x6e, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x74, 0x6d, 0x69,
	0x6b, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_serialreader_proto_rawDescOnce sync.Once
	file_proto_serialreader_proto_rawDescData = file_proto_serialreader_proto_rawDesc
)

func file_proto_serialreader_proto_rawDescGZIP() []byte {
	file_proto_serialreader_proto_rawDescOnce.Do(func() {
		file_proto_serialreader_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_serialreader_proto_rawDescData)
	})
	return file_proto_serialreader_proto_rawDescData
}

var file_proto_serialreader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_serialreader_proto_goTypes = []interface{}{
	(*GetTimeSeriesData)(nil),                   // 0: proto.GetTimeSeriesData
	(*SparkFunWeatherShieldTimeSeriesData)(nil), // 1: proto.SparkFunWeatherShieldTimeSeriesData
	(*timestamp.Timestamp)(nil),                 // 2: google.protobuf.Timestamp
}
var file_proto_serialreader_proto_depIdxs = []int32{
	2, // 0: proto.SparkFunWeatherShieldTimeSeriesData.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: proto.SerialReader.GetSparkFunWeatherShieldData:input_type -> proto.GetTimeSeriesData
	1, // 2: proto.SerialReader.GetSparkFunWeatherShieldData:output_type -> proto.SparkFunWeatherShieldTimeSeriesData
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_serialreader_proto_init() }
func file_proto_serialreader_proto_init() {
	if File_proto_serialreader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_serialreader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTimeSeriesData); i {
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
		file_proto_serialreader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparkFunWeatherShieldTimeSeriesData); i {
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
			RawDescriptor: file_proto_serialreader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_serialreader_proto_goTypes,
		DependencyIndexes: file_proto_serialreader_proto_depIdxs,
		MessageInfos:      file_proto_serialreader_proto_msgTypes,
	}.Build()
	File_proto_serialreader_proto = out.File
	file_proto_serialreader_proto_rawDesc = nil
	file_proto_serialreader_proto_goTypes = nil
	file_proto_serialreader_proto_depIdxs = nil
}
