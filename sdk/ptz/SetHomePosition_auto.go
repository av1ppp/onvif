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

// Call_SetHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHomePositionResponse.
func Call_SetHomePosition(ctx context.Context, dev *onvif.Device, request ptz.SetHomePosition) (ptz.SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse ptz.SetHomePositionResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHomePosition")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHomePosition")
	}
	return reply.Body.SetHomePositionResponse, nil
}

// CallWithLogging_SetHomePosition works like Call_SetHomePosition but also logs the response body.
func CallWithLogging_SetHomePosition(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.SetHomePosition) (ptz.SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse ptz.SetHomePositionResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHomePosition")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetHomePosition")
	if err != nil {
		return reply.Body.SetHomePositionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHomePosition")
	}
	return reply.Body.SetHomePositionResponse, nil
}
