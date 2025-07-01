// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetDot1XConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationsResponse.
func Call_GetDot1XConfigurations(ctx context.Context, dev *onvif.Device, request device.GetDot1XConfigurations) (device.GetDot1XConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationsResponse device.GetDot1XConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDot1XConfigurations")
		return reply.Body.GetDot1XConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot1XConfigurations")
	}
}
