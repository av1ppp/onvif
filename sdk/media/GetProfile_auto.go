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

// Call_GetProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a GetProfileResponse.
func Call_GetProfile(ctx context.Context, dev *onvif.Device, request media.GetProfile) (media.GetProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfileResponse media.GetProfileResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfile")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfile")
	}
	return reply.Body.GetProfileResponse, nil
}

// CallWithLogging_GetProfile works like Call_GetProfile but also logs the response body.
func CallWithLogging_GetProfile(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetProfile) (media.GetProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfileResponse media.GetProfileResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfile")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetProfile")
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfile")
	}
	return reply.Body.GetProfileResponse, nil
}
