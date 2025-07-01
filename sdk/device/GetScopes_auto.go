// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetScopesResponse.
func Call_GetScopes(ctx context.Context, dev *onvif.Device, request device.GetScopes) (device.GetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetScopesResponse device.GetScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetScopes")
		return reply.Body.GetScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetScopes")
	}
}
