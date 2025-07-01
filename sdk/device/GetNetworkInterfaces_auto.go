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

// Call_GetNetworkInterfaces forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkInterfacesResponse.
func Call_GetNetworkInterfaces(ctx context.Context, dev *onvif.Device, request device.GetNetworkInterfaces) (device.GetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkInterfacesResponse device.GetNetworkInterfacesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkInterfaces")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNetworkInterfaces")
	}
}

// CallWithLogging_GetNetworkInterfaces works like Call_GetNetworkInterfaces but also logs the response body.
func CallWithLogging_GetNetworkInterfaces(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetNetworkInterfaces) (device.GetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkInterfacesResponse device.GetNetworkInterfacesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkInterfaces")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNetworkInterfaces")
		return reply.Body.GetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNetworkInterfaces")
	}
}
