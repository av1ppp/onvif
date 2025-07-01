// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetSystemSupportInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemSupportInformationResponse.
func Call_GetSystemSupportInformation(ctx context.Context, dev *onvif.Device, request device.GetSystemSupportInformation) (device.GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse device.GetSystemSupportInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSystemSupportInformation")
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	}
}
