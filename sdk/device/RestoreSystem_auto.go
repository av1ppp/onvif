// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// RestoreSystem forwards the call to onvif.Do then parses the payload of the reply as a RestoreSystemResponse.
func RestoreSystem(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.RestoreSystem]) (device.RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse device.RestoreSystemResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.RestoreSystemResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RestoreSystem")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RestoreSystem")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RestoreSystemResponse, err
}
