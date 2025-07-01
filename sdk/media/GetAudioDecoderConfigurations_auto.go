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

// Call_GetAudioDecoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationsResponse.
func Call_GetAudioDecoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfigurations) (media.GetAudioDecoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationsResponse media.GetAudioDecoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	}
	return reply.Body.GetAudioDecoderConfigurationsResponse, nil
}

// CallWithLogging_GetAudioDecoderConfigurations works like Call_GetAudioDecoderConfigurations but also logs the response body.
func CallWithLogging_GetAudioDecoderConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioDecoderConfigurations) (media.GetAudioDecoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationsResponse media.GetAudioDecoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioDecoderConfigurations")
	if err != nil {
		return reply.Body.GetAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioDecoderConfigurations")
	}
	return reply.Body.GetAudioDecoderConfigurationsResponse, nil
}
