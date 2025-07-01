// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// LoadCertificates forwards the call to onvif.Do then parses the payload of the reply as a LoadCertificatesResponse.
func LoadCertificates(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.LoadCertificates]) (device.LoadCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificatesResponse device.LoadCertificatesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.LoadCertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificates")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "LoadCertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.LoadCertificatesResponse, err
}
