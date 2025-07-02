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

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetRelayOutputSettingsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRelayOutputSettings")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRelayOutputSettings")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetRelayOutputSettingsResponse, err
}
