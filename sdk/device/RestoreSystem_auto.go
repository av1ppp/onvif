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

// Call_RestoreSystem forwards the call to dev.CallMethod() then parses the payload of the reply as a RestoreSystemResponse.
func Call_RestoreSystem(ctx context.Context, dev *onvif.Device, request device.RestoreSystem) (device.RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse device.RestoreSystemResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RestoreSystemResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RestoreSystem")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RestoreSystemResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RestoreSystem")
	}
	return reply.Body.RestoreSystemResponse, nil
}

// CallWithLogging_RestoreSystem works like Call_RestoreSystem but also logs the response body.
func CallWithLogging_RestoreSystem(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.RestoreSystem) (device.RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse device.RestoreSystemResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.RestoreSystemResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RestoreSystem")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RestoreSystem")
	if err != nil {
		return reply.Body.RestoreSystemResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RestoreSystem")
	}
	return reply.Body.RestoreSystemResponse, nil
}
