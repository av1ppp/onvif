// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetZeroConfiguration forwards the call to onvif.Do then parses the payload of the reply as a GetZeroConfigurationResponse.
func GetZeroConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetZeroConfiguration]) (device.GetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetZeroConfigurationResponse device.GetZeroConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetZeroConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetZeroConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetZeroConfigurationResponse, err
}
