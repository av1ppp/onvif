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

// Call_GetAudioEncoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationOptionsResponse.
func Call_GetAudioEncoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioEncoderConfigurationOptions) (media.GetAudioEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationOptionsResponse media.GetAudioEncoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurationOptions")
	}
}

// CallWithLogging_GetAudioEncoderConfigurationOptions works like Call_GetAudioEncoderConfigurationOptions but also logs the response body.
func CallWithLogging_GetAudioEncoderConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioEncoderConfigurationOptions) (media.GetAudioEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationOptionsResponse media.GetAudioEncoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioEncoderConfigurationOptions")
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurationOptions")
	}
}
