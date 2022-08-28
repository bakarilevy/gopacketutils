package base

import (
	"gopacketutils/primitives"
	"github.com/google/gopacket/pcap"
)

func ReadHttpTraffic(handle *pcap.Handle) {
	primitives.SetFilter(primitives.FILTERS["HTTP"], handle)
	primitives.ShowPacket(handle)
}
