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

// Call_StartFirmwareUpgrade forwards the call to dev.CallMethod() then parses the payload of the reply as a StartFirmwareUpgradeResponse.
func Call_StartFirmwareUpgrade(ctx context.Context, dev *onvif.Device, request device.StartFirmwareUpgrade) (device.StartFirmwareUpgradeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartFirmwareUpgradeResponse device.StartFirmwareUpgradeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartFirmwareUpgrade")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartFirmwareUpgrade")
	}
	return reply.Body.StartFirmwareUpgradeResponse, nil
}

// CallWithLogging_StartFirmwareUpgrade works like Call_StartFirmwareUpgrade but also logs the response body.
func CallWithLogging_StartFirmwareUpgrade(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.StartFirmwareUpgrade) (device.StartFirmwareUpgradeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartFirmwareUpgradeResponse device.StartFirmwareUpgradeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartFirmwareUpgrade")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartFirmwareUpgrade")
	if err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartFirmwareUpgrade")
	}
	return reply.Body.StartFirmwareUpgradeResponse, nil
}
