// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetAudioDecoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationOptionsResponse.
func Call_GetAudioDecoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfigurationOptions) (media.GetAudioDecoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationOptionsResponse media.GetAudioDecoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurationOptions")
	}
}

// CallWithLogging_GetAudioDecoderConfigurationOptions works like Call_GetAudioDecoderConfigurationOptions but also logs the response body.
func CallWithLogging_GetAudioDecoderConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioDecoderConfigurationOptions) (media.GetAudioDecoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationOptionsResponse media.GetAudioDecoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioDecoderConfigurationOptions")
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurationOptions")
	}
}
