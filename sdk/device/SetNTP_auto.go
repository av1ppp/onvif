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

// Call_SetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNTPResponse.
func Call_SetNTP(ctx context.Context, dev *onvif.Device, request device.SetNTP) (device.SetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNTPResponse device.SetNTPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNTP")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNTP")
	}
	return reply.Body.SetNTPResponse, nil
}

// CallWithLogging_SetNTP works like Call_SetNTP but also logs the response body.
func CallWithLogging_SetNTP(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetNTP) (device.SetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNTPResponse device.SetNTPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNTP")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNTP")
	if err != nil {
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNTP")
	}
	return reply.Body.SetNTPResponse, nil
}
