// Code generated : DO NOT EDIT.

package device

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

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ScanAvailableDot11Networks")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ScanAvailableDot11Networks")
	}
	return reply.Body.ScanAvailableDot11NetworksResponse, nil
}
