// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_Stop forwards the call to dev.CallMethod() then parses the payload of the reply as a StopResponse.
func Call_Stop(ctx context.Context, dev *onvif.Device, request ptz.Stop) (ptz.StopResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopResponse ptz.StopResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Stop")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Stop")
		return reply.Body.StopResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Stop")
	}
}
