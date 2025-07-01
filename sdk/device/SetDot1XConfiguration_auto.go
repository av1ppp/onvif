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

// Call_SetDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDot1XConfigurationResponse.
func Call_SetDot1XConfiguration(ctx context.Context, dev *onvif.Device, request device.SetDot1XConfiguration) (device.SetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDot1XConfigurationResponse device.SetDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDot1XConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDot1XConfiguration")
	}
	return reply.Body.SetDot1XConfigurationResponse, nil
}

// CallWithLogging_SetDot1XConfiguration works like Call_SetDot1XConfiguration but also logs the response body.
func CallWithLogging_SetDot1XConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetDot1XConfiguration) (device.SetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDot1XConfigurationResponse device.SetDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDot1XConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDot1XConfiguration")
	if err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDot1XConfiguration")
	}
	return reply.Body.SetDot1XConfigurationResponse, nil
}
