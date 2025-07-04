// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetAudioEncoderConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetAudioEncoderConfigurationResponse.
func SetAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetAudioEncoderConfiguration]) (media.SetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioEncoderConfigurationResponse media.SetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioEncoderConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetAudioEncoderConfigurationResponse, err
}
