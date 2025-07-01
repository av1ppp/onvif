// Code generated : DO NOT EDIT.

package event

import (
	"context"

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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Unsubscribe")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Unsubscribe")
		return reply.Body.UnsubscribeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Unsubscribe")
	}
}
