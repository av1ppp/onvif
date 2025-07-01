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

// Call_GetSystemLog forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemLogResponse.
func Call_GetSystemLog(ctx context.Context, dev *onvif.Device, request device.GetSystemLog) (device.GetSystemLogResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemLogResponse device.GetSystemLogResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemLogResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemLog")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetSystemLogResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemLog")
	}
}

// CallWithLogging_GetSystemLog works like Call_GetSystemLog but also logs the response body.
func CallWithLogging_GetSystemLog(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetSystemLog) (device.GetSystemLogResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemLogResponse device.GetSystemLogResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemLogResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemLog")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemLog")
		return reply.Body.GetSystemLogResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemLog")
	}
}
