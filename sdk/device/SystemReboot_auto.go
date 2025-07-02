// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SystemReboot forwards the call to onvif.Do then parses the payload of the reply as a SystemRebootResponse.
func SystemReboot(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SystemReboot]) (device.SystemRebootResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SystemRebootResponse device.SystemRebootResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SystemRebootResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SystemReboot")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SystemReboot")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SystemRebootResponse, err
}
