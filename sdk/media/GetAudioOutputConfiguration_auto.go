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

// Call_GetAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationResponse.
func Call_GetAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfiguration) (media.GetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationResponse media.GetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	}
	return reply.Body.GetAudioOutputConfigurationResponse, nil
}

// CallWithLogging_GetAudioOutputConfiguration works like Call_GetAudioOutputConfiguration but also logs the response body.
func CallWithLogging_GetAudioOutputConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioOutputConfiguration) (media.GetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationResponse media.GetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputConfiguration")
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfiguration")
	}
	return reply.Body.GetAudioOutputConfigurationResponse, nil
}
