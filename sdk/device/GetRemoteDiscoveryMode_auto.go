// Code generated : DO NOT EDIT.

package device

import (
	"context"

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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRemoteDiscoveryMode")
		return reply.Body.GetRemoteDiscoveryModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetRemoteDiscoveryMode")
	}
}
