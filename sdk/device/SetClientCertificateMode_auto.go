// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetClientCertificateModeResponse.
func Call_SetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.SetClientCertificateMode) (device.SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse device.SetClientCertificateModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetClientCertificateMode")
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	}
}
