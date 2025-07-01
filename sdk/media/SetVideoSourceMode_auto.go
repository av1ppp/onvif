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

// Call_SetVideoSourceMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceModeResponse.
func Call_SetVideoSourceMode(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceMode) (media.SetVideoSourceModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceModeResponse media.SetVideoSourceModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	}
	return reply.Body.SetVideoSourceModeResponse, nil
}

// CallWithLogging_SetVideoSourceMode works like Call_SetVideoSourceMode but also logs the response body.
func CallWithLogging_SetVideoSourceMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetVideoSourceMode) (media.SetVideoSourceModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceModeResponse media.SetVideoSourceModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoSourceMode")
	if err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	}
	return reply.Body.SetVideoSourceModeResponse, nil
}
