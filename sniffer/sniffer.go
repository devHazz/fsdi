package sniffer

import (
	"github.com/google/gopacket/pcap"
	"log"
)

type Sniffer struct {
	netInterface string
	buffer       int32
	filter       string
}

func NewSniffer(netInterface string, filter string) *Sniffer {
	return &Sniffer{netInterface: netInterface, buffer: int32(1600), filter: filter}
}

func deviceExists(netInterface string) bool {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic(err)
	}
	for _, device := range devices {
		if device.Name == netInterface {
			return true
		}
	}
	return false
}

func (s Sniffer) Sniff() (*pcap.Handle, error) {
	if b := deviceExists(s.netInterface); b == false {
		log.Panic("network interface does not exist: ", s.netInterface)
	}
	handle, err := pcap.OpenLive(s.netInterface, s.buffer, false, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	err = handle.SetBPFFilter(s.filter)
	return handle, nil
}
