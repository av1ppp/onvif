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

// Call_ModifyPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a ModifyPresetTourResponse.
func Call_ModifyPresetTour(ctx context.Context, dev *onvif.Device, request ptz.ModifyPresetTour) (ptz.ModifyPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyPresetTourResponse ptz.ModifyPresetTourResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyPresetTour")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ModifyPresetTour")
	}
	return reply.Body.ModifyPresetTourResponse, nil
}

// CallWithLogging_ModifyPresetTour works like Call_ModifyPresetTour but also logs the response body.
func CallWithLogging_ModifyPresetTour(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.ModifyPresetTour) (ptz.ModifyPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyPresetTourResponse ptz.ModifyPresetTourResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyPresetTour")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ModifyPresetTour")
	if err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ModifyPresetTour")
	}
	return reply.Body.ModifyPresetTourResponse, nil
}
