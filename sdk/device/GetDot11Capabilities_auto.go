// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetDot11Capabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11CapabilitiesResponse.
func Call_GetDot11Capabilities(ctx context.Context, dev *onvif.Device, request device.GetDot11Capabilities) (device.GetDot11CapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11CapabilitiesResponse device.GetDot11CapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDot11Capabilities")
		return reply.Body.GetDot11CapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Capabilities")
	}
}
