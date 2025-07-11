// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// AddIPAddressFilter forwards the call to onvif.Do then parses the payload of the reply as a AddIPAddressFilterResponse.
func AddIPAddressFilter(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.AddIPAddressFilter]) (device.AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse device.AddIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddIPAddressFilter")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AddIPAddressFilterResponse, err
}
