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

// Call_GetOSDOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDOptionsResponse.
func Call_GetOSDOptions(ctx context.Context, dev *onvif.Device, request media.GetOSDOptions) (media.GetOSDOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDOptionsResponse media.GetOSDOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetOSDOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSDOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetOSDOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSDOptions")
	}
	return reply.Body.GetOSDOptionsResponse, nil
}

// CallWithLogging_GetOSDOptions works like Call_GetOSDOptions but also logs the response body.
func CallWithLogging_GetOSDOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetOSDOptions) (media.GetOSDOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDOptionsResponse media.GetOSDOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetOSDOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSDOptions")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetOSDOptions")
	if err != nil {
		return reply.Body.GetOSDOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSDOptions")
	}
	return reply.Body.GetOSDOptionsResponse, nil
}
