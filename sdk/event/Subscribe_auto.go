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

// Call_Subscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a SubscribeResponse.
func Call_Subscribe(ctx context.Context, dev *onvif.Device, request event.Subscribe) (event.SubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SubscribeResponse event.SubscribeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Subscribe")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SubscribeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Subscribe")
	}
	return reply.Body.SubscribeResponse, nil
}

// CallWithLogging_Subscribe works like Call_Subscribe but also logs the response body.
func CallWithLogging_Subscribe(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request event.Subscribe) (event.SubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SubscribeResponse event.SubscribeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Subscribe")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "Subscribe")
	if err != nil {
		return reply.Body.SubscribeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Subscribe")
	}
	return reply.Body.SubscribeResponse, nil
}
