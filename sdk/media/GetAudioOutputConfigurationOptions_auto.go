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

// Call_GetAudioOutputConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationOptionsResponse.
func Call_GetAudioOutputConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfigurationOptions) (media.GetAudioOutputConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationOptionsResponse media.GetAudioOutputConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfigurationOptions")
	}
}

// CallWithLogging_GetAudioOutputConfigurationOptions works like Call_GetAudioOutputConfigurationOptions but also logs the response body.
func CallWithLogging_GetAudioOutputConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioOutputConfigurationOptions) (media.GetAudioOutputConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationOptionsResponse media.GetAudioOutputConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputConfigurationOptions")
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputConfigurationOptions")
	}
}
