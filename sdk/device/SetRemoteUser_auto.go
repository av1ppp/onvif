// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteUserResponse.
func Call_SetRemoteUser(ctx context.Context, dev *onvif.Device, request device.SetRemoteUser) (device.SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse device.SetRemoteUserResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteUser")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRemoteUser")
		return reply.Body.SetRemoteUserResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRemoteUser")
	}
}
