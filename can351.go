package SMACanMessages

import (
	"encoding/binary"
	"github.com/brutella/can"
)

// Details of the CAN message with ID 351 - Contains the battery information
/*
Battery Charge Voltage (16 bit)
Battery Charge Current Limitation (16 bit)
Battery Discharge Current Limitation (16 bit)
Battery Discharge Voltage Limitation (16 bit)
*/

type Can351 struct {
	data [8]byte
}

//func New(d []byte) Can351 {
//	c := Can351{binary.LittleEndian.Uint16(d[0:]), uint16(binary.LittleEndian.Uint16(d[2:])), uint16(binary.LittleEndian.Uint16(d[4:])), binary.LittleEndian.Uint16(d[6:])}
//	return c
//}

func NewCan351(vSetpoint float32, iSetpoint float32, iDischarge float32, vDischarge float32) Can351 {
	this := Can351{}
	binary.LittleEndian.PutUint16(this.data[0:], uint16(vSetpoint*10.0))
	binary.LittleEndian.PutUint16(this.data[2:], uint16(iSetpoint*10.0))
	binary.LittleEndian.PutUint16(this.data[4:], uint16(iDischarge*10.0))
	binary.LittleEndian.PutUint16(this.data[6:], uint16(vDischarge*10.0))
	return this
}

func (setpoint *Can351) SetVSetpoint(vSetpoint float32) {
	binary.LittleEndian.PutUint16(setpoint.data[0:], uint16(vSetpoint*10.0))
}

func (setpoint *Can351) VSetpoint() float32 {
	return float32(binary.LittleEndian.Uint16(setpoint.data[0:])) / 10
}

func (setpoint *Can351) SetISetpoint(iSetpoint float32) {
	binary.LittleEndian.PutUint16(setpoint.data[2:], uint16(iSetpoint*10.0))
}

func (setpoint *Can351) ISetpoint() float32 {
	return float32(binary.LittleEndian.Uint16(setpoint.data[2:])) / 10
}

func (setpoint *Can351) SetIDischarge(iDischarge float32) {
	binary.LittleEndian.PutUint16(setpoint.data[4:], uint16(iDischarge*10.0))
}

func (setpoint *Can351) IDischarge() float32 {
	return float32(binary.LittleEndian.Uint16(setpoint.data[4:])) / 10
}

func (setpoint *Can351) SetVDischarge(vDischarge float32) {
	binary.LittleEndian.PutUint16(setpoint.data[6:], uint16(vDischarge*10.0))
}

func (setpoint *Can351) VDischarge() float32 {
	return float32(binary.LittleEndian.Uint16(setpoint.data[6:])) / 10
}

func (setpoint *Can351) CANData() []byte {
	return setpoint.data[0:]
}

func (setpoint *Can351) Frame() can.Frame {
	return can.Frame{
		ID:     0x351,
		Length: 8,
		Flags:  0,
		Res0:   0,
		Res1:   0,
		Data:   setpoint.data,
	}
}
