// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetZeroConfigurationResponse.
func Call_GetZeroConfiguration(ctx context.Context, dev *onvif.Device, request device.GetZeroConfiguration) (device.GetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetZeroConfigurationResponse device.GetZeroConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetZeroConfiguration")
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	}
}
