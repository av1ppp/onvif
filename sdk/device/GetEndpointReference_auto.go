// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetEndpointReference forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEndpointReferenceResponse.
func Call_GetEndpointReference(ctx context.Context, dev *onvif.Device, request device.GetEndpointReference) (device.GetEndpointReferenceResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEndpointReferenceResponse device.GetEndpointReferenceResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEndpointReferenceResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetEndpointReference")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetEndpointReferenceResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetEndpointReference")
	}
}

// CallWithLogging_GetEndpointReference works like Call_GetEndpointReference but also logs the response body.
func CallWithLogging_GetEndpointReference(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetEndpointReference) (device.GetEndpointReferenceResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEndpointReferenceResponse device.GetEndpointReferenceResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEndpointReferenceResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetEndpointReference")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetEndpointReference")
		return reply.Body.GetEndpointReferenceResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetEndpointReference")
	}
}
