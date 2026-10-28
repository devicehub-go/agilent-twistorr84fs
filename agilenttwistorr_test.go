package agilenttwistorr_test

import (
	"testing"
	"time"

	twistorr84fs "github.com/devicehub-go/agilent-twistorr84fs"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestReadings(t *testing.T) {
	sensor := twistorr84fs.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.3.35",
			Port:         4001,
			ReadTimeout:  1000 * time.Millisecond,
			WriteTimeout: 1000 * time.Millisecond,
		},
	}, 0)
	if err := sensor.Connect(); err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer sensor.Disconnect()

	// Simple integer readings
	intReadings := []struct {
		name string
		unit string
		read func() (int, error)
	}{
		{"PumpCurrent", "mA", sensor.GetPumpCurrent},
		{"PumpVoltage", "Vdc", sensor.GetPumpVoltage},
		{"PumpPower", "W", sensor.GetPumpPower},
		{"DrivingFrequency", "Hz", sensor.GetDrivingFrequency},
		{"PumpTemperature", "°C", sensor.GetPumpTemperature},
		{"ControllerHeatsinkTemperature", "°C", sensor.GetControllerHeatsinkTemperature},
		{"ControllerAirTemperature", "°C", sensor.GetControllerAirTemperature},
		{"PressureReadingFormat", "", sensor.GetPressureReadingFormat},
		{"RotationFrequency", "rpm", sensor.GetRotationFrequency},
		{"CycleTime", "min", sensor.GetCycleTime},
		{"CycleNumber", "", sensor.GetCycleNumber},
		{"PumpLife", "h", sensor.GetPumpLife},
	}

	for _, r := range intReadings {
		t.Run(r.name, func(t *testing.T) {
			value, err := r.read()
			if err != nil {
				t.Fatalf("%s: %v", r.name, err)
			}
			t.Logf("%s = %d %s", r.name, value, r.unit)
		})
	}

	t.Run("PumpStatus", func(t *testing.T) {
		status, desc, err := sensor.GetPumpStatus()
		if err != nil {
			t.Fatalf("PumpStatus: %v", err)
		}
		t.Logf("PumpStatus = %d (%s)", status, desc)
	})

	t.Run("ErrorCode", func(t *testing.T) {
		code, errs, err := sensor.GetErrorCode()
		if err != nil {
			t.Fatalf("ErrorCode: %v", err)
		}
		t.Logf("ErrorCode = %d, active errors: %v", code, errs)
	})

	t.Run("GaugeStatus", func(t *testing.T) {
		status, desc, err := sensor.GetGaugeStatus()
		if err != nil {
			t.Fatalf("GaugeStatus: %v", err)
		}
		t.Logf("GaugeStatus = %d (%s)", status, desc)
	})
}
