// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetIPAddressFilter forwards the call to onvif.Do then parses the payload of the reply as a GetIPAddressFilterResponse.
func GetIPAddressFilter(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetIPAddressFilter]) (device.GetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetIPAddressFilterResponse device.GetIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetIPAddressFilter")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetIPAddressFilterResponse, err
}
