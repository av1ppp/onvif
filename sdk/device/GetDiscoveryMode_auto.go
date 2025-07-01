// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetDiscoveryMode forwards the call to onvif.Do then parses the payload of the reply as a GetDiscoveryModeResponse.
func GetDiscoveryMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetDiscoveryMode]) (device.GetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDiscoveryModeResponse device.GetDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDiscoveryMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetDiscoveryModeResponse, err
}
