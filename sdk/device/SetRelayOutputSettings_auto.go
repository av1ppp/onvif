// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetRelayOutputSettings forwards the call to onvif.Do then parses the payload of the reply as a SetRelayOutputSettingsResponse.
func SetRelayOutputSettings(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetRelayOutputSettings]) (device.SetRelayOutputSettingsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRelayOutputSettingsResponse device.SetRelayOutputSettingsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetRelayOutputSettingsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputSettings")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetRelayOutputSettingsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRelayOutputSettings")
	}
	return reply.Body.SetRelayOutputSettingsResponse, nil
}
