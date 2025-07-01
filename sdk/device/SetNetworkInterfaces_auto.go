// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetNetworkInterfaces forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkInterfacesResponse.
func Call_SetNetworkInterfaces(ctx context.Context, dev *onvif.Device, request device.SetNetworkInterfaces) (device.SetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkInterfacesResponse device.SetNetworkInterfacesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	}
	return reply.Body.SetNetworkInterfacesResponse, nil
}

// CallWithLogging_SetNetworkInterfaces works like Call_SetNetworkInterfaces but also logs the response body.
func CallWithLogging_SetNetworkInterfaces(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetNetworkInterfaces) (device.SetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkInterfacesResponse device.SetNetworkInterfacesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNetworkInterfaces")
	if err != nil {
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	}
	return reply.Body.SetNetworkInterfacesResponse, nil
}
