package main

import (
	"encoding/json"
	"fmt"
	"github.com/devHazz/fsdi/parser"
	"github.com/devHazz/fsdi/sniffer"
	"github.com/google/gopacket"
	"io"
	"log"
	"os"
)

type Config struct {
	InterfaceName string `json:"interfaceName"`
	SnapLen       string `json:"SnapLen"`
}

var config Config

func init() {
	path := "config.json"
	if _, err := os.Open(path); err != nil {
		log.Panic("error loading FSDI config")
	}
	f, _ := os.Open(path)
	configData, _ := io.ReadAll(f)
	_ = json.Unmarshal(configData, &config)
}

func main() {
	s := sniffer.NewSniffer(config.InterfaceName, "tcp src port 6809")
	h, _ := s.Sniff()
	defer h.Close()
	source := gopacket.NewPacketSource(h, h.LinkType())
	for packet := range source.Packets() {
		if packet != nil && packet.ApplicationLayer() != nil {
			payload := string(packet.ApplicationLayer().Payload())
			if payload == "" || len(payload) == 0 {
				src := packet.NetworkLayer().NetworkFlow().Src().String()
				dst := packet.NetworkLayer().NetworkFlow().Dst().String()
				fmt.Printf("could not get payload for packet: %s->%s", src, dst)
			}
			parser.Parse(payload)
		}
	}
}
