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

// Call_GetDot11Status forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11StatusResponse.
func Call_GetDot11Status(ctx context.Context, dev *onvif.Device, request device.GetDot11Status) (device.GetDot11StatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11StatusResponse device.GetDot11StatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Status")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Status")
	}
	return reply.Body.GetDot11StatusResponse, nil
}

// CallWithLogging_GetDot11Status works like Call_GetDot11Status but also logs the response body.
func CallWithLogging_GetDot11Status(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDot11Status) (device.GetDot11StatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11StatusResponse device.GetDot11StatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Status")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDot11Status")
	if err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Status")
	}
	return reply.Body.GetDot11StatusResponse, nil
}
