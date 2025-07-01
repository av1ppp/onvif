// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GotoHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoHomePositionResponse.
func Call_GotoHomePosition(ctx context.Context, dev *onvif.Device, request ptz.GotoHomePosition) (ptz.GotoHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoHomePositionResponse ptz.GotoHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GotoHomePositionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GotoHomePosition")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GotoHomePosition")
		return reply.Body.GotoHomePositionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GotoHomePosition")
	}
}
