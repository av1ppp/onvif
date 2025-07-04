// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// ScanAvailableDot11Networks forwards the call to onvif.Do then parses the payload of the reply as a ScanAvailableDot11NetworksResponse.
func ScanAvailableDot11Networks(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.ScanAvailableDot11Networks]) (device.ScanAvailableDot11NetworksResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ScanAvailableDot11NetworksResponse device.ScanAvailableDot11NetworksResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ScanAvailableDot11Networks")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ScanAvailableDot11Networks")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.ScanAvailableDot11NetworksResponse, err
}
