// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// SendAuxiliaryCommand forwards the call to onvif.Do then parses the payload of the reply as a SendAuxiliaryCommandResponse.
func SendAuxiliaryCommand(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.SendAuxiliaryCommand]) (ptz.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse ptz.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SendAuxiliaryCommand")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SendAuxiliaryCommandResponse, err
}
