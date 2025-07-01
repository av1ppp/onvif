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

// Call_SetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioDecoderConfigurationResponse.
func Call_SetAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioDecoderConfiguration) (media.SetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioDecoderConfigurationResponse media.SetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioDecoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioDecoderConfiguration")
	}
	return reply.Body.SetAudioDecoderConfigurationResponse, nil
}

// CallWithLogging_SetAudioDecoderConfiguration works like Call_SetAudioDecoderConfiguration but also logs the response body.
func CallWithLogging_SetAudioDecoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetAudioDecoderConfiguration) (media.SetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioDecoderConfigurationResponse media.SetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioDecoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioDecoderConfiguration")
	if err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioDecoderConfiguration")
	}
	return reply.Body.SetAudioDecoderConfigurationResponse, nil
}
