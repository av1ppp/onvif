// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetServices forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServicesResponse.
func Call_GetServices(ctx context.Context, dev *onvif.Device, request device.GetServices) (device.GetServicesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServicesResponse device.GetServicesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServicesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServices")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServices")
		return reply.Body.GetServicesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetServices")
	}
}
