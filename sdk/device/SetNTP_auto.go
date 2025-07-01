// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNTPResponse.
func Call_SetNTP(ctx context.Context, dev *onvif.Device, request device.SetNTP) (device.SetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNTPResponse device.SetNTPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetNTP")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetNTP")
		return reply.Body.SetNTPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetNTP")
	}
}
