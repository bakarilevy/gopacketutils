package primitives

import (
	"fmt"
	"log"
	"time"
	"strings"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var ALL_DEVICES = FindAllDevices()

var FILTERS = map[string]string{
	"FTP":"tcp and port 21",
	"HTTP": "tcp and port 80",
	"DNS":"dns",
}

// Users shouldn't need to use this function, simply used to populate devices on startup to avoid multiple loops over pcap.FindAllDevs
func FindAllDevices() ([]pcap.Interface) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln("Error while trying to find all devices")
		panic(err)
	}
	return devices
}

func ListDevices() {

	fmt.Println("Devices found:\n")
	for _, device := range ALL_DEVICES {
		fmt.Println("Name: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println()
		for _, address := range device.Addresses {
			fmt.Println("- IP Address: ", address.IP)
			fmt.Println("- Subnet Mask: ", address.Netmask)
			fmt.Println()
		}
	}
}

func SetTimeout(seconds int) (time.Duration) {
	timeout := time.Duration(seconds) * time.Second
	return timeout
}

// Setting default timeout to -1 seconds to immedeately flush packets.
func SetDeviceA(device string) (*pcap.Handle) {
	handle := SetDeviceB(device, -1)
	return handle
}

func SetDeviceB(device string, seconds_timeout int) (*pcap.Handle) {
	var (
		snapshot_len int32 = 65535
		promiscuous bool = true
	)
	timeout := SetTimeout(seconds_timeout)
	handle, err := SetDeviceEx(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatalln("Error while getting handle to network device")
		log.Fatalln(err)
	}
	return handle
}

func SetDeviceEx(device string, snapshot_len int32, promiscuous bool, timeout time.Duration) (*pcap.Handle, error) {
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	
	if err != nil {
		return nil, err
	}

	return handle, nil
}

func SetFilter(filter string, handle *pcap.Handle) {
	// Set the filter on a specific network interface
	err := handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatalln("Error while attempting to set filter on interface")
		log.Fatalln(err)
	} else {
		log.Println("Successfully applied filter: " + filter + " to interface")
	}

}

func SetDefaultWiFiDevice() (*pcap.Handle) {

	var default_device_name string
	for _, device := range ALL_DEVICES {
		if (strings.Contains(device.Description, "Wi-Fi") || strings.Contains(device.Description, "WiFi"))  && (!strings.Contains(device.Description, "Virtual") && !strings.Contains(device.Description, "virtual")) {
			default_device_name = device.Name
		}
	}

	log.Println("Setting default device to: " + default_device_name)
	default_device := SetDeviceA(default_device_name)
	return default_device
}

func GetDefaultWiFiDeviceInfo() {
	
	for _, device := range ALL_DEVICES {
		if (strings.Contains(device.Description, "Wi-Fi") || strings.Contains(device.Description, "WiFi"))  && (!strings.Contains(device.Description, "Virtual") && !strings.Contains(device.Description, "virtual")) {
			fmt.Println("Name: ", device.Name)
			fmt.Println("Description: ", device.Description)
			fmt.Println()
		
			for _, address := range device.Addresses {
				fmt.Println("- IP Address: ", address.IP)
				fmt.Println("- Subnet Mask: ", address.Netmask)
				fmt.Println()
			}
		}
	}
}

func ReadPackets(handle *pcap.Handle) {
	packet_source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packet_source.Packets() {
		fmt.Println(packet.String())
	}
}
