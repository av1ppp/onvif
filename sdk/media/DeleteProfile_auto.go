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

// Call_DeleteProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteProfileResponse.
func Call_DeleteProfile(ctx context.Context, dev *onvif.Device, request media.DeleteProfile) (media.DeleteProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteProfileResponse media.DeleteProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteProfile")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.DeleteProfileResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteProfile")
	}
}

// CallWithLogging_DeleteProfile works like Call_DeleteProfile but also logs the response body.
func CallWithLogging_DeleteProfile(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.DeleteProfile) (media.DeleteProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteProfileResponse media.DeleteProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteProfile")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteProfile")
		return reply.Body.DeleteProfileResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteProfile")
	}
}
