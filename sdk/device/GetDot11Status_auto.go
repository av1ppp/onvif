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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Status")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Status")
	}
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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDot11Status")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDot11Status")
		return reply.Body.GetDot11StatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDot11Status")
	}
}
