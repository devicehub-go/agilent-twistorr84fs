package protocol

import "fmt"

/*
Sets the low speed activation status.

Parameters:
	- value: 0 to disable, 1 to enable.
*/
func(t *TwissTorr84FS) SetLowSpeedActivation(value int) error {
	if(value < 0 || value > 1) {
		return fmt.Errorf("invalid value for soft start")
	}
	return t.SetInt("001", int(value))
}

/*
Sets the communication mode.

Parameters:
	- value: 0 for serial, 1 for remote.
*/
func(t *TwissTorr84FS) SetCommunicationMode(value int) error {
	return t.SetInt("008", int(value))
}

/*
Gets the current communication mode.
*/
func(t *TwissTorr84FS) GetCommunicationMode() (int, error) {
	return t.GetInt("008")
}

/*
Sets the soft start status.

Parameters:
	- value: 0 to disable, 1 to enable.
*/
func(t *TwissTorr84FS) SetSoftStart(value int) error {
	status, _ := t.GetStatus()
	if status != 0 {
		return fmt.Errorf("the device must be off to set the soft start")
	}

	if(value < 0 || value > 1) {
		return fmt.Errorf("invalid value for soft start")
	}
	return t.SetInt("100", int(value))
}

/*
Gets the current soft start status.
*/
func(t *TwissTorr84FS) GetSoftStart() (int, error) {
	return t.GetInt("100")
}

/*
Sets the R1 set point type

Parameters:
	- value:
		0 = Pressure
		1 = Power
		2 = Time
		3 = Normal
		4 = Pressure (only if gauge is connected)
*/
func (t *TwissTorr84FS) SetR1SetpointType(value int) error {
	if value < 0 || value > 4 {
		return fmt.Errorf("invalid value for R1 set point type")
	}
	return t.SetInt("101", int(value))
}

/*
Sets R1 set point value.
*/
func (t *TwissTorr84FS) SetR1Setpoint(value int) error {
	return t.SetInt("102", int(value))
}

/*
Gets the R1 set point value.
*/
func (t *TwissTorr84FS) GetR1Setpoint() (int, error) {
	return t.GetInt("102")
}

/*
Sets the time between the pump start and the set
point check (in seconds)
*/
func (t *TwissTorr84FS) SetTimeDelay(value uint) error {
	if value > 99999 {
		return fmt.Errorf("invalid value for time delay")
	}
	return t.SetInt("103", int(value))
}

/*
Gets the target time between the pump start and
the set point check (in seconds)
*/
func (t *TwissTorr84FS) GetTimeDelay() (int, error) {
	return t.GetInt("103")
}

/*
Sets the signal activation type

Parameters:
	- value:
		0 = high level active
		1 = low level active
*/
func (t *TwissTorr84FS) SetSignalActivationType(value uint) error {
	if value > 1 {
		return fmt.Errorf("invalid value for signal activation type")
	}
	return t.SetInt("104", int(value))
}

/*
Gets the signal activation type
*/
func (t *TwissTorr84FS) GetSignalActivationType() (int, error) {
	return t.GetInt("104")
}

/*
Sets the hysterisis set point (in % of value)

Parameters:
	- value: 0 to 100
*/
func (t *TwissTorr84FS) SetHysteresis(value uint) error {
	if value > 100 {
		return fmt.Errorf("invalid value for hysteresis")
	}
	return t.SetInt("105", int(value))
}

/*
Gets the hysterisis set point (in % of value)
*/
func (t *TwissTorr84FS) GetHysteresis() (int, error) {
	return t.GetInt("105")
}

/*
Sets water cooling status

Parameters:
	- value: 0 to NO, 1 to YES
*/
func (t *TwissTorr84FS) SetWaterCooling(value uint) error {
	if value > 1 {
		return fmt.Errorf("invalid value for water cooling")
	}
	return t.SetInt("106", int(value))
}

/*
Gets water cooling status
*/
func (t *TwissTorr84FS) GetWaterCooling() (int, error) {
	return t.GetInt("106")
}

/*
Sets the active stop status
*/
func (t *TwissTorr84FS) SetActiveStop(value uint) error {
	if value > 1 {
		return fmt.Errorf("invalid value for active stop")
	}
	return t.SetInt("107", int(value))
}

