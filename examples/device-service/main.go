package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	goonvif "github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/device"
	sdkdevice "github.com/av1ppp/onvif/sdk/device"
)

var (
	addr     string
	username string
	password string
)

func init() {
	flag.StringVar(&addr, "addr", "192.168.13.14:80", "Device address")
	flag.StringVar(&username, "username", "admin", "Device username")
	flag.StringVar(&password, "password", "admin", "Device password")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	// Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      addr,
		Username:   username,
		Password:   password,
		HttpClient: new(http.Client),
		// Logger:     logging.New(),
	})
	if err != nil {
		panic(err)
	}

	// Getting device information
	systemDateAndTymeResponse, err := sdkdevice.GetSystemDateAndTime(ctx, dev, goonvif.Request(device.GetSystemDateAndTime{}))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}

	// Check error handling
	getCapabilitiesResponse, err := sdkdevice.GetCapabilities(ctx, dev, goonvif.Request(device.GetCapabilities{
		Category: "WRONG",
	}))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}
}
