// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetRelayOutputState forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRelayOutputStateResponse.
func Call_SetRelayOutputState(ctx context.Context, dev *onvif.Device, request device.SetRelayOutputState) (device.SetRelayOutputStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputStateResponse device.SetRelayOutputStateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputState")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRelayOutputState")
		return reply.Body.SetRelayOutputStateResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRelayOutputState")
	}
}
