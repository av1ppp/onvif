// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetHostnameFromDHCP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameFromDHCPResponse.
func Call_SetHostnameFromDHCP(ctx context.Context, dev *onvif.Device, request device.SetHostnameFromDHCP) (device.SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse device.SetHostnameFromDHCPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetHostnameFromDHCP")
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	}
}
