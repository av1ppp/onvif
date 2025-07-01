package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	goonvif "github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/device"
	sdkdevice "github.com/av1ppp/onvif/sdk/device"
	"github.com/av1ppp/onvif/xsd/onvif"
)

const (
	login    = "admin"
	password = "Supervisor"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.13.14:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	// Getting device information
	systemDateAndTymeResponse, err := sdkdevice.GetSystemDateAndTime(ctx, dev, &goonvif.Req[device.GetSystemDateAndTime]{
		Body: device.GetSystemDateAndTime{},
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}

	getCapabilitiesResponse, err := sdkdevice.GetCapabilities(ctx, dev, &goonvif.Req[device.GetCapabilities]{
		Body: device.GetCapabilities{Category: "All"},
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

	// Creating user
	createUserResponse, err := sdkdevice.CreateUsers(ctx, dev, &goonvif.Req[device.CreateUsers]{
		Body: device.CreateUsers{
			User: onvif.User{
				Username:  "TestUser",
				Password:  "TestPassword",
				UserLevel: "User",
			},
		},
	})
	if err != nil {
		log.Println(err)
	} else {
		// You could use https://github.com/av1ppp/onvif/gosoap for pretty printing response
		fmt.Println(createUserResponse)
	}

}
