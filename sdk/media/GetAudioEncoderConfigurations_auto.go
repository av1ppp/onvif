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

// Call_GetAudioEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationsResponse.
func Call_GetAudioEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioEncoderConfigurations) (media.GetAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationsResponse media.GetAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurations")
	}
	return reply.Body.GetAudioEncoderConfigurationsResponse, nil
}

// CallWithLogging_GetAudioEncoderConfigurations works like Call_GetAudioEncoderConfigurations but also logs the response body.
func CallWithLogging_GetAudioEncoderConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioEncoderConfigurations) (media.GetAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationsResponse media.GetAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioEncoderConfigurations")
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioEncoderConfigurations")
	}
	return reply.Body.GetAudioEncoderConfigurationsResponse, nil
}
