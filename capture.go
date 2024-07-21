package main

import (
	"log"
	"os"

	"github.com/google/gopacket"

	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Capture() {
	handle, err := pcap.OpenLive("\\Device\\NPF_{95B57B90-88CC-4B60-A478-70FEAFFC6BF7}", int32(1024), false, pcap.BlockForever)

	var fileName = "test.txt"

	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		ipLayer := packet.Layer(layers.LayerTypeIPv4)

		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			Filter(fileName)
			f.WriteString("This is Source ip: ")
			f.Write([]byte(ip.SrcIP.String()))
			f.WriteString("\n")
			f.WriteString("This is Destination ip: ")
			f.Write([]byte(ip.SrcIP.String()))
			f.WriteString("\n")

		}

	}

}
