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

// Call_StopMulticastStreaming forwards the call to dev.CallMethod() then parses the payload of the reply as a StopMulticastStreamingResponse.
func Call_StopMulticastStreaming(ctx context.Context, dev *onvif.Device, request media.StopMulticastStreaming) (media.StopMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopMulticastStreamingResponse media.StopMulticastStreamingResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StopMulticastStreaming")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.StopMulticastStreamingResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StopMulticastStreaming")
	}
}

// CallWithLogging_StopMulticastStreaming works like Call_StopMulticastStreaming but also logs the response body.
func CallWithLogging_StopMulticastStreaming(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.StopMulticastStreaming) (media.StopMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopMulticastStreamingResponse media.StopMulticastStreamingResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StopMulticastStreaming")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StopMulticastStreaming")
		return reply.Body.StopMulticastStreamingResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StopMulticastStreaming")
	}
}
