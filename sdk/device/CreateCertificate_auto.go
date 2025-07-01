// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_CreateCertificate forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateCertificateResponse.
func Call_CreateCertificate(ctx context.Context, dev *onvif.Device, request device.CreateCertificate) (device.CreateCertificateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateCertificateResponse device.CreateCertificateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateCertificateResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateCertificate")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateCertificate")
		return reply.Body.CreateCertificateResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateCertificate")
	}
}
