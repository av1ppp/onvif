// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GotoPreset forwards the call to onvif.Do then parses the payload of the reply as a GotoPresetResponse.
func GotoPreset(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GotoPreset]) (ptz.GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse ptz.GotoPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GotoPreset")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GotoPreset")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GotoPreset")
	}
	return reply.Body.GotoPresetResponse, nil
}
