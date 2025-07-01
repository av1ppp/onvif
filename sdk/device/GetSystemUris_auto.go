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

// Call_GetSystemUris forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemUrisResponse.
func Call_GetSystemUris(ctx context.Context, dev *onvif.Device, request device.GetSystemUris) (device.GetSystemUrisResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemUrisResponse device.GetSystemUrisResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemUrisResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemUris")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetSystemUrisResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemUris")
	}
}

// CallWithLogging_GetSystemUris works like Call_GetSystemUris but also logs the response body.
func CallWithLogging_GetSystemUris(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetSystemUris) (device.GetSystemUrisResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemUrisResponse device.GetSystemUrisResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemUrisResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemUris")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemUris")
		return reply.Body.GetSystemUrisResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemUris")
	}
}
