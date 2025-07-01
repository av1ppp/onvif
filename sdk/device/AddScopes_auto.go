// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_AddScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a AddScopesResponse.
func Call_AddScopes(ctx context.Context, dev *onvif.Device, request device.AddScopes) (device.AddScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddScopesResponse device.AddScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddScopes")
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddScopes")
	}
}
