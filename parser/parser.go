package parser

import "fmt"

var FSDTypes = map[string]int{
	"DI": SERVER_IDENTIFY,
	"ID": CLIENT_IDENTIFY,
	"ER": SERVER_ERROR,
	"AA": ADD_ATC,
	"AP": ADD_PILOT,
	"TM": MOTD,
}

var FSDTypeNames = map[int]string{
	SERVER_IDENTIFY: "Server Identification",
	CLIENT_IDENTIFY: "Client Identification",
	ADD_ATC:         "Add Controller Request",
	ADD_PILOT:       "Add Pilot Request",
	MOTD:            "Message of the Day",
	DELETE_ATC:      "Remove Controller Client",
	DELETE_PILOT:    "Remove Pilot Client",
	CLIENT_REQUEST:  "Client Request",
	CLIENT_RESPONSE: "Client Response",
}

type FSDHeader struct {
	Id   int
	Name string
}

func GetType(packetType string) (*FSDHeader, bool) {
	if len(packetType) > 3 {
		packetType = packetType[1:3]
	}
	if _, ok := FSDTypes[packetType]; !ok {
		fmt.Println("FSD Packet Type Not Supported: ", packetType)
		return nil, false
	}
	id := FSDTypes[packetType]
	return &FSDHeader{Id: id, Name: FSDTypeNames[id]}, true
}

func Parse(payload string) {
}
