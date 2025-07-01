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

// Call_AbsoluteMove forwards the call to dev.CallMethod() then parses the payload of the reply as a AbsoluteMoveResponse.
func Call_AbsoluteMove(ctx context.Context, dev *onvif.Device, request ptz.AbsoluteMove) (ptz.AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse ptz.AbsoluteMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AbsoluteMove")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AbsoluteMove")
	}
	return reply.Body.AbsoluteMoveResponse, nil
}

// CallWithLogging_AbsoluteMove works like Call_AbsoluteMove but also logs the response body.
func CallWithLogging_AbsoluteMove(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.AbsoluteMove) (ptz.AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse ptz.AbsoluteMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AbsoluteMove")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AbsoluteMove")
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AbsoluteMove")
	}
	return reply.Body.AbsoluteMoveResponse, nil
}
