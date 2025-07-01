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

// Call_GetOSDs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDsResponse.
func Call_GetOSDs(ctx context.Context, dev *onvif.Device, request media.GetOSDs) (media.GetOSDsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDsResponse media.GetOSDsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetOSDsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSDs")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetOSDsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSDs")
	}
	return reply.Body.GetOSDsResponse, nil
}

// CallWithLogging_GetOSDs works like Call_GetOSDs but also logs the response body.
func CallWithLogging_GetOSDs(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetOSDs) (media.GetOSDsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDsResponse media.GetOSDsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetOSDsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSDs")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetOSDs")
	if err != nil {
		return reply.Body.GetOSDsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSDs")
	}
	return reply.Body.GetOSDsResponse, nil
}
