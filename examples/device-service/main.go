package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/device"
	sdkdevice "github.com/av1ppp/onvif/sdk/device"
)

var (
	addr     string
	username string
	password string
)

func init() {
	flag.StringVar(&addr, "addr", "192.168.12.34:80", "Device address")
	flag.StringVar(&username, "username", "admin", "Device username")
	flag.StringVar(&password, "password", "admin", "Device password")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	// Getting an camera instance
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}

	// Getting device information
	systemDateAndTymeResponse, err := sdkdevice.GetSystemDateAndTime(ctx, dev, onvif.Request(device.GetSystemDateAndTime{}))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}

	// Check error handling
	getCapabilitiesResponse, err := sdkdevice.GetCapabilities(ctx, dev, onvif.Request(device.GetCapabilities{
		Category: "WRONG",
	}))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}
}
