package cereal

import (
	"fmt"
	"github.com/devHazz/fsdi/model"
	"github.com/devHazz/fsdi/parser"
	"golang.org/x/exp/slices"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func Deserialize(p string) (model.FSDPacket, error) {
	var packet model.FSDPacket
	id, valid := parser.GetType(p)
	if valid {
		packet.Id = id
		packet.CommandType = slices.IndexFunc(parser.CommandPrefixes, func(c string) bool { return c == p[:1] })
		p := parser.CleanPayload(p[3:])
		if strings.Contains(p, "::") {
			p = strings.Replace(p, "::", ":", -1)
		}
		payloadFormat := strings.Split(p, ":")

		packetStruct := parser.FSDPacketStructs[packet.Id]
		v := reflect.ValueOf(&packetStruct).Elem()

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
						if v := parser.MultiPayload(corrValue); v != nil {
							corrValue = v[0]
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
			packet.Data = v.Interface()
			return packet, nil
		}
		return model.FSDPacket{}, fmt.Errorf("could not handle struct: %q", v.Type().Name())
	}
	return model.FSDPacket{}, fmt.Errorf("could not get fsd packet type for: %s", p[1:3])
}

func Serialize(p model.FSDPacket) string {
	v := reflect.ValueOf(p)
	payload := ""
	prefixId := v.FieldByName("CommandType").Int()
	commandPrefix := parser.CommandPrefixes[prefixId]
	command := ""
	for k, val := range parser.FSDTypes {
		id := v.FieldByName("Id").Int()
		if int64(val) == id {
			command = k
			break
		}
	}
	payload += commandPrefix + command
	data := v.FieldByName("Data")
	v = reflect.Indirect(data)
	for i := 0; i < v.Elem().NumField(); i++ {
		f := v.Elem().Field(i).Interface()
		switch reflect.ValueOf(f).Kind() {
		case reflect.String:
			payload += f.(string)
		case reflect.Int:
			payload += strconv.Itoa(f.(int))
		default:
			fmt.Printf("Serialization Error: %s | Invalid Field Type At: %s", reflect.TypeOf(v).Name(), reflect.TypeOf(v.Elem().Field(i)).Name())
		}
		if i != (v.Elem().NumField() - 1) {
			payload += ":"
		}
	}

	return payload
}
