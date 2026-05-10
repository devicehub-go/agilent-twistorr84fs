package agilenttwistorr

import (
	"github.com/devicehub-go/agilent-twistorr84fs/protocol"
	"github.com/devicehub-go/unicomm"
)

/*
Creates a new Agilent TwisTorr instance that allows to communicate
with the device through RS232 protocol.

Notes:
  - In case of RS232 communication, the device number must be zero.
    Otherwise, this variable represents the unit address of the device
    in the RS485 network.
*/
func New(options unicomm.Options, deviceNumber int) *protocol.TwissTorr84FS {
	options.Delimiter = ""
	return &protocol.TwissTorr84FS{
		DeviceNumber:  deviceNumber,
		Communication: unicomm.New(options),
	}
}
