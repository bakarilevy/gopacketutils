package main

import (
	"log"
	"time"
	"strings"
//	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

//ALL_DEVICES := pcap.FindAllDevs()

func ListDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	// Print device information
	log.Println("Devices found:\n")
	for _, device := range devices {
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

func SetTimeout(minutes int) (time.Duration) {
	timeout := time.Duration(minutes) * time.Minute
	return timeout
}

// Setting default timeout to 3 minutes for tesitng.
func SetDevice(device string) (*pcap.Handle){
	var (
		snapshot_len int32 = 1024
		promiscuous bool = true
	)
	timeout := SetTimeout(3)
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

func SetFilterEx(filter string, handle *pcap.Handle) {
	// Set the filter on as specific network interface
}

func SetDefaultWiFiDevice() (*pcap.Handle) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	var default_device_name string
	for _, device := range devices {
		if (strings.Contains(device.Description, "Wi-Fi") || strings.Contains(device.Description, "WiFi"))  && (!strings.Contains(device.Description, "Virtual")) {
			default_device_name = device.Name
		}
	}

	log.Println("Setting default device to: " + default_device_name)
	default_device := SetDevice(default_device_name)
	return default_device
}

func GetDefaultWiFiDeviceInfo() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}
	
	for _, device := range devices {
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


func main() {
	GetDefaultWiFiDeviceInfo()
	//_ = SetDefaultWiFiDevice()
}