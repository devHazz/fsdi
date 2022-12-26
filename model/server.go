package model

type ServerIdentify struct {
	Version         string
	VersionExpanded string
}

type MessageOfTheDay struct {
	Callsign string
	Message  string
}

type FSDHeader struct {
	Id   int
	Name string
}

type FSDPacket struct {
	*FSDHeader
	data interface{}
}
