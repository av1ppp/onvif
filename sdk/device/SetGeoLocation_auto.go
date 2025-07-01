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

// Call_SetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a SetGeoLocationResponse.
func Call_SetGeoLocation(ctx context.Context, dev *onvif.Device, request device.SetGeoLocation) (device.SetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetGeoLocationResponse device.SetGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetGeoLocation")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetGeoLocation")
	}
	return reply.Body.SetGeoLocationResponse, nil
}

// CallWithLogging_SetGeoLocation works like Call_SetGeoLocation but also logs the response body.
func CallWithLogging_SetGeoLocation(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetGeoLocation) (device.SetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetGeoLocationResponse device.SetGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetGeoLocation")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetGeoLocation")
	if err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetGeoLocation")
	}
	return reply.Body.SetGeoLocationResponse, nil
}
