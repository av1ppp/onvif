// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// UpgradeSystemFirmware forwards the call to onvif.Do then parses the payload of the reply as a UpgradeSystemFirmwareResponse.
func UpgradeSystemFirmware(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.UpgradeSystemFirmware]) (device.UpgradeSystemFirmwareResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UpgradeSystemFirmwareResponse device.UpgradeSystemFirmwareResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "UpgradeSystemFirmware")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.UpgradeSystemFirmwareResponse, err
}
