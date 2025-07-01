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

// Call_RemoveAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioSourceConfigurationResponse.
func Call_RemoveAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioSourceConfiguration) (media.RemoveAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioSourceConfigurationResponse media.RemoveAudioSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioSourceConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioSourceConfiguration")
	}
	return reply.Body.RemoveAudioSourceConfigurationResponse, nil
}

// CallWithLogging_RemoveAudioSourceConfiguration works like Call_RemoveAudioSourceConfiguration but also logs the response body.
func CallWithLogging_RemoveAudioSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemoveAudioSourceConfiguration) (media.RemoveAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioSourceConfigurationResponse media.RemoveAudioSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveAudioSourceConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveAudioSourceConfiguration")
	if err != nil {
		return reply.Body.RemoveAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveAudioSourceConfiguration")
	}
	return reply.Body.RemoveAudioSourceConfigurationResponse, nil
}
