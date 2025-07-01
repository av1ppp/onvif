// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// AddAudioEncoderConfiguration forwards the call to onvif.Do then parses the payload of the reply as a AddAudioEncoderConfigurationResponse.
func AddAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.AddAudioEncoderConfiguration]) (media.AddAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioEncoderConfigurationResponse media.AddAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioEncoderConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddAudioEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AddAudioEncoderConfigurationResponse, err
}
