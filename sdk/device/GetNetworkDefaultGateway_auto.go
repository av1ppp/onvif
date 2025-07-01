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

// Call_GetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkDefaultGatewayResponse.
func Call_GetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request device.GetNetworkDefaultGateway) (device.GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse device.GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkDefaultGateway")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNetworkDefaultGateway")
	}
	return reply.Body.GetNetworkDefaultGatewayResponse, nil
}

// CallWithLogging_GetNetworkDefaultGateway works like Call_GetNetworkDefaultGateway but also logs the response body.
func CallWithLogging_GetNetworkDefaultGateway(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetNetworkDefaultGateway) (device.GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse device.GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkDefaultGateway")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNetworkDefaultGateway")
	if err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNetworkDefaultGateway")
	}
	return reply.Body.GetNetworkDefaultGatewayResponse, nil
}
