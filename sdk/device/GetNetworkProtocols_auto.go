// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetNetworkProtocols forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkProtocolsResponse.
func Call_GetNetworkProtocols(ctx context.Context, dev *onvif.Device, request device.GetNetworkProtocols) (device.GetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkProtocolsResponse device.GetNetworkProtocolsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNetworkProtocols")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNetworkProtocols")
		return reply.Body.GetNetworkProtocolsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNetworkProtocols")
	}
}
