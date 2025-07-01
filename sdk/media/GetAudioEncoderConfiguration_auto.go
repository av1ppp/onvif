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

// Call_GetAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationResponse.
func Call_GetAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioEncoderConfiguration) (media.GetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationResponse media.GetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfiguration")
	}
	return reply.Body.GetAudioEncoderConfigurationResponse, nil
}

// CallWithLogging_GetAudioEncoderConfiguration works like Call_GetAudioEncoderConfiguration but also logs the response body.
func CallWithLogging_GetAudioEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioEncoderConfiguration) (media.GetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationResponse media.GetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioEncoderConfiguration")
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfiguration")
	}
	return reply.Body.GetAudioEncoderConfigurationResponse, nil
}
