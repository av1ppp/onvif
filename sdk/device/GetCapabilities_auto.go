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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCapabilities")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCapabilities")
	}
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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCapabilities")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCapabilities")
		return reply.Body.GetCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCapabilities")
	}
}
