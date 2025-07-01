// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SendAuxiliaryCommand forwards the call to onvif.Do then parses the payload of the reply as a SendAuxiliaryCommandResponse.
func SendAuxiliaryCommand(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SendAuxiliaryCommand]) (device.SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse device.SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SendAuxiliaryCommand")
	}
	return reply.Body.SendAuxiliaryCommandResponse, nil
}
