// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioEncoderConfigurationResponse.
func Call_RemoveAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioEncoderConfiguration) (media.RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse media.RemoveAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioEncoderConfiguration")
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	}
}
