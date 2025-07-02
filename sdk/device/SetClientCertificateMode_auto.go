// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetClientCertificateMode forwards the call to onvif.Do then parses the payload of the reply as a SetClientCertificateModeResponse.
func SetClientCertificateMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetClientCertificateMode]) (device.SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse device.SetClientCertificateModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetClientCertificateMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetClientCertificateModeResponse, err
}
