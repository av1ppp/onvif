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

// Call_GetProfiles forwards the call to dev.CallMethod() then parses the payload of the reply as a GetProfilesResponse.
func Call_GetProfiles(ctx context.Context, dev *onvif.Device, request media.GetProfiles) (media.GetProfilesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfilesResponse media.GetProfilesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfiles")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfiles")
	}
}

// CallWithLogging_GetProfiles works like Call_GetProfiles but also logs the response body.
func CallWithLogging_GetProfiles(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetProfiles) (media.GetProfilesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfilesResponse media.GetProfilesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfiles")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetProfiles")
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfiles")
	}
}
