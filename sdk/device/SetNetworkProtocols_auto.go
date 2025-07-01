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

// Call_SetNetworkProtocols forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkProtocolsResponse.
func Call_SetNetworkProtocols(ctx context.Context, dev *onvif.Device, request device.SetNetworkProtocols) (device.SetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkProtocolsResponse device.SetNetworkProtocolsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkProtocols")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkProtocols")
	}
}

// CallWithLogging_SetNetworkProtocols works like Call_SetNetworkProtocols but also logs the response body.
func CallWithLogging_SetNetworkProtocols(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetNetworkProtocols) (device.SetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkProtocolsResponse device.SetNetworkProtocolsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkProtocols")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetNetworkProtocols")
		return reply.Body.SetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkProtocols")
	}
}
