// Code generated : DO NOT EDIT.

package sdkevent

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// CreatePullPointSubscription forwards the call to onvif.Do then parses the payload of the reply as a CreatePullPointSubscriptionResponse.
func CreatePullPointSubscription(ctx context.Context, dev *onvif.Device, request *onvif.Req[event.CreatePullPointSubscription]) (event.CreatePullPointSubscriptionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePullPointSubscriptionResponse event.CreatePullPointSubscriptionResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.CreatePullPointSubscriptionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreatePullPointSubscription")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreatePullPointSubscription")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreatePullPointSubscriptionResponse, err
}
