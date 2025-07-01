// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// ModifyPresetTour forwards the call to onvif.Do then parses the payload of the reply as a ModifyPresetTourResponse.
func ModifyPresetTour(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.ModifyPresetTour]) (ptz.ModifyPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyPresetTourResponse ptz.ModifyPresetTourResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyPresetTour")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ModifyPresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.ModifyPresetTourResponse, err
}
