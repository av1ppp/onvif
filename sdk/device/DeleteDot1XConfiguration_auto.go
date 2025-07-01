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

// Call_DeleteDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteDot1XConfigurationResponse.
func Call_DeleteDot1XConfiguration(ctx context.Context, dev *onvif.Device, request device.DeleteDot1XConfiguration) (device.DeleteDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteDot1XConfigurationResponse device.DeleteDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteDot1XConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteDot1XConfiguration")
	}
	return reply.Body.DeleteDot1XConfigurationResponse, nil
}

// CallWithLogging_DeleteDot1XConfiguration works like Call_DeleteDot1XConfiguration but also logs the response body.
func CallWithLogging_DeleteDot1XConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.DeleteDot1XConfiguration) (device.DeleteDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteDot1XConfigurationResponse device.DeleteDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.DeleteDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteDot1XConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteDot1XConfiguration")
	if err != nil {
		return reply.Body.DeleteDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteDot1XConfiguration")
	}
	return reply.Body.DeleteDot1XConfigurationResponse, nil
}
