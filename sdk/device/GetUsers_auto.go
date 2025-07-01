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

// Call_GetUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a GetUsersResponse.
func Call_GetUsers(ctx context.Context, dev *onvif.Device, request device.GetUsers) (device.GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse device.GetUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetUsers")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetUsers")
	}
	return reply.Body.GetUsersResponse, nil
}

// CallWithLogging_GetUsers works like Call_GetUsers but also logs the response body.
func CallWithLogging_GetUsers(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetUsers) (device.GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse device.GetUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetUsers")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetUsers")
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetUsers")
	}
	return reply.Body.GetUsersResponse, nil
}
