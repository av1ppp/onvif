// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_LoadCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCertificatesResponse.
func Call_LoadCertificates(ctx context.Context, dev *onvif.Device, request device.LoadCertificates) (device.LoadCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificatesResponse device.LoadCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "LoadCertificates")
		return reply.Body.LoadCertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCertificates")
	}
}
