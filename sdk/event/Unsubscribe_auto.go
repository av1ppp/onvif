// Code generated : DO NOT EDIT.

package event

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// Unsubscribe forwards the call to onvif.Do then parses the payload of the reply as a UnsubscribeResponse.
func Unsubscribe(ctx context.Context, dev *onvif.Device, request *onvif.Req[event.Unsubscribe]) (event.UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse event.UnsubscribeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Unsubscribe")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "Unsubscribe")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.UnsubscribeResponse, err
}
