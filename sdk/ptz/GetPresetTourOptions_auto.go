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

// Call_GetPresetTourOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetTourOptionsResponse.
func Call_GetPresetTourOptions(ctx context.Context, dev *onvif.Device, request ptz.GetPresetTourOptions) (ptz.GetPresetTourOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourOptionsResponse ptz.GetPresetTourOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetTourOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTourOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetPresetTourOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTourOptions")
	}
}

// CallWithLogging_GetPresetTourOptions works like Call_GetPresetTourOptions but also logs the response body.
func CallWithLogging_GetPresetTourOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetPresetTourOptions) (ptz.GetPresetTourOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourOptionsResponse ptz.GetPresetTourOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetTourOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTourOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPresetTourOptions")
		return reply.Body.GetPresetTourOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresetTourOptions")
	}
}
