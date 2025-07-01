// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioOutputConfigurationResponse.
func Call_RemoveAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioOutputConfiguration) (media.RemoveAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioOutputConfigurationResponse media.RemoveAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioOutputConfiguration")
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioOutputConfiguration")
	}
}
