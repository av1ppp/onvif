// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// DeleteCertificates forwards the call to onvif.Do then parses the payload of the reply as a DeleteCertificatesResponse.
func DeleteCertificates(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.DeleteCertificates]) (device.DeleteCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteCertificatesResponse device.DeleteCertificatesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.DeleteCertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteCertificates")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteCertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.DeleteCertificatesResponse, err
}
