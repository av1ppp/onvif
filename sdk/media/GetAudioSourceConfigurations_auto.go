// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetAudioSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationsResponse.
func Call_GetAudioSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioSourceConfigurations) (media.GetAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationsResponse media.GetAudioSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioSourceConfigurations")
		return reply.Body.GetAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfigurations")
	}
}
