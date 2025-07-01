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

// Call_GetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDResponse.
func Call_GetOSD(ctx context.Context, dev *onvif.Device, request media.GetOSD) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSD")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSD")
	}
	return reply.Body.GetOSDResponse, nil
}

// CallWithLogging_GetOSD works like Call_GetOSD but also logs the response body.
func CallWithLogging_GetOSD(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetOSD) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSD")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetOSD")
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSD")
	}
	return reply.Body.GetOSDResponse, nil
}
