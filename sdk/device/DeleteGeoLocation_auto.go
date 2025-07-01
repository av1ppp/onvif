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

// Call_DeleteGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteGeoLocationResponse.
func Call_DeleteGeoLocation(ctx context.Context, dev *onvif.Device, request device.DeleteGeoLocation) (device.DeleteGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteGeoLocationResponse device.DeleteGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteGeoLocation")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteGeoLocation")
	}
	return reply.Body.DeleteGeoLocationResponse, nil
}

// CallWithLogging_DeleteGeoLocation works like Call_DeleteGeoLocation but also logs the response body.
func CallWithLogging_DeleteGeoLocation(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.DeleteGeoLocation) (device.DeleteGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteGeoLocationResponse device.DeleteGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteGeoLocation")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteGeoLocation")
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteGeoLocation")
	}
	return reply.Body.DeleteGeoLocationResponse, nil
}
