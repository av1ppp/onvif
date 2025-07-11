// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetNetworkDefaultGateway forwards the call to onvif.Do then parses the payload of the reply as a GetNetworkDefaultGatewayResponse.
func GetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetNetworkDefaultGateway]) (device.GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse device.GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkDefaultGateway")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNetworkDefaultGateway")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetNetworkDefaultGatewayResponse, err
}
