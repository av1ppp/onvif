// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetCACertificates forwards the call to onvif.Do then parses the payload of the reply as a GetCACertificatesResponse.
func GetCACertificates(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetCACertificates]) (device.GetCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCACertificatesResponse device.GetCACertificatesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCACertificates")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCACertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetCACertificatesResponse, err
}
