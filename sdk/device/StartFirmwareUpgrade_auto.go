// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// StartFirmwareUpgrade forwards the call to onvif.Do then parses the payload of the reply as a StartFirmwareUpgradeResponse.
func StartFirmwareUpgrade(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.StartFirmwareUpgrade]) (device.StartFirmwareUpgradeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartFirmwareUpgradeResponse device.StartFirmwareUpgradeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartFirmwareUpgrade")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartFirmwareUpgrade")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.StartFirmwareUpgradeResponse, err
}
