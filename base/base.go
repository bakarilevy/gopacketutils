package base

import (
	"gopacketutils/primitives"
	"github.com/google/gopacket/pcap"
)

//Alias SetDeviceA
func SetDevice(device string) (*pcap.Handle) {
	handle := primitives.SetDeviceA(device)
	return handle
}

func ReadFtpTraffic(handle *pcap.Handle) {
	primitives.SetFilter(primitives.FILTERS["FTP"], handle)
	primitives.ReadPackets(handle)
	defer handle.Close()
}

func ReadHttpTraffic(handle *pcap.Handle) {
	primitives.SetFilter(primitives.FILTERS["HTTP"], handle)
	primitives.ReadPackets(handle)
	defer handle.Close()
}

func ReadDnsTraffic(handle *pcap.Handle) {
	primitives.SetFilter(primitives.FILTERS["DNS"], handle)
	primitives.ReadPackets(handle)
	defer handle.Close()
}