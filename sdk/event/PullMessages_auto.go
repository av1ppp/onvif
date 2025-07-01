// Code generated : DO NOT EDIT.

package event

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// Call_PullMessages forwards the call to dev.CallMethod() then parses the payload of the reply as a PullMessagesResponse.
func Call_PullMessages(ctx context.Context, dev *onvif.Device, request event.PullMessages) (event.PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.PullMessagesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "PullMessages")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.PullMessagesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "PullMessages")
	}
}

// CallWithLogging_PullMessages works like Call_PullMessages but also logs the response body.
func CallWithLogging_PullMessages(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request event.PullMessages) (event.PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.PullMessagesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "PullMessages")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "PullMessages")
		return reply.Body.PullMessagesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "PullMessages")
	}
}
