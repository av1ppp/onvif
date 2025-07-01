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

// Call_GetCompatibleAudioOutputConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioOutputConfigurationsResponse.
func Call_GetCompatibleAudioOutputConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleAudioOutputConfigurations) (media.GetCompatibleAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioOutputConfigurationsResponse media.GetCompatibleAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioOutputConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioOutputConfigurations")
	}
	return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, nil
}

// CallWithLogging_GetCompatibleAudioOutputConfigurations works like Call_GetCompatibleAudioOutputConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleAudioOutputConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetCompatibleAudioOutputConfigurations) (media.GetCompatibleAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioOutputConfigurationsResponse media.GetCompatibleAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioOutputConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleAudioOutputConfigurations")
	if err != nil {
		return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioOutputConfigurations")
	}
	return reply.Body.GetCompatibleAudioOutputConfigurationsResponse, nil
}
