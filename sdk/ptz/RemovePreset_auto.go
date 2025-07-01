// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemovePreset forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetResponse.
func Call_RemovePreset(ctx context.Context, dev *onvif.Device, request ptz.RemovePreset) (ptz.RemovePresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetResponse ptz.RemovePresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePreset")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.RemovePresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePreset")
	}
}

// CallWithLogging_RemovePreset works like Call_RemovePreset but also logs the response body.
func CallWithLogging_RemovePreset(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.RemovePreset) (ptz.RemovePresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetResponse ptz.RemovePresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePreset")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemovePreset")
		return reply.Body.RemovePresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePreset")
	}
}
