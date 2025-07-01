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

// Call_GetServiceCapabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServiceCapabilitiesResponse.
func Call_GetServiceCapabilities(ctx context.Context, dev *onvif.Device, request device.GetServiceCapabilities) (device.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse device.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	}
	return reply.Body.GetServiceCapabilitiesResponse, nil
}

// CallWithLogging_GetServiceCapabilities works like Call_GetServiceCapabilities but also logs the response body.
func CallWithLogging_GetServiceCapabilities(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetServiceCapabilities) (device.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse device.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetServiceCapabilities")
	if err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	}
	return reply.Body.GetServiceCapabilitiesResponse, nil
}
