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

// Call_CreatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePresetTourResponse.
func Call_CreatePresetTour(ctx context.Context, dev *onvif.Device, request ptz.CreatePresetTour) (ptz.CreatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePresetTourResponse ptz.CreatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreatePresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.CreatePresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreatePresetTour")
	}
}

// CallWithLogging_CreatePresetTour works like Call_CreatePresetTour but also logs the response body.
func CallWithLogging_CreatePresetTour(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.CreatePresetTour) (ptz.CreatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePresetTourResponse ptz.CreatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreatePresetTour")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreatePresetTour")
		return reply.Body.CreatePresetTourResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreatePresetTour")
	}
}
