// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetStreamUri forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStreamUriResponse.
func Call_GetStreamUri(ctx context.Context, dev *onvif.Device, request media.GetStreamUri) (media.GetStreamUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStreamUriResponse media.GetStreamUriResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStreamUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStreamUri")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStreamUri")
		return reply.Body.GetStreamUriResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetStreamUri")
	}
}
