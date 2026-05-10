package protocol

var PumpStatusDesc = map[int]string{
	0: "Stop",
	1: "Waiting interlock",
	2: "Starting",
	3: "Auto-tuning",
	4: "Braking",
	5: "Normal",
	6: "Fail",
}

var ErrorCodeBitDesc = map[int]string{
	0: "No connection",
	1: "Pump overtemperature",
	2: "Controller overtemperature",
	3: "Power failure",
	4: "Auxiliary power failure",
	5: "Overvoltage",
	6: "Short circuit",
	7: "Too high load",
}

var GaugeStatusDesc = map[int]string{
	0: "No gauge connected",
	1: "Gauge connected",
	2: "Under range / gauge error",
	3: "Over range / gauge error",
	4: "Rid unknown",
}

// Gets the pump current in mA DC
func (t *TwissTorr84FS) GetPumpCurrent() (int, error) {
	return t.GetInt("200")
}

// Gets the pump voltage in Vdc
func (t *TwissTorr84FS) GetPumpVoltage() (int, error) {
	return t.GetInt("201")
}

// Gets the pump power in watts. This paramater is calculated
// throught: pump current * pump voltage duty cycle
func (t *TwissTorr84FS) GetPumpPower() (int, error) {
	return t.GetInt("202")
}

// Gets the driving frequency in Hz
func (t *TwissTorr84FS) GetDrivingFrequency() (int, error) {
	return t.GetInt("203")
}

// Gets the pump temperature in Celsius degrees
func (t *TwissTorr84FS) GetPumpTemperature() (int, error) {
	return t.GetInt("204")
}

// Gets the pump status
func (t *TwissTorr84FS) GetPumpStatus() (int, string, error) {
	status, err := t.GetInt("205")
	if err != nil {
		return 0, "", err
	}
	return status, PumpStatusDesc[status], nil
}

// Gets the error code
func (t *TwissTorr84FS) GetErrorCode() (int, []string, error) {
	code, err := t.GetInt("206")
	if err != nil {
		return 0, nil, err
	}

	var errors []string
	for i := range 8 {
		if code&(1<<i) != 0 {
			errors = append(errors, ErrorCodeBitDesc[i])
		}
	}
	return code, errors, nil
}

// Gets the controller heatsink temperature in Celsius degrees
func (t *TwissTorr84FS) GetControllerHeatsinkTemperature() (int, error) {
	return t.GetInt("211")
}

// Gets the controller air temperature in Celsius degrees
func (t *TwissTorr84FS) GetControllerAirTemperature() (int, error) {
	return t.GetInt("216")
}

// Gets pressure reading format
func (t *TwissTorr84FS) GetPressureReadingFormat() (int, error) {
	return t.GetInt("224")
}

// Gets the rotation frequency in rpm
func (t *TwissTorr84FS) GetRotationFrequency() (int, error) {
	return t.GetInt("226")
}

// Gets the gauge status
func (t *TwissTorr84FS) GetGaugeStatus() (int, string, error) {
	status, err := t.GetInt("257")
	if err != nil {
		return 0, "", err
	}
	return status, GaugeStatusDesc[status], nil
}

// Gets the cycle time in minutes (zeroed by the reset command)
func (t *TwissTorr84FS) GetCycleTime() (int, error) {
	return t.GetInt("300")
}

// Gets the cycle number (zeroed by the reset command)
func (t *TwissTorr84FS) GetCycleNumber() (int, error) {
	return t.GetInt("301")
}

// Gets the pump life in hours (zeroed by the reset command)
func (t *TwissTorr84FS) GetPumpLife() (int, error) {
	return t.GetInt("302")
}
