// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetUser forwards the call to onvif.Do then parses the payload of the reply as a SetUserResponse.
func SetUser(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetUser]) (device.SetUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetUserResponse device.SetUserResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetUser")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetUser")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetUserResponse, err
}
