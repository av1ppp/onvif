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

// Call_GetVideoEncoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationOptionsResponse.
func Call_GetVideoEncoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetVideoEncoderConfigurationOptions) (media.GetVideoEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationOptionsResponse media.GetVideoEncoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetVideoEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurationOptions")
	}
}

// CallWithLogging_GetVideoEncoderConfigurationOptions works like Call_GetVideoEncoderConfigurationOptions but also logs the response body.
func CallWithLogging_GetVideoEncoderConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoEncoderConfigurationOptions) (media.GetVideoEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationOptionsResponse media.GetVideoEncoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoEncoderConfigurationOptions")
		return reply.Body.GetVideoEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurationOptions")
	}
}
