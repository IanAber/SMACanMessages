package SMACanMessages

import "encoding/binary"
import "github.com/brutella/can"

// Details of the CAN message with ID 305 - Contains the battery information
/*
Actual Battery Voltage (16 bit)
Actual Battery Current (16 bit)
Actual Temperature (16 bit)
Unused (16 bit)
*/

type Can356 struct {
	data [8]byte
}

func NewCan356(voltage float32, current float32, temperature float32) Can356 {
	this := Can356{}
	binary.LittleEndian.PutUint16(this.data[0:], uint16(voltage*100.0))
	binary.LittleEndian.PutUint16(this.data[2:], uint16(current*10.0))
	binary.LittleEndian.PutUint16(this.data[4:], uint16(temperature*10.0))
	return this
}

func (canMessage *Can356) SetVoltage(voltage float32) {
	binary.LittleEndian.PutUint16(canMessage.data[0:], uint16(voltage*100))
}

func (canMessage *Can356) Voltage() float32 {
	return float32(binary.LittleEndian.Uint16(canMessage.data[0:])) / 100.0
}

func (canMessage *Can356) SetCurrent(current float32) {
	binary.LittleEndian.PutUint16(canMessage.data[2:], uint16(current*10.0))
}

func (canMessage *Can356) Current() float32 {
	return float32(binary.LittleEndian.Uint16(canMessage.data[2:])) / 10.0
}

func (canMessage *Can356) SetTemperature(temperature float32) {
	binary.LittleEndian.PutUint16(canMessage.data[4:], uint16(temperature*10.0))
}

func (canMessage *Can356) Temperature() float32 {
	return float32(binary.LittleEndian.Uint16(canMessage.data[4:])) / 10.0
}

func (canMessage *Can356) CANData() []byte {
	return canMessage.data[0:]
}

func (canMessage *Can356) Frame() can.Frame {
	return can.Frame{
		ID:     0x356,
		Length: 8,
		Flags:  0,
		Res0:   0,
		Res1:   0,
		Data:   canMessage.data,
	}
}
