package model

type ServerIdentify struct {
	Unk1            string
	Unk2            string
	VersionExpanded string
	Token           string
}

type MessageOfTheDay struct {
	Unk1     string
	Callsign string
	Message  string
}

type FSDPacket struct {
	Id          int
	CommandType int
	Data        interface{}
}
