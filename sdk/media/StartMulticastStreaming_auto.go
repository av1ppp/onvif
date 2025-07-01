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

// Call_StartMulticastStreaming forwards the call to dev.CallMethod() then parses the payload of the reply as a StartMulticastStreamingResponse.
func Call_StartMulticastStreaming(ctx context.Context, dev *onvif.Device, request media.StartMulticastStreaming) (media.StartMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartMulticastStreamingResponse media.StartMulticastStreamingResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.StartMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartMulticastStreaming")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.StartMulticastStreamingResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartMulticastStreaming")
	}
	return reply.Body.StartMulticastStreamingResponse, nil
}

// CallWithLogging_StartMulticastStreaming works like Call_StartMulticastStreaming but also logs the response body.
func CallWithLogging_StartMulticastStreaming(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.StartMulticastStreaming) (media.StartMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartMulticastStreamingResponse media.StartMulticastStreamingResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.StartMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartMulticastStreaming")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartMulticastStreaming")
	if err != nil {
		return reply.Body.StartMulticastStreamingResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartMulticastStreaming")
	}
	return reply.Body.StartMulticastStreamingResponse, nil
}
