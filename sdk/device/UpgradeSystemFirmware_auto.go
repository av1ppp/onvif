// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_UpgradeSystemFirmware forwards the call to dev.CallMethod() then parses the payload of the reply as a UpgradeSystemFirmwareResponse.
func Call_UpgradeSystemFirmware(ctx context.Context, dev *onvif.Device, request device.UpgradeSystemFirmware) (device.UpgradeSystemFirmwareResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UpgradeSystemFirmwareResponse device.UpgradeSystemFirmwareResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "UpgradeSystemFirmware")
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	}
}
