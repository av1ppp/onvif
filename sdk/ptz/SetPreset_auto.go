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

// Call_SetPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a SetPresetResponse.
func Call_SetPreset(ctx context.Context, dev *onvif.Device, request ptz.SetPreset) (ptz.SetPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetPresetResponse ptz.SetPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetPreset")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetPreset")
	}
	return reply.Body.SetPresetResponse, nil
}

// CallWithLogging_SetPreset works like Call_SetPreset but also logs the response body.
func CallWithLogging_SetPreset(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.SetPreset) (ptz.SetPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetPresetResponse ptz.SetPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetPreset")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetPreset")
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetPreset")
	}
	return reply.Body.SetPresetResponse, nil
}
