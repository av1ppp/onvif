// Code generated : DO NOT EDIT.

package event

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/event"
	"github.com/av1ppp/onvif/errors"
)

// GetEventProperties forwards the call to onvif.Do then parses the payload of the reply as a GetEventPropertiesResponse.
func GetEventProperties(ctx context.Context, dev *onvif.Device, request *onvif.Req[event.GetEventProperties]) (event.GetEventPropertiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEventPropertiesResponse event.GetEventPropertiesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetEventPropertiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetEventProperties")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetEventPropertiesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetEventProperties")
	}
	return reply.Body.GetEventPropertiesResponse, nil
}
