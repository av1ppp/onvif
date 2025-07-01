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

// Call_GetVideoSources forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourcesResponse.
func Call_GetVideoSources(ctx context.Context, dev *onvif.Device, request media.GetVideoSources) (media.GetVideoSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourcesResponse media.GetVideoSourcesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourcesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSources")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetVideoSourcesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSources")
	}
}

// CallWithLogging_GetVideoSources works like Call_GetVideoSources but also logs the response body.
func CallWithLogging_GetVideoSources(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoSources) (media.GetVideoSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourcesResponse media.GetVideoSourcesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourcesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSources")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoSources")
		return reply.Body.GetVideoSourcesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSources")
	}
}
