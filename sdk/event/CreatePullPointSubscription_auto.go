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

// Call_CreatePullPointSubscription forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePullPointSubscriptionResponse.
func Call_CreatePullPointSubscription(ctx context.Context, dev *onvif.Device, request event.CreatePullPointSubscription) (event.CreatePullPointSubscriptionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePullPointSubscriptionResponse event.CreatePullPointSubscriptionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreatePullPointSubscription")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreatePullPointSubscription")
	}
}

// CallWithLogging_CreatePullPointSubscription works like Call_CreatePullPointSubscription but also logs the response body.
func CallWithLogging_CreatePullPointSubscription(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request event.CreatePullPointSubscription) (event.CreatePullPointSubscriptionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePullPointSubscriptionResponse event.CreatePullPointSubscriptionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreatePullPointSubscription")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreatePullPointSubscription")
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreatePullPointSubscription")
	}
}
