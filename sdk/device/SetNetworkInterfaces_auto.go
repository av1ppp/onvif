// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetNetworkInterfaces forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkInterfacesResponse.
func Call_SetNetworkInterfaces(ctx context.Context, dev *onvif.Device, request device.SetNetworkInterfaces) (device.SetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkInterfacesResponse device.SetNetworkInterfacesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetNetworkInterfaces")
		return reply.Body.SetNetworkInterfacesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNetworkInterfaces")
	}
}
