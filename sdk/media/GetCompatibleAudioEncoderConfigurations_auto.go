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

// Call_GetCompatibleAudioEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioEncoderConfigurationsResponse.
func Call_GetCompatibleAudioEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleAudioEncoderConfigurations) (media.GetCompatibleAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioEncoderConfigurationsResponse media.GetCompatibleAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	}
	return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, nil
}

// CallWithLogging_GetCompatibleAudioEncoderConfigurations works like Call_GetCompatibleAudioEncoderConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleAudioEncoderConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetCompatibleAudioEncoderConfigurations) (media.GetCompatibleAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioEncoderConfigurationsResponse media.GetCompatibleAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleAudioEncoderConfigurations")
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	}
	return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, nil
}
