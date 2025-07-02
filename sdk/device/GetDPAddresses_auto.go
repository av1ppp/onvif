// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetDPAddresses forwards the call to onvif.Do then parses the payload of the reply as a GetDPAddressesResponse.
func GetDPAddresses(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetDPAddresses]) (device.GetDPAddressesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDPAddressesResponse device.GetDPAddressesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDPAddresses")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDPAddresses")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetDPAddressesResponse, err
}
