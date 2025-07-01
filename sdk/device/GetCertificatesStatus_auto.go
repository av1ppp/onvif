// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetCertificatesStatus forwards the call to onvif.Do then parses the payload of the reply as a GetCertificatesStatusResponse.
func GetCertificatesStatus(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetCertificatesStatus]) (device.GetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesStatusResponse device.GetCertificatesStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCertificatesStatus")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCertificatesStatus")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetCertificatesStatusResponse, err
}
