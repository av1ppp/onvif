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

// Call_GotoPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoPresetResponse.
func Call_GotoPreset(ctx context.Context, dev *onvif.Device, request ptz.GotoPreset) (ptz.GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse ptz.GotoPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GotoPreset")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GotoPreset")
	}
	return reply.Body.GotoPresetResponse, nil
}

// CallWithLogging_GotoPreset works like Call_GotoPreset but also logs the response body.
func CallWithLogging_GotoPreset(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GotoPreset) (ptz.GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse ptz.GotoPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GotoPreset")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GotoPreset")
	if err != nil {
		return reply.Body.GotoPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GotoPreset")
	}
	return reply.Body.GotoPresetResponse, nil
}
