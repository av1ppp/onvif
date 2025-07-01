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

// Call_GetHostname forwards the call to dev.CallMethod() then parses the payload of the reply as a GetHostnameResponse.
func Call_GetHostname(ctx context.Context, dev *onvif.Device, request device.GetHostname) (device.GetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetHostnameResponse device.GetHostnameResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetHostname")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetHostnameResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetHostname")
	}
}

// CallWithLogging_GetHostname works like Call_GetHostname but also logs the response body.
func CallWithLogging_GetHostname(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetHostname) (device.GetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetHostnameResponse device.GetHostnameResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetHostname")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetHostname")
		return reply.Body.GetHostnameResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetHostname")
	}
}
