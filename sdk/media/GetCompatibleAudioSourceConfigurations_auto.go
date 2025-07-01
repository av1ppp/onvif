// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCompatibleAudioSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioSourceConfigurationsResponse.
func Call_GetCompatibleAudioSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleAudioSourceConfigurations) (media.GetCompatibleAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioSourceConfigurationsResponse media.GetCompatibleAudioSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleAudioSourceConfigurations")
		return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioSourceConfigurations")
	}
}
