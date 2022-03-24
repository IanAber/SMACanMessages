package SMACanMessages

import "encoding/binary"

// Details of the CAN message with ID 305 - Contains the battery information

type Can305 struct {
	vBatt   uint16
	iBatt   int16
	tBatt   int16
	socBatt uint16
}

func NewCan305(d []byte) Can305 {
	c := Can305{binary.LittleEndian.Uint16(d[0:]), int16(binary.LittleEndian.Uint16(d[2:])), int16(binary.LittleEndian.Uint16(d[4:])), binary.LittleEndian.Uint16(d[6:])}
	return c
}

func (c Can305) VBatt() float32 {
	return float32(c.vBatt) / 10
}

func (c Can305) IBatt() float32 {
	return float32(c.iBatt) / 10
}

func (c Can305) TBatt() float32 {
	return float32(c.tBatt) / 10
}

func (c Can305) SocBatt() float32 {
	return float32(c.socBatt) / 10
}
