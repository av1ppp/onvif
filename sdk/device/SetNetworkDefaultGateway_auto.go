// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkDefaultGatewayResponse.
func Call_SetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request device.SetNetworkDefaultGateway) (device.SetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkDefaultGatewayResponse device.SetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkDefaultGateway")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkDefaultGateway")
	}
}

// CallWithLogging_SetNetworkDefaultGateway works like Call_SetNetworkDefaultGateway but also logs the response body.
func CallWithLogging_SetNetworkDefaultGateway(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetNetworkDefaultGateway) (device.SetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkDefaultGatewayResponse device.SetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkDefaultGateway")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNetworkDefaultGateway")
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkDefaultGateway")
	}
}
