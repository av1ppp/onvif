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

// Call_SetSystemDateAndTime forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemDateAndTimeResponse.
func Call_SetSystemDateAndTime(ctx context.Context, dev *onvif.Device, request device.SetSystemDateAndTime) (device.SetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemDateAndTimeResponse device.SetSystemDateAndTimeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemDateAndTime")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSystemDateAndTime")
	}
	return reply.Body.SetSystemDateAndTimeResponse, nil
}

// CallWithLogging_SetSystemDateAndTime works like Call_SetSystemDateAndTime but also logs the response body.
func CallWithLogging_SetSystemDateAndTime(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetSystemDateAndTime) (device.SetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemDateAndTimeResponse device.SetSystemDateAndTimeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemDateAndTime")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSystemDateAndTime")
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSystemDateAndTime")
	}
	return reply.Body.SetSystemDateAndTimeResponse, nil
}
