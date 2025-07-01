// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetServiceCapabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServiceCapabilitiesResponse.
func Call_GetServiceCapabilities(ctx context.Context, dev *onvif.Device, request ptz.GetServiceCapabilities) (ptz.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse ptz.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServiceCapabilities")
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	}
}
