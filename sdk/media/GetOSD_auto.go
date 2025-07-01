// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDResponse.
func Call_GetOSD(ctx context.Context, dev *onvif.Device, request media.GetOSD) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSD")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOSD")
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSD")
	}
}
