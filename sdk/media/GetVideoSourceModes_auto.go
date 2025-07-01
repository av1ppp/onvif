// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetVideoSourceModes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceModesResponse.
func Call_GetVideoSourceModes(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceModes) (media.GetVideoSourceModesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceModesResponse media.GetVideoSourceModesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceModesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceModes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceModes")
		return reply.Body.GetVideoSourceModesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceModes")
	}
}
