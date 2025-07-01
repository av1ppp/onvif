// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetPresetTours forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetToursResponse.
func Call_GetPresetTours(ctx context.Context, dev *onvif.Device, request ptz.GetPresetTours) (ptz.GetPresetToursResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetToursResponse ptz.GetPresetToursResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTours")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPresetTours")
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTours")
	}
}
