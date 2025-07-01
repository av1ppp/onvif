// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemovePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetTourResponse.
func Call_RemovePresetTour(ctx context.Context, dev *onvif.Device, request ptz.RemovePresetTour) (ptz.RemovePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetTourResponse ptz.RemovePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePresetTour")
		return reply.Body.RemovePresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePresetTour")
	}
}
