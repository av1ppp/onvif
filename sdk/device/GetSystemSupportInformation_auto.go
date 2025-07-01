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

// Call_GetSystemSupportInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemSupportInformationResponse.
func Call_GetSystemSupportInformation(ctx context.Context, dev *onvif.Device, request device.GetSystemSupportInformation) (device.GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse device.GetSystemSupportInformationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	}
	return reply.Body.GetSystemSupportInformationResponse, nil
}

// CallWithLogging_GetSystemSupportInformation works like Call_GetSystemSupportInformation but also logs the response body.
func CallWithLogging_GetSystemSupportInformation(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetSystemSupportInformation) (device.GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse device.GetSystemSupportInformationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemSupportInformation")
	if err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemSupportInformation")
	}
	return reply.Body.GetSystemSupportInformationResponse, nil
}
