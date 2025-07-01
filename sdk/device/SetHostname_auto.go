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

// Call_SetHostname forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameResponse.
func Call_SetHostname(ctx context.Context, dev *onvif.Device, request device.SetHostname) (device.SetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameResponse device.SetHostnameResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostname")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetHostnameResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHostname")
	}
	return reply.Body.SetHostnameResponse, nil
}

// CallWithLogging_SetHostname works like Call_SetHostname but also logs the response body.
func CallWithLogging_SetHostname(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetHostname) (device.SetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameResponse device.SetHostnameResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostname")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetHostname")
	if err != nil {
		return reply.Body.SetHostnameResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHostname")
	}
	return reply.Body.SetHostnameResponse, nil
}
