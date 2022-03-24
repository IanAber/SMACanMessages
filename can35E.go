package SMACanMessages

import (
	"github.com/brutella/can"
)

// Details of the CAN message with ID 35E - Contains the battery manufacturers name

type Can35E struct {
	data [8]byte
}

func NewCan35E(manufacturer string) Can35E {
	this := Can35E{}
	copy(this.data[:], manufacturer)
	return this
}

func (canMessage *Can35E) SetManufacturer(manufacturer string) {
	copy(canMessage.data[:], manufacturer)
}

func (canMessage *Can35E) Manufacturer() string {
	return string(canMessage.data[:])
}

func (canMessage *Can35E) CANData() []byte {
	return canMessage.data[0:]
}

func (canMessage *Can35E) Frame() can.Frame {
	return can.Frame{
		ID:     0x35E,
		Length: 8,
		Flags:  0,
		Res0:   0,
		Res1:   0,
		Data:   canMessage.data,
	}
}
