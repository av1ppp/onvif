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

// Call_DeleteUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteUsersResponse.
func Call_DeleteUsers(ctx context.Context, dev *onvif.Device, request device.DeleteUsers) (device.DeleteUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteUsersResponse device.DeleteUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteUsers")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteUsers")
	}
	return reply.Body.DeleteUsersResponse, nil
}

// CallWithLogging_DeleteUsers works like Call_DeleteUsers but also logs the response body.
func CallWithLogging_DeleteUsers(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.DeleteUsers) (device.DeleteUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteUsersResponse device.DeleteUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteUsers")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteUsers")
	if err != nil {
		return reply.Body.DeleteUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteUsers")
	}
	return reply.Body.DeleteUsersResponse, nil
}
