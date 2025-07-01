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

// Call_SetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteUserResponse.
func Call_SetRemoteUser(ctx context.Context, dev *onvif.Device, request device.SetRemoteUser) (device.SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse device.SetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteUser")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRemoteUser")
	}
	return reply.Body.SetRemoteUserResponse, nil
}

// CallWithLogging_SetRemoteUser works like Call_SetRemoteUser but also logs the response body.
func CallWithLogging_SetRemoteUser(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetRemoteUser) (device.SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse device.SetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteUser")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRemoteUser")
	if err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRemoteUser")
	}
	return reply.Body.SetRemoteUserResponse, nil
}
