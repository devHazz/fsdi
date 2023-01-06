package parser

import (
	"github.com/devHazz/fsdi/model"
	"strings"
)

var FSDTypes = map[string]int{
	"DI": SERVER_IDENTIFY,
	"ID": CLIENT_IDENTIFY,
	"ER": SERVER_ERROR,
	"AA": ADD_ATC,
	"AP": ADD_PILOT,
	"TM": MOTD,
	"DA": DELETE_ATC,
	"DP": DELETE_PILOT,
}

var FSDPacketStructs = map[int]interface{}{
	SERVER_IDENTIFY: model.ServerIdentify{},
	CLIENT_IDENTIFY: model.ClientIdentify{},
	ADD_ATC:         model.AddControllerRequest{},
	ADD_PILOT:       model.AddPilotRequest{},
	MOTD:            model.MessageOfTheDay{},
	DELETE_ATC:      model.DeleteClientRequest{},
	DELETE_PILOT:    model.DeleteClientRequest{},
	CLIENT_REQUEST:  model.ClientRequest{},
	//CLIENT_RESPONSE: model.ClientResponse{},
}

var CommandPrefixes = []string{
	"#", "%", "$", "@",
}

func GetType(packetType string) (int, bool) {
	if len(packetType) > 3 {
		packetType = packetType[1:3]
	}
	if _, ok := FSDTypes[packetType]; !ok {
		//fmt.Println("FSD Packet Type Not Supported: ", packetType)
		return 0, false
	}
	id := FSDTypes[packetType]
	return id, true
}

func CleanPayload(payload string) string {
	p := payload
	if strings.Contains(payload, ":SERVER:") {
		p = strings.Replace(payload, ":SERVER:", ":", 1)
	}
	return CleanEscSequence(p)
}

func MultiPayload(field string) []string {
	np := CleanEscSequence(field)
	for _, v := range CommandPrefixes {
		if strings.Contains(np, v) {
			payloadArr := strings.Split(np, v)
			if _, valid := GetType(v + payloadArr[1]); valid {
				return payloadArr
			}
		}
	}
	return nil
}
