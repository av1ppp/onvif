// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceConfigurationResponse.
func Call_SetVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceConfiguration) (media.SetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceConfigurationResponse media.SetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoSourceConfiguration")
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	}
}
