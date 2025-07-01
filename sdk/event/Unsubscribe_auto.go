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

// Call_Unsubscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a UnsubscribeResponse.
func Call_Unsubscribe(ctx context.Context, dev *onvif.Device, request event.Unsubscribe) (event.UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse event.UnsubscribeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Unsubscribe")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Unsubscribe")
	}
	return reply.Body.UnsubscribeResponse, nil
}

// CallWithLogging_Unsubscribe works like Call_Unsubscribe but also logs the response body.
func CallWithLogging_Unsubscribe(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request event.Unsubscribe) (event.UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse event.UnsubscribeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Unsubscribe")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "Unsubscribe")
	if err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Unsubscribe")
	}
	return reply.Body.UnsubscribeResponse, nil
}
