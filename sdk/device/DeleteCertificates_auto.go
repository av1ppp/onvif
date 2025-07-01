// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_DeleteCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteCertificatesResponse.
func Call_DeleteCertificates(ctx context.Context, dev *onvif.Device, request device.DeleteCertificates) (device.DeleteCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteCertificatesResponse device.DeleteCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteCertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteCertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteCertificates")
		return reply.Body.DeleteCertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteCertificates")
	}
}
