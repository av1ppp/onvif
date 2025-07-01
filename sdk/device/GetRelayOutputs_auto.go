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

// Call_GetRelayOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRelayOutputsResponse.
func Call_GetRelayOutputs(ctx context.Context, dev *onvif.Device, request device.GetRelayOutputs) (device.GetRelayOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRelayOutputsResponse device.GetRelayOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRelayOutputs")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRelayOutputs")
	}
	return reply.Body.GetRelayOutputsResponse, nil
}

// CallWithLogging_GetRelayOutputs works like Call_GetRelayOutputs but also logs the response body.
func CallWithLogging_GetRelayOutputs(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetRelayOutputs) (device.GetRelayOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRelayOutputsResponse device.GetRelayOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRelayOutputs")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRelayOutputs")
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRelayOutputs")
	}
	return reply.Body.GetRelayOutputsResponse, nil
}
