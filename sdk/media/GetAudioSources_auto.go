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

// Call_GetAudioSources forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourcesResponse.
func Call_GetAudioSources(ctx context.Context, dev *onvif.Device, request media.GetAudioSources) (media.GetAudioSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourcesResponse media.GetAudioSourcesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioSourcesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSources")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioSourcesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSources")
	}
	return reply.Body.GetAudioSourcesResponse, nil
}

// CallWithLogging_GetAudioSources works like Call_GetAudioSources but also logs the response body.
func CallWithLogging_GetAudioSources(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioSources) (media.GetAudioSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourcesResponse media.GetAudioSourcesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioSourcesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSources")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioSources")
	if err != nil {
		return reply.Body.GetAudioSourcesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSources")
	}
	return reply.Body.GetAudioSourcesResponse, nil
}
