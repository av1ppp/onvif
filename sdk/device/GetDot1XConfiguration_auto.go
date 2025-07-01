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

// Call_GetDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationResponse.
func Call_GetDot1XConfiguration(ctx context.Context, dev *onvif.Device, request device.GetDot1XConfiguration) (device.GetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationResponse device.GetDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot1XConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot1XConfiguration")
	}
	return reply.Body.GetDot1XConfigurationResponse, nil
}

// CallWithLogging_GetDot1XConfiguration works like Call_GetDot1XConfiguration but also logs the response body.
func CallWithLogging_GetDot1XConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDot1XConfiguration) (device.GetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationResponse device.GetDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot1XConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDot1XConfiguration")
	if err != nil {
		return reply.Body.GetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot1XConfiguration")
	}
	return reply.Body.GetDot1XConfigurationResponse, nil
}
