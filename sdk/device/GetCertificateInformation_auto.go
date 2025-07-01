// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCertificateInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificateInformationResponse.
func Call_GetCertificateInformation(ctx context.Context, dev *onvif.Device, request device.GetCertificateInformation) (device.GetCertificateInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificateInformationResponse device.GetCertificateInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificateInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCertificateInformation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCertificateInformation")
		return reply.Body.GetCertificateInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCertificateInformation")
	}
}
