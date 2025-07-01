// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCompatibleAudioEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioEncoderConfigurationsResponse.
func Call_GetCompatibleAudioEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleAudioEncoderConfigurations) (media.GetCompatibleAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioEncoderConfigurationsResponse media.GetCompatibleAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleAudioEncoderConfigurations")
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	}
}
