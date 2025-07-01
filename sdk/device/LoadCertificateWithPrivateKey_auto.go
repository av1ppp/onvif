// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_LoadCertificateWithPrivateKey forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCertificateWithPrivateKeyResponse.
func Call_LoadCertificateWithPrivateKey(ctx context.Context, dev *onvif.Device, request device.LoadCertificateWithPrivateKey) (device.LoadCertificateWithPrivateKeyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificateWithPrivateKeyResponse device.LoadCertificateWithPrivateKeyResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "LoadCertificateWithPrivateKey")
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	}
}
