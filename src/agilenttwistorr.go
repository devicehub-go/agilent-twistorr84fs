/*
Author: Leonardo Rossi Leao
Contact: leonardo.leao@cnpem.br
Last Update: March 7th, 2025
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

	//options := manager.CommunicationOptions {
	//	Protocol: manager.Serial,
	//	Serial: manager.SerialOptions{
	//		Port: port,
	//		BaudRate: baudrate,
	//		Parity: parity,
	//		DataBits: databits,
	//		StopBits: stopbits,
	//	},
	//	Delimiter: "",
	//}

	options := manager.CommunicationOptions {
		Protocol: manager.TCP,
		TCP: manager.TCPOptions{
			Host: "localhost",
			Port: 65432,
		},
		Delimiter: "",
	}

	return &protocol.TwissTorr84FS{
		DeviceNumber: deviceNumber,
		Serial: manager.NewCommunication(options),
	}
}

func main() {
	twisstor := New("COM1", 9600, "None", 8, 1, 0)
	twisstor.Serial.Connect()
	fmt.Print(twisstor.TurnOn())
}