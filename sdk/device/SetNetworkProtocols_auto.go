// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetNetworkProtocols forwards the call to onvif.Do then parses the payload of the reply as a SetNetworkProtocolsResponse.
func SetNetworkProtocols(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetNetworkProtocols]) (device.SetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkProtocolsResponse device.SetNetworkProtocolsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkProtocols")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNetworkProtocols")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetNetworkProtocolsResponse, err
}
