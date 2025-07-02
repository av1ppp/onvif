// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetGeoLocation forwards the call to onvif.Do then parses the payload of the reply as a SetGeoLocationResponse.
func SetGeoLocation(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetGeoLocation]) (device.SetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetGeoLocationResponse device.SetGeoLocationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetGeoLocation")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetGeoLocation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetGeoLocationResponse, err
}
