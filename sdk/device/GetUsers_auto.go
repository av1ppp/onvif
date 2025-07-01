// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a GetUsersResponse.
func Call_GetUsers(ctx context.Context, dev *onvif.Device, request device.GetUsers) (device.GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse device.GetUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetUsers")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetUsers")
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetUsers")
	}
}
