// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetCompatibleAudioOutputConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetCompatibleAudioOutputConfigurationsResponse.
func GetCompatibleAudioOutputConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetCompatibleAudioOutputConfigurations]) (media.GetCompatibleAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioOutputConfigurationsResponse media.GetCompatibleAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioOutputConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleAudioOutputConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, err
}
