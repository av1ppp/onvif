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

// Call_GetCapabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCapabilitiesResponse.
func Call_GetCapabilities(ctx context.Context, dev *onvif.Device, request device.GetCapabilities) (device.GetCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCapabilitiesResponse device.GetCapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCapabilities")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCapabilities")
	}
	return reply.Body.GetCapabilitiesResponse, nil
}

// CallWithLogging_GetCapabilities works like Call_GetCapabilities but also logs the response body.
func CallWithLogging_GetCapabilities(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetCapabilities) (device.GetCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCapabilitiesResponse device.GetCapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCapabilities")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCapabilities")
	if err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCapabilities")
	}
	return reply.Body.GetCapabilitiesResponse, nil
}
