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

// Call_GetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteUserResponse.
func Call_GetRemoteUser(ctx context.Context, dev *onvif.Device, request device.GetRemoteUser) (device.GetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteUserResponse device.GetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteUser")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRemoteUser")
	}
	return reply.Body.GetRemoteUserResponse, nil
}

// CallWithLogging_GetRemoteUser works like Call_GetRemoteUser but also logs the response body.
func CallWithLogging_GetRemoteUser(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetRemoteUser) (device.GetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteUserResponse device.GetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteUser")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRemoteUser")
	if err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRemoteUser")
	}
	return reply.Body.GetRemoteUserResponse, nil
}
