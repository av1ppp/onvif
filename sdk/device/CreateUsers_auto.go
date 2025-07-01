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

// Call_CreateUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateUsersResponse.
func Call_CreateUsers(ctx context.Context, dev *onvif.Device, request device.CreateUsers) (device.CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse device.CreateUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateUsers")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateUsers")
	}
	return reply.Body.CreateUsersResponse, nil
}

// CallWithLogging_CreateUsers works like Call_CreateUsers but also logs the response body.
func CallWithLogging_CreateUsers(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.CreateUsers) (device.CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse device.CreateUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateUsers")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateUsers")
	if err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateUsers")
	}
	return reply.Body.CreateUsersResponse, nil
}
