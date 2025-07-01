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

// Call_GetDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDiscoveryModeResponse.
func Call_GetDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.GetDiscoveryMode) (device.GetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDiscoveryModeResponse device.GetDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDiscoveryMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDiscoveryMode")
	}
	return reply.Body.GetDiscoveryModeResponse, nil
}

// CallWithLogging_GetDiscoveryMode works like Call_GetDiscoveryMode but also logs the response body.
func CallWithLogging_GetDiscoveryMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDiscoveryMode) (device.GetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDiscoveryModeResponse device.GetDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDiscoveryMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDiscoveryMode")
	if err != nil {
		return reply.Body.GetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDiscoveryMode")
	}
	return reply.Body.GetDiscoveryModeResponse, nil
}
