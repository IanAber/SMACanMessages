package SMACanMessages

import "math"

// Details of the CAN message with ID 010 - Contains the frequency of the mains

type Can010 struct {
	fDecimal uint8
	fInteger uint8
}

func NewCan010(d []byte) Can010 {
	c := Can010{d[2], d[3]}
	return c
}

func (c Can010) Frequency() float64 {
	return math.Round(((float64(c.fDecimal)/256)+float64(c.fInteger))*100) / 100
}
