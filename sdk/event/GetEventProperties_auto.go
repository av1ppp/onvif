// Code generated : DO NOT EDIT.

package event

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetEventProperties forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEventPropertiesResponse.
func Call_GetEventProperties(ctx context.Context, dev *onvif.Device, request event.GetEventProperties) (event.GetEventPropertiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEventPropertiesResponse event.GetEventPropertiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEventPropertiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetEventProperties")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetEventProperties")
		return reply.Body.GetEventPropertiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetEventProperties")
	}
}
