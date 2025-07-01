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

// Call_SendAuxiliaryCommand forwards the call to dev.CallMethod() then parses the payload of the reply as a SendAuxiliaryCommandResponse.
func Call_SendAuxiliaryCommand(ctx context.Context, dev *onvif.Device, request ptz.SendAuxiliaryCommand) (ptz.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse ptz.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}
	return reply.Body.SendAuxiliaryCommandResponse, nil
}

// CallWithLogging_SendAuxiliaryCommand works like Call_SendAuxiliaryCommand but also logs the response body.
func CallWithLogging_SendAuxiliaryCommand(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.SendAuxiliaryCommand) (ptz.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse ptz.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SendAuxiliaryCommand")
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}
	return reply.Body.SendAuxiliaryCommandResponse, nil
}
