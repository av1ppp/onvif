// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetRemoteDiscoveryMode forwards the call to onvif.Do then parses the payload of the reply as a SetRemoteDiscoveryModeResponse.
func SetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetRemoteDiscoveryMode]) (device.SetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteDiscoveryModeResponse device.SetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteDiscoveryMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRemoteDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetRemoteDiscoveryModeResponse, err
}
