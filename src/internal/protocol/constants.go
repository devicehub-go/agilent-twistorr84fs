package protocol

import manager "gitlab.cnpem.br/ads/vacuum-devices/communication"

type SerialType uint8

type TwissTorr84FS struct {
	DeviceNumber int
	Serial manager.Communication
}

const (
	ACK       byte = 0x06
	NACK      byte = 0x15
	UNKNOWN   byte = 0x32
	TYPE_ERR  byte = 0x33
	OUT_RANGE byte = 0x34
	DISABLED  byte = 0x35

	STX       byte = 0x02
	ADDR      byte = 0x80
	ETX       byte = 0x03
	
	COM_READ  byte = 0x30
	COM_WRITE byte = 0x31
)
