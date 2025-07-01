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

// Call_GetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteDiscoveryModeResponse.
func Call_GetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.GetRemoteDiscoveryMode) (device.GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse device.GetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	}
	return reply.Body.GetRemoteDiscoveryModeResponse, nil
}

// CallWithLogging_GetRemoteDiscoveryMode works like Call_GetRemoteDiscoveryMode but also logs the response body.
func CallWithLogging_GetRemoteDiscoveryMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetRemoteDiscoveryMode) (device.GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse device.GetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRemoteDiscoveryMode")
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	}
	return reply.Body.GetRemoteDiscoveryModeResponse, nil
}
