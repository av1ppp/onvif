// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetDiscoveryMode forwards the call to onvif.Do then parses the payload of the reply as a SetDiscoveryModeResponse.
func SetDiscoveryMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetDiscoveryMode]) (device.SetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDiscoveryModeResponse device.SetDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDiscoveryMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetDiscoveryModeResponse, err
}
