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

// Call_SetSystemFactoryDefault forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func Call_SetSystemFactoryDefault(ctx context.Context, dev *onvif.Device, request device.SetSystemFactoryDefault) (device.SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse device.SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemFactoryDefault")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSystemFactoryDefault")
	}
	return reply.Body.SetSystemFactoryDefaultResponse, nil
}

// CallWithLogging_SetSystemFactoryDefault works like Call_SetSystemFactoryDefault but also logs the response body.
func CallWithLogging_SetSystemFactoryDefault(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetSystemFactoryDefault) (device.SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse device.SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemFactoryDefault")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSystemFactoryDefault")
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSystemFactoryDefault")
	}
	return reply.Body.SetSystemFactoryDefaultResponse, nil
}
