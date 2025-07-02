// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioSourceConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetAudioSourceConfigurationOptionsResponse.
func GetAudioSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioSourceConfigurationOptions]) (media.GetAudioSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationOptionsResponse media.GetAudioSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioSourceConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAudioSourceConfigurationOptionsResponse, err
}
