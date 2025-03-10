package protocol

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
Calculates the checksum of a message using
the XOR operation.

Parameters:
	- message: byte array that represents the message.
*/
func calculateChecksum(message []byte) []byte {
	checksum := byte(0)
	for _, b := range []byte(message) {
		checksum ^= b
	}
	return []byte(fmt.Sprintf("%02X", checksum))
}

/*
Formats a write command to be sent to TwisTorr 84 FS AG 
using the window protocol.
*/
func (t *TwissTorr84FS) FormatWriteCmd(window string, data string) string {
	var sb strings.Builder

	sb.WriteByte(STX)
	sb.WriteByte(ADDR + byte(t.DeviceNumber))
	sb.Write([]byte(window))
	sb.WriteByte(COM_WRITE)
	sb.Write([]byte(data))
	sb.WriteByte(ETX)

	message := []byte(sb.String())
	sb.Write(calculateChecksum(message[1:]))

	return sb.String()
}

/*
Formats a read command to be sent to TwisTorr 84 FS AG
using the window protocol.
*/
func (t *TwissTorr84FS) FormatReadCmd(window string) string {
	var sb strings.Builder

	sb.WriteByte(STX)
	sb.WriteByte(ADDR + byte(t.DeviceNumber))
	sb.Write([]byte(window))
	sb.WriteByte(COM_READ)
	sb.WriteByte(ETX)

	message := []byte(sb.String())
	sb.Write(calculateChecksum(message[1:]))

	return sb.String()
}

/*
Validates the status response of a command.
*/
func ValidateStatusReponse(status byte) error{
	switch status {
	case ACK:
		return nil
	case NACK:
		return fmt.Errorf("the command execution failed")
	case UNKNOWN:
		return fmt.Errorf("the command is not a valid window")
	case TYPE_ERR:
		return fmt.Errorf("the data type is not valid for this window")
	case OUT_RANGE:
		return fmt.Errorf("the value is out of range")
	case DISABLED:
		return fmt.Errorf("the window is disabled")
	default:
		return fmt.Errorf("unknown status: %#v", status)
	}
}

/*
Requests a message to the device and returns the response
based on the window protocol.
*/
func (t *TwissTorr84FS) Request(message string) ([]byte, error) {
	t.Serial.Write(message)

	t.Serial.SetDelimiter(string(ETX))
	response, err := t.Serial.ReadUntil()
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}
	t.Serial.SetDelimiter("")

	checksum, err := t.Serial.Read(2)
	if err != nil {
		return nil, fmt.Errorf("error reading checksum: %v", err)
	}
	targetChecksum := calculateChecksum(response[1:])
	if !reflect.DeepEqual(checksum, targetChecksum) {
		return nil, fmt.Errorf("invalid checksum: %v", checksum)
	}

	end := len(response)-1
	return response[2:end], nil
}

/*
Request the value of a window and returns it as an integer.
*/
func (t *TwissTorr84FS) GetInt(window string) (int, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return 0, fmt.Errorf("error getting the value: %v", err)
	}

	responseStr := string(response)
	fmt.Printf("response: %v\n", responseStr)
	if(strings.Contains(responseStr, ".")) {
		fmt.Print("This is a float number")
		return 0, nil
	}

	return strconv.Atoi(responseStr)
}

/*
Request the value of a window and returns it as a float.
*/
func (t *TwissTorr84FS) GetFloat(window string) (float64, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return 0, fmt.Errorf("error getting the value: %v", err)
	}

	responseStr := string(response)
	fmt.Printf("response: %v\n", responseStr)
	return strconv.ParseFloat(responseStr, 64)
}

/*
Request the value of a window and returns it as a string.
*/
func (t *TwissTorr84FS) GetAlphanumeric(window string) (string, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return "", fmt.Errorf("error getting the value: %v", err)
	}
	return string(response), nil
}

/*
Sets the value of a window as an integer.
*/
func (t *TwissTorr84FS) SetInt(window string, value int) error {
	message := t.FormatWriteCmd(window, fmt.Sprintf("%d", value))
	response, err := t.Request(message)
	if err != nil {
		return fmt.Errorf("error setting the value: %v", err)
	}
	return ValidateStatusReponse(response[0])
}