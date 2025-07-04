// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetSystemDateAndTime forwards the call to onvif.Do then parses the payload of the reply as a SetSystemDateAndTimeResponse.
func SetSystemDateAndTime(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetSystemDateAndTime]) (device.SetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemDateAndTimeResponse device.SetSystemDateAndTimeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemDateAndTime")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSystemDateAndTime")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetSystemDateAndTimeResponse, err
}
