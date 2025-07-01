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

// Call_GetPresetTours forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetToursResponse.
func Call_GetPresetTours(ctx context.Context, dev *onvif.Device, request ptz.GetPresetTours) (ptz.GetPresetToursResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetToursResponse ptz.GetPresetToursResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTours")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTours")
	}
	return reply.Body.GetPresetToursResponse, nil
}

// CallWithLogging_GetPresetTours works like Call_GetPresetTours but also logs the response body.
func CallWithLogging_GetPresetTours(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetPresetTours) (ptz.GetPresetToursResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetToursResponse ptz.GetPresetToursResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTours")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPresetTours")
	if err != nil {
		return reply.Body.GetPresetToursResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTours")
	}
	return reply.Body.GetPresetToursResponse, nil
}
