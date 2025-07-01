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

// Call_GetAudioOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputsResponse.
func Call_GetAudioOutputs(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputs) (media.GetAudioOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputsResponse media.GetAudioOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputs")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputs")
	}
	return reply.Body.GetAudioOutputsResponse, nil
}

// CallWithLogging_GetAudioOutputs works like Call_GetAudioOutputs but also logs the response body.
func CallWithLogging_GetAudioOutputs(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioOutputs) (media.GetAudioOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputsResponse media.GetAudioOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputs")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputs")
	if err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioOutputs")
	}
	return reply.Body.GetAudioOutputsResponse, nil
}
