// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationResponse.
func Call_GetAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfiguration) (media.GetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationResponse media.GetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputConfiguration")
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	}
}
