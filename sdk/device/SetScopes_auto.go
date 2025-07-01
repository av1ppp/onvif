// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a SetScopesResponse.
func Call_SetScopes(ctx context.Context, dev *onvif.Device, request device.SetScopes) (device.SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse device.SetScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetScopes")
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetScopes")
	}
}
