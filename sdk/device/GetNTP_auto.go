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

// Call_GetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNTPResponse.
func Call_GetNTP(ctx context.Context, dev *onvif.Device, request device.GetNTP) (device.GetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNTPResponse device.GetNTPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNTP")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetNTPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNTP")
	}
	return reply.Body.GetNTPResponse, nil
}

// CallWithLogging_GetNTP works like Call_GetNTP but also logs the response body.
func CallWithLogging_GetNTP(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetNTP) (device.GetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNTPResponse device.GetNTPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNTP")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNTP")
	if err != nil {
		return reply.Body.GetNTPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNTP")
	}
	return reply.Body.GetNTPResponse, nil
}
