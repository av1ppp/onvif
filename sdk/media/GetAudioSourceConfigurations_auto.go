// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioSourceConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetAudioSourceConfigurationsResponse.
func GetAudioSourceConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioSourceConfigurations]) (media.GetAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationsResponse media.GetAudioSourceConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAudioSourceConfigurationsResponse, err
}
