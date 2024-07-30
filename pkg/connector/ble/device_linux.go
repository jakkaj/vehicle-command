package ble

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
)

func newDevice() (ble.Device, error) {
	devices := []string{"hci0", "hci1", "hci2"} // Add more devices if needed
	var device ble.Device
	var err error

	for _, dev := range devices {
		// Extract the integer part from the device string
		devIndexStr := strings.TrimPrefix(dev, "hci")
		devIndex, convErr := strconv.Atoi(devIndexStr)
		if convErr != nil {
			continue
		}

		device, err = linux.NewDevice(ble.OptDeviceID(devIndex))
		if err == nil {
			return device, nil
		}
	}

	return nil, fmt.Errorf("all devices failed")
}
