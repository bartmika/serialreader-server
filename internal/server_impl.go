package internal

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	pb "github.com/bartmika/serialreader-server/proto"
)

type SerialReaderServerImpl struct {
	arduinoReader *ArduinoReader
	pb.SerialReaderServer
}

func (s *SerialReaderServerImpl) GetSparkFunWeatherShieldData(ctx context.Context, in *pb.GetTimeSeriesData) (*pb.SparkFunWeatherShieldTimeSeriesData, error) {
	datum := s.arduinoReader.GetSparkFunWeatherShieldData()
	return &pb.SparkFunWeatherShieldTimeSeriesData{
		Status:                 true,
		Timestamp:              ptypes.TimestampNow(), // Note: https://godoc.org/github.com/golang/protobuf/ptypes#Timestamp
		HumidityValue:          datum.HumidityValue,
		HumidityUnit:           datum.HumidityUnit,
		TemperatureValue:       datum.TemperatureValue,
		TemperatureUnit:        datum.TemperatureUnit,
		PressureValue:          datum.PressureValue,
		PressureUnit:           datum.PressureUnit,
		TemperatureBackupValue: datum.TemperatureBackupValue,
		TemperatureBackupUnit:  datum.TemperatureBackupUnit,
		AltitudeValue:          datum.AltitudeValue,
		AltitudeUnit:           datum.AltitudeUnit,
		IlluminanceValue:       datum.IlluminanceValue,
		IlluminanceUnit:        datum.IlluminanceUnit,
	}, nil
}
