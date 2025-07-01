// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// StartMulticastStreaming forwards the call to onvif.Do then parses the payload of the reply as a StartMulticastStreamingResponse.
func StartMulticastStreaming(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.StartMulticastStreaming]) (media.StartMulticastStreamingResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartMulticastStreamingResponse media.StartMulticastStreamingResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.StartMulticastStreamingResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartMulticastStreaming")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartMulticastStreaming")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.StartMulticastStreamingResponse, err
}
