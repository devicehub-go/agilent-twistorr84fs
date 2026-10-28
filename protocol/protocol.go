package protocol

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/devicehub-go/unicomm"
)

type TwissTorr84FS struct {
	DeviceNumber  int
	Communication unicomm.Unicomm

	mutex sync.Mutex
}

// Calculates the checksum of a message using XOR of bytes
func calculateChecksum(message []byte) []byte {
	checksum := byte(0)
	for _, b := range []byte(message) {
		checksum ^= b
	}
	return []byte(fmt.Sprintf("%02X", checksum))
}

// Establishes a connection with the device
func (t *TwissTorr84FS) Connect() error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.Communication.Connect()
}

// Closes the connection with the device
func (t *TwissTorr84FS) Disconnect() error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.Communication.Disconnect()
}

// Checks if the device is connected
func (t *TwissTorr84FS) IsConnected() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.Communication.IsConnected()
}

// Formats a write command to be sent to TwisTorr 84 FS AG
// using the window protocol
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

// Formats a read command to be sent to TwisTorr 84 FS AG
// using the window protocol
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

// Validates the status response of a command
func ValidateStatusReponse(status byte) error {
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

// Requests a message to the device and returns the response
func (t *TwissTorr84FS) Request(message string) ([]byte, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Communication.Write([]byte(message))

	response, err := t.Communication.ReadUntil(string(ETX))
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	// This time is required to wait for the device send the checksum
	time.Sleep(20 * time.Millisecond)

	checksum, err := t.Communication.Read(2)
	if err != nil {
		return nil, fmt.Errorf("error reading checksum: %v", err)
	}
	targetChecksum := calculateChecksum(response[1:])
	if !reflect.DeepEqual(checksum, targetChecksum) {
		return nil, fmt.Errorf("invalid checksum: %v", checksum)
	}

	end := len(response) - 1
	if end < 2 {
		return nil, fmt.Errorf("invalid response length: %d", len(response))
	}

	return response[2:end], nil
}

// Request the value of a window and returns it as an integer
func (t *TwissTorr84FS) GetInt(window string) (int, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return 0, fmt.Errorf("error getting the value: %v", err)
	}

	window_rx := string(response[0:3])
	if window_rx != window {
		return 0, fmt.Errorf("response window %s does not match requested window %s", window_rx, window)
	}

	responseStr := string(response[3:])
	if strings.Contains(responseStr, ".") {
		fmt.Print("This is a float number")
		return 0, nil
	}

	return strconv.Atoi(responseStr)
}

// Request the value of a window and returns it as a float
func (t *TwissTorr84FS) GetFloat(window string) (float64, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return 0, fmt.Errorf("error getting the value: %v", err)
	}

	window_rx := string(response[0:3])
	if window_rx != window {
		return 0, fmt.Errorf("response window %s does not match requested window %s", window_rx, window)
	}

	responseStr := string(response)
	return strconv.ParseFloat(responseStr, 64)
}

// Request the value of a window and returns it as a string.
func (t *TwissTorr84FS) GetAlphanumeric(window string) (string, error) {
	message := t.FormatReadCmd(window)
	response, err := t.Request(message)
	if err != nil {
		return "", fmt.Errorf("error getting the value: %v", err)
	}

	window_rx := string(response[0:3])
	if window_rx != window {
		return "", fmt.Errorf("response window %s does not match requested window %s", window_rx, window)
	}

	return string(response), nil
}

// Sets the value of a window as a logic
func (t *TwissTorr84FS) SetLogic(window string, value int) error {
	message := t.FormatWriteCmd(window, fmt.Sprintf("%d", value))
	response, err := t.Request(message)
	if err != nil {
		return fmt.Errorf("error setting the value: %v", err)
	}
	return ValidateStatusReponse(response[0])
}

// Sets the value of a window as a numeric
func (t *TwissTorr84FS) SetNumeric(window string, value int) error {
	message := t.FormatWriteCmd(window, fmt.Sprintf("%06d", value))
	response, err := t.Request(message)
	if err != nil {
		return fmt.Errorf("error setting the value: %v", err)
	}
	return ValidateStatusReponse(response[0])
}
