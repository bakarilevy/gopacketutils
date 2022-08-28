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
	"HTTP": "tcp and port 80",
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

func SetTimeout(minutes int) (time.Duration) {
	timeout := time.Duration(minutes) * time.Minute
	return timeout
}

//Alias SetDeviceA
func SetDevice(device string) (*pcap.Handle) {
	handle := SetDeviceA(device)
	return handle
}

// Setting default timeout to 5 minutes for tesitng.
func SetDeviceA(device string) (*pcap.Handle) {
	handle := SetDeviceB(device, 5)
	return handle
}

func SetDeviceB(device string, minute_timeout int) (*pcap.Handle) {
	var (
		snapshot_len int32 = 1024
		promiscuous bool = true
	)
	timeout := SetTimeout(minute_timeout)
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
		if (strings.Contains(device.Description, "Wi-Fi") || strings.Contains(device.Description, "WiFi"))  && (!strings.Contains(device.Description, "Virtual")) {
			default_device_name = device.Name
		}
	}

	log.Println("Setting default device to: " + default_device_name)
	default_device := SetDevice(default_device_name)
	return default_device
}

func GetDefaultWiFiDeviceInfo() {
	
	for _, device := range ALL_DEVICES {
		if (strings.Contains(device.Description, "Wi-Fi") || strings.Contains(device.Description, "WiFi"))  && (!strings.Contains(device.Description, "Virtual")) {
			log.Println("Name: ", device.Name)
			log.Println("Description: ", device.Description)
			log.Println()
		
			for _, address := range device.Addresses {
				log.Println("- IP Address: ", address.IP)
				log.Println("- Subnet Mask: ", address.Netmask)
				log.Println()
			}
		}
	}
}

func ShowPacket(handle *pcap.Handle) {
	packet_source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packet_source.Packets() {
		fmt.Println(packet)
	}
}
