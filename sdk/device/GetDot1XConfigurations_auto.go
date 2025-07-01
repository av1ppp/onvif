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

// Call_GetDot1XConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationsResponse.
func Call_GetDot1XConfigurations(ctx context.Context, dev *onvif.Device, request device.GetDot1XConfigurations) (device.GetDot1XConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationsResponse device.GetDot1XConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	}
	return reply.Body.GetDot1XConfigurationsResponse, nil
}

// CallWithLogging_GetDot1XConfigurations works like Call_GetDot1XConfigurations but also logs the response body.
func CallWithLogging_GetDot1XConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDot1XConfigurations) (device.GetDot1XConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationsResponse device.GetDot1XConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDot1XConfigurations")
	if err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	}
	return reply.Body.GetDot1XConfigurationsResponse, nil
}
