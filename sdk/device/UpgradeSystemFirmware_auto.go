// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	}
	return reply.Body.UpgradeSystemFirmwareResponse, nil
}

// CallWithLogging_UpgradeSystemFirmware works like Call_UpgradeSystemFirmware but also logs the response body.
func CallWithLogging_UpgradeSystemFirmware(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.UpgradeSystemFirmware) (device.UpgradeSystemFirmwareResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UpgradeSystemFirmwareResponse device.UpgradeSystemFirmwareResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "UpgradeSystemFirmware")
	if err != nil {
		return reply.Body.UpgradeSystemFirmwareResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "UpgradeSystemFirmware")
	}
	return reply.Body.UpgradeSystemFirmwareResponse, nil
}
