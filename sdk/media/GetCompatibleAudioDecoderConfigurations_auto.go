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

// Call_GetCompatibleAudioDecoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioDecoderConfigurationsResponse.
func Call_GetCompatibleAudioDecoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleAudioDecoderConfigurations) (media.GetCompatibleAudioDecoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioDecoderConfigurationsResponse media.GetCompatibleAudioDecoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioDecoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetCompatibleAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioDecoderConfigurations")
	}
}

// CallWithLogging_GetCompatibleAudioDecoderConfigurations works like Call_GetCompatibleAudioDecoderConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleAudioDecoderConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetCompatibleAudioDecoderConfigurations) (media.GetCompatibleAudioDecoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioDecoderConfigurationsResponse media.GetCompatibleAudioDecoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioDecoderConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleAudioDecoderConfigurations")
		return reply.Body.GetCompatibleAudioDecoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioDecoderConfigurations")
	}
}
