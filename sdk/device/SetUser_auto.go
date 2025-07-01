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

// Call_SetUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetUserResponse.
func Call_SetUser(ctx context.Context, dev *onvif.Device, request device.SetUser) (device.SetUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetUserResponse device.SetUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetUser")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetUser")
	}
	return reply.Body.SetUserResponse, nil
}

// CallWithLogging_SetUser works like Call_SetUser but also logs the response body.
func CallWithLogging_SetUser(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetUser) (device.SetUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetUserResponse device.SetUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetUser")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetUser")
	if err != nil {
		return reply.Body.SetUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetUser")
	}
	return reply.Body.SetUserResponse, nil
}
