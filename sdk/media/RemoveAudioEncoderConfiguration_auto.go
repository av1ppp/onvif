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

// Call_RemoveAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioEncoderConfigurationResponse.
func Call_RemoveAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioEncoderConfiguration) (media.RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse media.RemoveAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	}
	return reply.Body.RemoveAudioEncoderConfigurationResponse, nil
}

// CallWithLogging_RemoveAudioEncoderConfiguration works like Call_RemoveAudioEncoderConfiguration but also logs the response body.
func CallWithLogging_RemoveAudioEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemoveAudioEncoderConfiguration) (media.RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse media.RemoveAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveAudioEncoderConfiguration")
	if err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioEncoderConfiguration")
	}
	return reply.Body.RemoveAudioEncoderConfigurationResponse, nil
}
