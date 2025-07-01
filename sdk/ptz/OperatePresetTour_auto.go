// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_OperatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a OperatePresetTourResponse.
func Call_OperatePresetTour(ctx context.Context, dev *onvif.Device, request ptz.OperatePresetTour) (ptz.OperatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			OperatePresetTourResponse ptz.OperatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.OperatePresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "OperatePresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "OperatePresetTour")
		return reply.Body.OperatePresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "OperatePresetTour")
	}
}
