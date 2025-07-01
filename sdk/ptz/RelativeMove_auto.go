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

// Call_RelativeMove forwards the call to dev.CallMethod() then parses the payload of the reply as a RelativeMoveResponse.
func Call_RelativeMove(ctx context.Context, dev *onvif.Device, request ptz.RelativeMove) (ptz.RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse ptz.RelativeMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RelativeMove")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RelativeMove")
	}
	return reply.Body.RelativeMoveResponse, nil
}

// CallWithLogging_RelativeMove works like Call_RelativeMove but also logs the response body.
func CallWithLogging_RelativeMove(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.RelativeMove) (ptz.RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse ptz.RelativeMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RelativeMove")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RelativeMove")
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RelativeMove")
	}
	return reply.Body.RelativeMoveResponse, nil
}
