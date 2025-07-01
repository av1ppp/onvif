// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetAudioDecoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationsResponse.
func Call_GetAudioDecoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfigurations) (media.GetAudioDecoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationsResponse media.GetAudioDecoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioDecoderConfigurations")
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	}
}
