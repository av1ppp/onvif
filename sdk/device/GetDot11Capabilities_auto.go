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

// Call_GetDot11Capabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11CapabilitiesResponse.
func Call_GetDot11Capabilities(ctx context.Context, dev *onvif.Device, request device.GetDot11Capabilities) (device.GetDot11CapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11CapabilitiesResponse device.GetDot11CapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	}
	return reply.Body.GetDot11CapabilitiesResponse, nil
}

// CallWithLogging_GetDot11Capabilities works like Call_GetDot11Capabilities but also logs the response body.
func CallWithLogging_GetDot11Capabilities(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDot11Capabilities) (device.GetDot11CapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11CapabilitiesResponse device.GetDot11CapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDot11Capabilities")
	if err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	}
	return reply.Body.GetDot11CapabilitiesResponse, nil
}
