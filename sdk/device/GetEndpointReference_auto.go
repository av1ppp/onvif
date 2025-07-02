// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetEndpointReference forwards the call to onvif.Do then parses the payload of the reply as a GetEndpointReferenceResponse.
func GetEndpointReference(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetEndpointReference]) (device.GetEndpointReferenceResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEndpointReferenceResponse device.GetEndpointReferenceResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetEndpointReferenceResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetEndpointReference")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetEndpointReference")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetEndpointReferenceResponse, err
}
