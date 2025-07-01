// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetRemoteDiscoveryMode forwards the call to onvif.Do then parses the payload of the reply as a GetRemoteDiscoveryModeResponse.
func GetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetRemoteDiscoveryMode]) (device.GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse device.GetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRemoteDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetRemoteDiscoveryModeResponse, err
}
