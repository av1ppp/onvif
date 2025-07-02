// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// LoadCertificateWithPrivateKey forwards the call to onvif.Do then parses the payload of the reply as a LoadCertificateWithPrivateKeyResponse.
func LoadCertificateWithPrivateKey(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.LoadCertificateWithPrivateKey]) (device.LoadCertificateWithPrivateKeyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificateWithPrivateKeyResponse device.LoadCertificateWithPrivateKeyResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "LoadCertificateWithPrivateKey")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.LoadCertificateWithPrivateKeyResponse, err
}
