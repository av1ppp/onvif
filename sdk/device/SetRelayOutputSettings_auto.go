// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetRelayOutputSettings forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRelayOutputSettingsResponse.
func Call_SetRelayOutputSettings(ctx context.Context, dev *onvif.Device, request device.SetRelayOutputSettings) (device.SetRelayOutputSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputSettingsResponse device.SetRelayOutputSettingsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRelayOutputSettingsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputSettings")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRelayOutputSettings")
		return reply.Body.SetRelayOutputSettingsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRelayOutputSettings")
	}
}
