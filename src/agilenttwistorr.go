/*
Author: Leonardo Rossi Leao
Contact: leonardo.leao@cnpem.br
Last Update: April 2nd, 2025
Descriptions: This module abstract the communication with
Agilent TwisTorr 84 FS AG Rack Controller, providing an interface
to monitor and control the device.
*/

package main

import (
	"fmt"

	"gitlab.cnpem.br/ads/vacuum-devices/agilenttwistorr/internal/protocol"
	manager "gitlab.cnpem.br/ads/vacuum-devices/communication"
)

/*
Creates a new Agilent TwisTorr instance that allows to communicate
with the device through RS232 protocol.

Notes:
	- In case of RS232 communication, the device number must be zero.
	  Otherwise, this variable represents the unit address of the device
	  in the RS485 network.
*/
func New(port string, baudrate int, parity string, databits int,
		 stopbits int, deviceNumber int) (*protocol.TwissTorr84FS) {

	options := manager.CommunicationOptions {
		Protocol: manager.Serial,
		Serial: manager.SerialOptions{
			Port: port,
			BaudRate: baudrate,
			Parity: parity,
			DataBits: databits,
			StopBits: stopbits,
		},
		Delimiter: "",
	}

	return &protocol.TwissTorr84FS{
		DeviceNumber: deviceNumber,
		Serial: manager.NewCommunication(options),
	}
}

func main() {
	twisstor := New("COM4", 9600, "N", 8, 1, 0)
	fmt.Println(twisstor.Serial.Connect())
	fmt.Println(twisstor.SetGaugeSetpointType(3))
	fmt.Println(twisstor.GetGaugeSetpointType())
}