/*
Gets the active stop status
*/
func (t *TwissTorr84FS) GetActiveStop() (int, error) {
	return t.GetInt("107")
}

/*
Sets interlock type

Parameters:
	- value: 0 for impulse, 1 for continuos
*/
func (t *TwissTorr84FS) SetInterlockType(value uint) error {
	if value > 1 {
		return fmt.Errorf("invalid value for interlock type")
	}
	return t.SetInt("110", int(value))
}

/*
Gets interlock type
*/
func (t *TwissTorr84FS) GetInterlockType() (int, error) {
	return t.GetInt("110")
}

/*
Sets the analog output type: output voltage signal proportional
to frequency, power, temperature or pressure.

Parameters:
	- value:
		0 = frequency
		1 = power (default)
		2 = temperature
		3 = pressure
*/
func (t *TwissTorr84FS) SetAnalogOutputType(value uint) error {
	if value > 3 {
		return fmt.Errorf("invalid value for analog output type")
	}
	return t.SetInt("111", int(value))
}

/*
Gets the analog output type
*/
func (t *TwissTorr84FS) GetAnalogOutputType() (int, error) {
	return t.GetInt("111")
}

/*
Sets low speed frequency

Parameters:
	- value: desired frequency in Hz (default=1100)
*/
func (t *TwissTorr84FS) SetLowSpeedFrequency(value uint) error {
	if value > 9999 {
		return fmt.Errorf("invalid value for low speed frequency")
	}
	return t.SetInt("117", int(value))
}

/*
Gets low speed frequency
*/
func (t *TwissTorr84FS) GetLowSpeedFrequency() (int, error) {
	return t.GetInt("117")
}

/*
Sets rotational frequency setting

Parameters:
	- value: desired frequency in Hz (default=1350)
*/
func (t *TwissTorr84FS) SetRotationalFrequency(value uint) error {
	if value > 9999 {
		return fmt.Errorf("invalid value for rotational frequency")
	}
	return t.SetInt("118", int(value))
}

/*
Gets rotational frequency setting
*/
func (t *TwissTorr84FS) GetRotationalFrequency() (int, error) {
	return t.GetInt("118")
}

/*
Sets the vent valve operation

Parameters:
	- value: 0 for automatic, 1 for on command
*/
func (t *TwissTorr84FS) SetVentValveOperation(value uint) error {
	if value > 1 {
		return fmt.Errorf("invalid value for vent valve operation")
	}
	return t.SetInt("125", int(value))
}

/*
Gets the vent valve operation
*/
func (t *TwissTorr84FS) GetVentValveOperation() (int, error) {
	return t.GetInt("125")
}

/*
Sets vent valve opening delay (expressed in 0.2 sec)

Parameters:
	- value: desired delay (0 to 65535)
*/
func (t *TwissTorr84FS) SetVentValveOpeningDelay(value uint) error {
	if value > 65535 {
		return fmt.Errorf("invalid value for vent valve opening delay")
	}
	return t.SetInt("126", int(value))
}

/*
Gets vent valve opening delay
*/
func (t *TwissTorr84FS) GetVentValveOpeningDelay() (int, error) {
	return t.GetInt("126")
}

/*
Sets gauge set point type

Parameters:
	- value:
		0 = Frequency
		1 = Power
		2 = Time
		3 = Normal (default)
*/
func (t *TwissTorr84FS) SetGaugeSetpointType(value uint) error {
	if value > 3 {
		return fmt.Errorf("invalid value for gauge set point type")
	}
	return t.SetInt("136", int(value))
}

/*
Gets gauge set point type
*/
func (t *TwissTorr84FS) GetGaugeSetpointType() (int, error) {
	return t.GetInt("136")
}


/*
Sets gauge set point value

Parameters:
	- value: desired value in Hz, W or sec (default = 867)
*/
func (t *TwissTorr84FS) SetGaugeSetpoint(value uint) error {
	if value > 9999 {
		return fmt.Errorf("invalid value for gauge set point")
	}
	return t.SetInt("137", int(value))
}

/*
Gets gauge set point value
*/
func (t *TwissTorr84FS) GetGaugeSetpoint() (int, error) {
	return t.GetInt("137")
}
