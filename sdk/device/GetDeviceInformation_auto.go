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

// Call_GetDeviceInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDeviceInformationResponse.
func Call_GetDeviceInformation(ctx context.Context, dev *onvif.Device, request device.GetDeviceInformation) (device.GetDeviceInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDeviceInformationResponse device.GetDeviceInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDeviceInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDeviceInformation")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetDeviceInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDeviceInformation")
	}
}

// CallWithLogging_GetDeviceInformation works like Call_GetDeviceInformation but also logs the response body.
func CallWithLogging_GetDeviceInformation(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDeviceInformation) (device.GetDeviceInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDeviceInformationResponse device.GetDeviceInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDeviceInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDeviceInformation")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDeviceInformation")
		return reply.Body.GetDeviceInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDeviceInformation")
	}
}
