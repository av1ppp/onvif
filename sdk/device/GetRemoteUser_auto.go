// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetRemoteUser forwards the call to onvif.Do then parses the payload of the reply as a GetRemoteUserResponse.
func GetRemoteUser(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetRemoteUser]) (device.GetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteUserResponse device.GetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteUser")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRemoteUser")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetRemoteUserResponse, err
}
