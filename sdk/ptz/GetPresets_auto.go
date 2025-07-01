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

// Call_GetPresets forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetsResponse.
func Call_GetPresets(ctx context.Context, dev *onvif.Device, request ptz.GetPresets) (ptz.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse ptz.GetPresetsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresets")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetPresetsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresets")
	}
}

// CallWithLogging_GetPresets works like Call_GetPresets but also logs the response body.
func CallWithLogging_GetPresets(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetPresets) (ptz.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse ptz.GetPresetsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresets")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPresets")
		return reply.Body.GetPresetsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetPresets")
	}
}
