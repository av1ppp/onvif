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

// Call_SystemReboot forwards the call to dev.CallMethod() then parses the payload of the reply as a SystemRebootResponse.
func Call_SystemReboot(ctx context.Context, dev *onvif.Device, request device.SystemReboot) (device.SystemRebootResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SystemRebootResponse device.SystemRebootResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SystemRebootResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SystemReboot")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SystemRebootResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SystemReboot")
	}
}

// CallWithLogging_SystemReboot works like Call_SystemReboot but also logs the response body.
func CallWithLogging_SystemReboot(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SystemReboot) (device.SystemRebootResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SystemRebootResponse device.SystemRebootResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SystemRebootResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SystemReboot")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SystemReboot")
		return reply.Body.SystemRebootResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SystemReboot")
	}
}
