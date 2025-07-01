// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioOutputConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetAudioOutputConfigurationOptionsResponse.
func GetAudioOutputConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioOutputConfigurationOptions]) (media.GetAudioOutputConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationOptionsResponse media.GetAudioOutputConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfigurationOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAudioOutputConfigurationOptionsResponse, err
}
