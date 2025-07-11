// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetRelayOutputs forwards the call to onvif.Do then parses the payload of the reply as a GetRelayOutputsResponse.
func GetRelayOutputs(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetRelayOutputs]) (device.GetRelayOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRelayOutputsResponse device.GetRelayOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRelayOutputs")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRelayOutputs")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetRelayOutputsResponse, err
}
