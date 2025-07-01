// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetNetworkDefaultGateway forwards the call to onvif.Do then parses the payload of the reply as a SetNetworkDefaultGatewayResponse.
func SetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetNetworkDefaultGateway]) (device.SetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkDefaultGatewayResponse device.SetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkDefaultGateway")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNetworkDefaultGateway")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetNetworkDefaultGatewayResponse, err
}
