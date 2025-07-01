// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_ScanAvailableDot11Networks forwards the call to dev.CallMethod() then parses the payload of the reply as a ScanAvailableDot11NetworksResponse.
func Call_ScanAvailableDot11Networks(ctx context.Context, dev *onvif.Device, request device.ScanAvailableDot11Networks) (device.ScanAvailableDot11NetworksResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ScanAvailableDot11NetworksResponse device.ScanAvailableDot11NetworksResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ScanAvailableDot11Networks")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ScanAvailableDot11Networks")
		return reply.Body.ScanAvailableDot11NetworksResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ScanAvailableDot11Networks")
	}
}
