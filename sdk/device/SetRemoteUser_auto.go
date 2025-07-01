// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetRemoteUser forwards the call to onvif.Do then parses the payload of the reply as a SetRemoteUserResponse.
func SetRemoteUser(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetRemoteUser]) (device.SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse device.SetRemoteUserResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteUser")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRemoteUser")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetRemoteUserResponse, err
}
