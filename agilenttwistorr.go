/*
Author: Leonardo Rossi Leao
Contact: leonardo.leao@cnpem.br
Last Update: April 2nd, 2025
Descriptions: This module abstract the communication with
Agilent TwisTorr 84 FS AG Rack Controller, providing an interface
to monitor and control the device.
*/

package agilenttwistorr

import (
	"gitlab.cnpem.br/ads/vacuum-devices/agilenttwistorr/internal/protocol"
	manager "gitlab.cnpem.br/ads/vacuum-devices/communication"
)

type AgilentTwisTorr = protocol.TwissTorr84FS

/*
Creates a new Agilent TwisTorr instance that allows to communicate
with the device through RS232 protocol.

Notes:
	- In case of RS232 communication, the device number must be zero.
	  Otherwise, this variable represents the unit address of the device
	  in the RS485 network.
*/
func New(options manager.CommunicationOptions, deviceNumber int) (*protocol.TwissTorr84FS) {
	options.Delimiter = ""
	return &protocol.TwissTorr84FS{
		DeviceNumber: deviceNumber,
		Serial: manager.NewCommunication(options),
	}
}