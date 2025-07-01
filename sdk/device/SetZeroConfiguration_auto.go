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

// Call_SetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetZeroConfigurationResponse.
func Call_SetZeroConfiguration(ctx context.Context, dev *onvif.Device, request device.SetZeroConfiguration) (device.SetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetZeroConfigurationResponse device.SetZeroConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetZeroConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetZeroConfiguration")
	}
}

// CallWithLogging_SetZeroConfiguration works like Call_SetZeroConfiguration but also logs the response body.
func CallWithLogging_SetZeroConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetZeroConfiguration) (device.SetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetZeroConfigurationResponse device.SetZeroConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetZeroConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetZeroConfiguration")
		return reply.Body.SetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetZeroConfiguration")
	}
}
