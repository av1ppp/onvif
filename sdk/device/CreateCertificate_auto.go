// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// CreateCertificate forwards the call to onvif.Do then parses the payload of the reply as a CreateCertificateResponse.
func CreateCertificate(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.CreateCertificate]) (device.CreateCertificateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateCertificateResponse device.CreateCertificateResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.CreateCertificateResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateCertificate")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateCertificate")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateCertificateResponse, err
}
