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

// Call_GetSystemDateAndTime forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemDateAndTimeResponse.
func Call_GetSystemDateAndTime(ctx context.Context, dev *onvif.Device, request device.GetSystemDateAndTime) (device.GetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemDateAndTimeResponse device.GetSystemDateAndTimeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemDateAndTime")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemDateAndTime")
	}
	return reply.Body.GetSystemDateAndTimeResponse, nil
}

// CallWithLogging_GetSystemDateAndTime works like Call_GetSystemDateAndTime but also logs the response body.
func CallWithLogging_GetSystemDateAndTime(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetSystemDateAndTime) (device.GetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemDateAndTimeResponse device.GetSystemDateAndTimeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemDateAndTime")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemDateAndTime")
	if err != nil {
		return reply.Body.GetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemDateAndTime")
	}
	return reply.Body.GetSystemDateAndTimeResponse, nil
}
