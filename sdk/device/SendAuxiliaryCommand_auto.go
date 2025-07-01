// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SendAuxiliaryCommand forwards the call to dev.CallMethod() then parses the payload of the reply as a SendAuxiliaryCommandResponse.
func Call_SendAuxiliaryCommand(ctx context.Context, dev *onvif.Device, request device.SendAuxiliaryCommand) (device.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse device.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}
}

// CallWithLogging_SendAuxiliaryCommand works like Call_SendAuxiliaryCommand but also logs the response body.
func CallWithLogging_SendAuxiliaryCommand(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SendAuxiliaryCommand) (device.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse device.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SendAuxiliaryCommand")
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}
}
