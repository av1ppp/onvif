// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetSynchronizationPoint forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSynchronizationPointResponse.
func Call_SetSynchronizationPoint(ctx context.Context, dev *onvif.Device, request media.SetSynchronizationPoint) (media.SetSynchronizationPointResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSynchronizationPointResponse media.SetSynchronizationPointResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSynchronizationPoint")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSynchronizationPoint")
	}
	return reply.Body.SetSynchronizationPointResponse, nil
}

// CallWithLogging_SetSynchronizationPoint works like Call_SetSynchronizationPoint but also logs the response body.
func CallWithLogging_SetSynchronizationPoint(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetSynchronizationPoint) (media.SetSynchronizationPointResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSynchronizationPointResponse media.SetSynchronizationPointResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSynchronizationPoint")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSynchronizationPoint")
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetSynchronizationPoint")
	}
	return reply.Body.SetSynchronizationPointResponse, nil
}
