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

// Call_GetWsdlUrl forwards the call to dev.CallMethod() then parses the payload of the reply as a GetWsdlUrlResponse.
func Call_GetWsdlUrl(ctx context.Context, dev *onvif.Device, request device.GetWsdlUrl) (device.GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse device.GetWsdlUrlResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetWsdlUrl")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetWsdlUrl")
	}
	return reply.Body.GetWsdlUrlResponse, nil
}

// CallWithLogging_GetWsdlUrl works like Call_GetWsdlUrl but also logs the response body.
func CallWithLogging_GetWsdlUrl(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetWsdlUrl) (device.GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse device.GetWsdlUrlResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetWsdlUrl")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetWsdlUrl")
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetWsdlUrl")
	}
	return reply.Body.GetWsdlUrlResponse, nil
}
