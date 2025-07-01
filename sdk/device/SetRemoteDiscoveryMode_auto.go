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

// Call_SetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteDiscoveryModeResponse.
func Call_SetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.SetRemoteDiscoveryMode) (device.SetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteDiscoveryModeResponse device.SetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteDiscoveryMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRemoteDiscoveryMode")
	}
	return reply.Body.SetRemoteDiscoveryModeResponse, nil
}

// CallWithLogging_SetRemoteDiscoveryMode works like Call_SetRemoteDiscoveryMode but also logs the response body.
func CallWithLogging_SetRemoteDiscoveryMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetRemoteDiscoveryMode) (device.SetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteDiscoveryModeResponse device.SetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetRemoteDiscoveryMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetRemoteDiscoveryMode")
	if err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetRemoteDiscoveryMode")
	}
	return reply.Body.SetRemoteDiscoveryModeResponse, nil
}
