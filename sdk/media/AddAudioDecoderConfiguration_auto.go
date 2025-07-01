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

// Call_AddAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioDecoderConfigurationResponse.
func Call_AddAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioDecoderConfiguration) (media.AddAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioDecoderConfigurationResponse media.AddAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioDecoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioDecoderConfiguration")
	}
	return reply.Body.AddAudioDecoderConfigurationResponse, nil
}

// CallWithLogging_AddAudioDecoderConfiguration works like Call_AddAudioDecoderConfiguration but also logs the response body.
func CallWithLogging_AddAudioDecoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddAudioDecoderConfiguration) (media.AddAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioDecoderConfigurationResponse media.AddAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioDecoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddAudioDecoderConfiguration")
	if err != nil {
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioDecoderConfiguration")
	}
	return reply.Body.AddAudioDecoderConfigurationResponse, nil
}
