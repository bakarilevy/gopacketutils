package main

import (
	"gopacketutils/base"
	"gopacketutils/primitives"
)

func main() {
	handle := primitives.SetDefaultWiFiDevice()
	base.ReadHttpTraffic(handle)
}