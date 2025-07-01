// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_CreateUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateUsersResponse.
func Call_CreateUsers(ctx context.Context, dev *onvif.Device, request device.CreateUsers) (device.CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse device.CreateUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateUsers")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateUsers")
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateUsers")
	}
}
