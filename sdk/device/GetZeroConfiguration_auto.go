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

// Call_GetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetZeroConfigurationResponse.
func Call_GetZeroConfiguration(ctx context.Context, dev *onvif.Device, request device.GetZeroConfiguration) (device.GetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetZeroConfigurationResponse device.GetZeroConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	}
	return reply.Body.GetZeroConfigurationResponse, nil
}

// CallWithLogging_GetZeroConfiguration works like Call_GetZeroConfiguration but also logs the response body.
func CallWithLogging_GetZeroConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetZeroConfiguration) (device.GetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetZeroConfigurationResponse device.GetZeroConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetZeroConfiguration")
	if err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	}
	return reply.Body.GetZeroConfigurationResponse, nil
}
