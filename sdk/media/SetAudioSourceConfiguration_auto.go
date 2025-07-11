// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetAudioSourceConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetAudioSourceConfigurationResponse.
func SetAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetAudioSourceConfiguration]) (media.SetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioSourceConfigurationResponse media.SetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioSourceConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetAudioSourceConfigurationResponse, err
}
