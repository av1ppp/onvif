// Code generated : DO NOT EDIT.

package event

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// PullMessages forwards the call to onvif.Do then parses the payload of the reply as a PullMessagesResponse.
func PullMessages(ctx context.Context, dev *onvif.Device, request *onvif.Req[event.PullMessages]) (event.PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.PullMessagesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "PullMessages")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "PullMessages")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.PullMessagesResponse, err
}
