package SMACanMessages

// Details of the CAN message with ID 010 - Contains relay sand generator state etc.

type Can307 struct {
	relayState byte
	func1      byte
	func2      byte
	func3      byte
	func4      byte
	func5      byte
	func6      byte
	func7      byte
}

func NewCan307(d []byte) Can307 {
	c := Can307{d[0], d[1], d[2], d[3], d[4], d[5], d[6], d[7]}
	return c
}

func (c Can307) Relay1Master() bool {
	return (c.relayState & 1) != 0
}

func (c Can307) Relay2Master() bool {
	return (c.relayState & 2) != 0
}

func (c Can307) Relay1Slave1() bool {
	return (c.relayState & 4) != 0
}

func (c Can307) Relay2Slave1() bool {
	return (c.relayState & 8) != 0
}

func (c Can307) Relay1Slave2() bool {
	return (c.relayState & 0x10) != 0
}

func (c Can307) Relay2Slave2() bool {
	return (c.relayState & 0x20) != 0
}

func (c Can307) GnRun() bool {
	return (c.func1 & 0x01) != 0
}

func (c Can307) GnRunSlave1() bool {
	return (c.func1 & 0x02) != 0
}

func (c Can307) GnRunSlave2() bool {
	return (c.func1 & 0x04) != 0
}

func (c Can307) AutoGn() bool {
	return (c.func2 & 0x01) != 0
}

func (c Can307) AutoLodExt() bool {
	return (c.func2 & 0x02) != 0
}

func (c Can307) AutoLodSoc() bool {
	return (c.func2 & 0x04) != 0
}

func (c Can307) Tm1() bool {
	return (c.func2 & 0x08) != 0
}

func (c Can307) Tm2() bool {
	return (c.func2 & 0x10) != 0
}

func (c Can307) ExtPwrDer() bool {
	return (c.func2 & 0x20) != 0
}

func (c Can307) ExtVfOk() bool {
	return (c.func2 & 0x40) != 0
}

func (c Can307) GdOn() bool {
	return (c.func2 & 0x80) != 0
}

func (c Can307) Error() bool {
	return (c.func3 & 0x01) != 0
}

func (c Can307) Run() bool {
	return (c.func3 & 0x02) != 0
}

func (c Can307) BatFan() bool {
	return (c.func3 & 0x04) != 0
}

func (c Can307) AcdCir() bool {
	return (c.func3 & 0x08) != 0
}

func (c Can307) MccBatFan() bool {
	return (c.func3 & 0x10) != 0
}

func (c Can307) MccAutoLod() bool {
	return (c.func3 & 0x20) != 0
}

func (c Can307) Chp() bool {
	return (c.func3 & 0x40) != 0
}

func (c Can307) ChpAdd() bool {
	return (c.func3 & 0x80) != 0
}

func (c Can307) SiComRemote() bool {
	return (c.func4 & 0x01) != 0
}

func (c Can307) Overload() bool {
	return (c.func4 & 0x02) != 0
}

func (c Can307) ExtSrcConn() bool {
	return (c.func5 & 0x01) != 0
}

func (c Can307) Silent() bool {
	return (c.func5 & 0x02) != 0
}

func (c Can307) Current() bool {
	return (c.func5 & 0x04) != 0
}

func (c Can307) FeedSelfC() bool {
	return (c.func5 & 0x08) != 0
}

func (c Can307) Esave() bool {
	return (c.func5 & 0x10) != 0
}
