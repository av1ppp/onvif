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

// Call_GetPkcs10Request forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPkcs10RequestResponse.
func Call_GetPkcs10Request(ctx context.Context, dev *onvif.Device, request device.GetPkcs10Request) (device.GetPkcs10RequestResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPkcs10RequestResponse device.GetPkcs10RequestResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPkcs10Request")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPkcs10Request")
	}
	return reply.Body.GetPkcs10RequestResponse, nil
}

// CallWithLogging_GetPkcs10Request works like Call_GetPkcs10Request but also logs the response body.
func CallWithLogging_GetPkcs10Request(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetPkcs10Request) (device.GetPkcs10RequestResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPkcs10RequestResponse device.GetPkcs10RequestResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPkcs10Request")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPkcs10Request")
	if err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPkcs10Request")
	}
	return reply.Body.GetPkcs10RequestResponse, nil
}
