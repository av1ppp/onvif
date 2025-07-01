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

// Call_GetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetGeoLocationResponse.
func Call_GetGeoLocation(ctx context.Context, dev *onvif.Device, request device.GetGeoLocation) (device.GetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGeoLocationResponse device.GetGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetGeoLocation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetGeoLocation")
	}
}

// CallWithLogging_GetGeoLocation works like Call_GetGeoLocation but also logs the response body.
func CallWithLogging_GetGeoLocation(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetGeoLocation) (device.GetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGeoLocationResponse device.GetGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetGeoLocation")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetGeoLocation")
		return reply.Body.GetGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetGeoLocation")
	}
}
