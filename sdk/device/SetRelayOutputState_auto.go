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

// Call_SetRelayOutputState forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRelayOutputStateResponse.
func Call_SetRelayOutputState(ctx context.Context, dev *onvif.Device, request device.SetRelayOutputState) (device.SetRelayOutputStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputStateResponse device.SetRelayOutputStateResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputState")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRelayOutputState")
	}
	return reply.Body.SetRelayOutputStateResponse, nil
}

// CallWithLogging_SetRelayOutputState works like Call_SetRelayOutputState but also logs the response body.
func CallWithLogging_SetRelayOutputState(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetRelayOutputState) (device.SetRelayOutputStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputStateResponse device.SetRelayOutputStateResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputState")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRelayOutputState")
	if err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRelayOutputState")
	}
	return reply.Body.SetRelayOutputStateResponse, nil
}
