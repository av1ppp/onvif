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

// Call_SetDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDiscoveryModeResponse.
func Call_SetDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.SetDiscoveryMode) (device.SetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDiscoveryModeResponse device.SetDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDiscoveryMode")
	}
}

// CallWithLogging_SetDiscoveryMode works like Call_SetDiscoveryMode but also logs the response body.
func CallWithLogging_SetDiscoveryMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetDiscoveryMode) (device.SetDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDiscoveryModeResponse device.SetDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDiscoveryMode")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDiscoveryMode")
		return reply.Body.SetDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDiscoveryMode")
	}
}
