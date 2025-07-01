// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// StopMulticastStreaming forwards the call to onvif.Do then parses the payload of the reply as a StopMulticastStreamingResponse.
func StopMulticastStreaming(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.StopMulticastStreaming]) (media.StopMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopMulticastStreamingResponse media.StopMulticastStreamingResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.StopMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StopMulticastStreaming")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StopMulticastStreaming")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.StopMulticastStreamingResponse, err
}
