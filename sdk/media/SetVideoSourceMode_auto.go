// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetVideoSourceMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceModeResponse.
func Call_SetVideoSourceMode(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceMode) (media.SetVideoSourceModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceModeResponse media.SetVideoSourceModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoSourceMode")
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	}
}
