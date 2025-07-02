// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// RemoveIPAddressFilter forwards the call to onvif.Do then parses the payload of the reply as a RemoveIPAddressFilterResponse.
func RemoveIPAddressFilter(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.RemoveIPAddressFilter]) (device.RemoveIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveIPAddressFilterResponse device.RemoveIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.RemoveIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveIPAddressFilter")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RemoveIPAddressFilterResponse, err
}
