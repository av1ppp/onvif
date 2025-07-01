// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveIPAddressFilterResponse.
func Call_RemoveIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.RemoveIPAddressFilter) (device.RemoveIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveIPAddressFilterResponse device.RemoveIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveIPAddressFilter")
		return reply.Body.RemoveIPAddressFilterResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveIPAddressFilter")
	}
}
