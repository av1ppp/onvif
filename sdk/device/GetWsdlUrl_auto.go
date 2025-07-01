// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetWsdlUrl forwards the call to dev.CallMethod() then parses the payload of the reply as a GetWsdlUrlResponse.
func Call_GetWsdlUrl(ctx context.Context, dev *onvif.Device, request device.GetWsdlUrl) (device.GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse device.GetWsdlUrlResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetWsdlUrl")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetWsdlUrl")
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetWsdlUrl")
	}
}
