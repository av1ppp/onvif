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

// Call_AddAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioEncoderConfigurationResponse.
func Call_AddAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioEncoderConfiguration) (media.AddAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioEncoderConfigurationResponse media.AddAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioEncoderConfiguration")
	}
	return reply.Body.AddAudioEncoderConfigurationResponse, nil
}

// CallWithLogging_AddAudioEncoderConfiguration works like Call_AddAudioEncoderConfiguration but also logs the response body.
func CallWithLogging_AddAudioEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddAudioEncoderConfiguration) (media.AddAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioEncoderConfigurationResponse media.AddAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddAudioEncoderConfiguration")
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioEncoderConfiguration")
	}
	return reply.Body.AddAudioEncoderConfigurationResponse, nil
}
