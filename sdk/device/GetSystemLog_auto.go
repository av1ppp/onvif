// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetSystemLog forwards the call to onvif.Do then parses the payload of the reply as a GetSystemLogResponse.
func GetSystemLog(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetSystemLog]) (device.GetSystemLogResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemLogResponse device.GetSystemLogResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetSystemLogResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemLog")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemLog")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetSystemLogResponse, err
}
