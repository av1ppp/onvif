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

// Call_GetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationResponse.
func Call_GetAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfiguration) (media.GetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationResponse media.GetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfiguration")
	}
}

// CallWithLogging_GetAudioDecoderConfiguration works like Call_GetAudioDecoderConfiguration but also logs the response body.
func CallWithLogging_GetAudioDecoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioDecoderConfiguration) (media.GetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationResponse media.GetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioDecoderConfiguration")
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfiguration")
	}
}
