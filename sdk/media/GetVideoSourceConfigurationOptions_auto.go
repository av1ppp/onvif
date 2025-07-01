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

// Call_GetVideoSourceConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationOptionsResponse.
func Call_GetVideoSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceConfigurationOptions) (media.GetVideoSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationOptionsResponse media.GetVideoSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	}
}

// CallWithLogging_GetVideoSourceConfigurationOptions works like Call_GetVideoSourceConfigurationOptions but also logs the response body.
func CallWithLogging_GetVideoSourceConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoSourceConfigurationOptions) (media.GetVideoSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationOptionsResponse media.GetVideoSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoSourceConfigurationOptions")
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	}
}
