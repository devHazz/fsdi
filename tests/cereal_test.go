package tests

import (
	cereal "github.com/devHazz/fsdi"
	"github.com/devHazz/fsdi/model"
	"testing"
)

func TestSerialize(t *testing.T) {
	s := model.FSDPacket{
		Id:          1,
		CommandType: 2,
		Data: model.ServerIdentify{
			Unk1:            "SERVER",
			Unk2:            "CLIENT",
			VersionExpanded: "VATSIM FSD V3.4m",
			Token:           "",
		},
	}
	k := cereal.Serialize(s)
	t.Logf("%s", k)
}

func TestDeserialize(t *testing.T) {
	payload := "$DISERVER:CLIENT:VATSIM FSD V3.4m:b378d821"
	packet, err := cereal.Deserialize(payload)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", packet)
}
