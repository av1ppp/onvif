// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioOutputConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetAudioOutputConfigurationsResponse.
func GetAudioOutputConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioOutputConfigurations]) (media.GetAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationsResponse media.GetAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAudioOutputConfigurationsResponse, err
}
