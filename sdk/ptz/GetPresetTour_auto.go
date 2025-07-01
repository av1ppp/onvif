// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetTourResponse.
func Call_GetPresetTour(ctx context.Context, dev *onvif.Device, request ptz.GetPresetTour) (ptz.GetPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourResponse ptz.GetPresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresetTour")
		return reply.Body.GetPresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTour")
	}
}
