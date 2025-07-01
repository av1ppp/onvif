// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// DeleteGeoLocation forwards the call to onvif.Do then parses the payload of the reply as a DeleteGeoLocationResponse.
func DeleteGeoLocation(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.DeleteGeoLocation]) (device.DeleteGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteGeoLocationResponse device.DeleteGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteGeoLocation")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteGeoLocation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.DeleteGeoLocationResponse, err
}
