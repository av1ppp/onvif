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

// Call_ContinuousMove forwards the call to dev.CallMethod() then parses the payload of the reply as a ContinuousMoveResponse.
func Call_ContinuousMove(ctx context.Context, dev *onvif.Device, request ptz.ContinuousMove) (ptz.ContinuousMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ContinuousMoveResponse ptz.ContinuousMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ContinuousMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ContinuousMove")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.ContinuousMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ContinuousMove")
	}
}

// CallWithLogging_ContinuousMove works like Call_ContinuousMove but also logs the response body.
func CallWithLogging_ContinuousMove(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.ContinuousMove) (ptz.ContinuousMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ContinuousMoveResponse ptz.ContinuousMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ContinuousMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ContinuousMove")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ContinuousMove")
		return reply.Body.ContinuousMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "ContinuousMove")
	}
}
