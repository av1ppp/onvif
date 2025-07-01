// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesResponse.
func Call_GetCertificates(ctx context.Context, dev *onvif.Device, request device.GetCertificates) (device.GetCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesResponse device.GetCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCertificates")
		return reply.Body.GetCertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCertificates")
	}
}
