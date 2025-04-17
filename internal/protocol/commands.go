package protocol

/*
Turns the turbo pump on.
*/
func (t *TwissTorr84FS) TurnOn() error {
	return t.SetInt("000", 1)
}

/*
Turns the turbo pump off.
*/
func (t *TwissTorr84FS) TurnOff() error {
	return t.SetInt("000", 0)
}

/*
Gets the operation status (on/off) of the turbo pump.
*/
func (t *TwissTorr84FS) GetStatus() (int, error) {
	return t.GetInt("000")
}

/*
Resets the pump life counter.
*/
func (t *TwissTorr84FS) ResetPumpLife() error {
	return t.SetInt("109", 0)
}

/*
Opens the vent valve.
*/
func (t *TwissTorr84FS) OpenVentValve() error {
	return t.SetInt("122", 1)
}

/*
Closes the vent valve.
*/
func (t *TwissTorr84FS) CloseVentValve() error {
	return t.SetInt("122", 0)
}

/*
Gets the vent valve status (open/closed).
*/
func (t *TwissTorr84FS) GetVentValveStatus() (int, error) {
	return t.GetInt("122")
}