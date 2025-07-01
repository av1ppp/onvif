// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetCompatibleAudioSourceConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetCompatibleAudioSourceConfigurationsResponse.
func GetCompatibleAudioSourceConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetCompatibleAudioSourceConfigurations]) (media.GetCompatibleAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioSourceConfigurationsResponse media.GetCompatibleAudioSourceConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioSourceConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleAudioSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, err
}
