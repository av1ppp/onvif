// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetAudioOutputConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationsResponse.
func Call_GetAudioOutputConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfigurations) (media.GetAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationsResponse media.GetAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputConfigurations")
		return reply.Body.GetAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfigurations")
	}
}
