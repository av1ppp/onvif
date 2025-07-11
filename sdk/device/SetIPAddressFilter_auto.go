// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetIPAddressFilter forwards the call to onvif.Do then parses the payload of the reply as a SetIPAddressFilterResponse.
func SetIPAddressFilter(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetIPAddressFilter]) (device.SetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetIPAddressFilterResponse device.SetIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetIPAddressFilter")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetIPAddressFilterResponse, err
}
