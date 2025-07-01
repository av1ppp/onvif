// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHomePositionResponse.
func Call_SetHomePosition(ctx context.Context, dev *onvif.Device, request ptz.SetHomePosition) (ptz.SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse ptz.SetHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHomePosition")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetHomePosition")
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHomePosition")
	}
}
