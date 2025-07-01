// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_ModifyPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a ModifyPresetTourResponse.
func Call_ModifyPresetTour(ctx context.Context, dev *onvif.Device, request ptz.ModifyPresetTour) (ptz.ModifyPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyPresetTourResponse ptz.ModifyPresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyPresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "ModifyPresetTour")
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ModifyPresetTour")
	}
}
