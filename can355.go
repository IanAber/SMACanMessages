package SMACanMessages

import "encoding/binary"
import "github.com/brutella/can"

// Details of the CAN message with ID 305 - Contains the battery information
/*
State Of Charge (16 bit)
State Of Health (16 bit)
High Resolution State Of Charge (16 bit)
Unused (16 bit)
*/

type Can355 struct {
	data [8]byte
}

func NewCan355(stateOfCharge uint16, stateOfHealth uint16, hiResolutionStateOfCharge float32) Can355 {
	this := Can355{}
	binary.LittleEndian.PutUint16(this.data[0:], stateOfCharge)
	binary.LittleEndian.PutUint16(this.data[2:], stateOfHealth)
	binary.LittleEndian.PutUint16(this.data[4:], uint16(hiResolutionStateOfCharge*100.0))
	return this
}

func (canMessage *Can355) SetStateOfCharge(stateOfCharge uint16) {
	binary.LittleEndian.PutUint16(canMessage.data[0:], stateOfCharge)
}

func (canMessage *Can355) StateOfCharge() uint16 {
	return binary.LittleEndian.Uint16(canMessage.data[0:])
}

func (canMessage *Can355) SetStateOfHealth(stateOfHealth uint16) {
	binary.LittleEndian.PutUint16(canMessage.data[2:], stateOfHealth)
}

func (canMessage *Can355) StateOfHealth() uint16 {
	return binary.LittleEndian.Uint16(canMessage.data[2:])
}

func (canMessage *Can355) SetHighResStateOfCharge(highResStateOfCharge float32) {
	binary.LittleEndian.PutUint16(canMessage.data[4:], uint16(highResStateOfCharge*100.0))
}

func (canMessage *Can355) HighResStateOfCharge() float32 {
	return float32(binary.LittleEndian.Uint16(canMessage.data[4:])) / 100
}

func (canMessage *Can355) CANData() []byte {
	return canMessage.data[0:]
}

func (canMessage *Can355) Frame() can.Frame {
	return can.Frame{
		ID:     0x355,
		Length: 8,
		Flags:  0,
		Res0:   0,
		Res1:   0,
		Data:   canMessage.data,
	}
}
