package SMACanMessages

import "encoding/binary"

// Details of the CAN message with ID 010 - Contains charging information

type Can306 struct {
	sohBatt            uint16
	chargingProcedure  byte
	operatingState     byte
	activeErrorMessage uint16
	chargeSetPoint     uint16
}

func NewCan306(d []byte) Can306 {
	c := Can306{binary.LittleEndian.Uint16(d[0:]), d[2], d[3], binary.LittleEndian.Uint16(d[4:]), binary.LittleEndian.Uint16(d[6:])}
	return c
}

func (c Can306) SohBatt() float32 {
	return float32(c.sohBatt) / 10
}

func (c Can306) ChargingProcedure() byte {
	return c.chargingProcedure
}

func (c Can306) OperatingState() byte {
	return c.operatingState
}

func (c Can306) ActiveErrorMessage() uint16 {
	return c.activeErrorMessage
}

func (c Can306) ChargeSetPoint() float32 {
	return float32(c.chargeSetPoint) / 10
}
