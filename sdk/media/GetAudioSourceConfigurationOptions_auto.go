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

// Call_GetAudioSourceConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationOptionsResponse.
func Call_GetAudioSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioSourceConfigurationOptions) (media.GetAudioSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationOptionsResponse media.GetAudioSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	}
}

// CallWithLogging_GetAudioSourceConfigurationOptions works like Call_GetAudioSourceConfigurationOptions but also logs the response body.
func CallWithLogging_GetAudioSourceConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioSourceConfigurationOptions) (media.GetAudioSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationOptionsResponse media.GetAudioSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioSourceConfigurationOptions")
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	}
}
