package parser

import (
	"fmt"
	"github.com/devHazz/fsdi/model"
	"log"
	"reflect"
	"strconv"
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

func GetType(packetType string) (*model.FSDHeader, bool) {
	if len(packetType) > 3 {
		packetType = packetType[1:3]
	}
	if _, ok := FSDTypes[packetType]; !ok {
		//fmt.Println("FSD Packet Type Not Supported: ", packetType)
		return nil, false
	}
	id := FSDTypes[packetType]
	return &model.FSDHeader{Id: id}, true
}

func CleanPayload(payload string) string {
	p := payload
	if strings.Contains(payload, ":SERVER:") {
		p = strings.Replace(payload, ":SERVER:", ":", 1)
	}
	return CleanEscSequence(p)
}

func isCommand(payload string) bool {
	for _, v := range CommandPrefixes {
		if payload[:1] == v {
			if _, valid := GetType(payload); valid {
				return true
			}
		}
	}
	return false
}

func parseMultiPayload(field string) []string {
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

func Parse(payload string) {
	var packet model.FSDPacket
	header, valid := GetType(payload)
	if valid {
		packet.FSDHeader = header
		p := CleanPayload(payload[3:])
		if strings.Contains(p, "::") {
			p = strings.Replace(p, "::", ":", -1)
		}
		payloadFormat := strings.Split(p, ":")

		packetStruct := FSDPacketStructs[packet.FSDHeader.Id]
		v := reflect.ValueOf(&packetStruct).Elem()

		//Handling Bad Packet
		if !v.IsNil() {
			tmp := reflect.New(v.Elem().Type()).Elem()
			tmp.Set(v.Elem())
			if tmp.NumField() == 0 {
				log.Panic("no fields on struct: ", v.Elem().Type().String())
			}

			for i := 0; i < tmp.NumField(); i++ {
				corrValue := payloadFormat[i]
				//Handling packets which may have more than 1 packets worth of data
				if tmp.Type().Field(i).Name == "NetworkId" {
					if len(corrValue) > 7 {
						if v := parseMultiPayload(corrValue); v != nil {
							corrValue = parseMultiPayload(corrValue)[0]
						}
					}
				}
				switch tmp.Type().Field(i).Type.Kind() {
				case reflect.String:
					tmp.Field(i).SetString(corrValue)
				case reflect.Int:
					n, _ := strconv.Atoi(corrValue)
					tmp.Field(i).SetInt(int64(n))
				}
			}
			v.Set(tmp)
			fmt.Printf("%s: %+v\n", v.Elem().Type().String(), packetStruct)
		}
	}
}
