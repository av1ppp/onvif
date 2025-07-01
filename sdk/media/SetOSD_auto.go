// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a SetOSDResponse.
func Call_SetOSD(ctx context.Context, dev *onvif.Device, request media.SetOSD) (media.SetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetOSDResponse media.SetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetOSD")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetOSD")
		return reply.Body.SetOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetOSD")
	}
}
