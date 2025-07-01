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

// Call_GetVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationsResponse.
func Call_GetVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoEncoderConfigurations) (media.GetVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationsResponse media.GetVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurations")
	}
}

// CallWithLogging_GetVideoEncoderConfigurations works like Call_GetVideoEncoderConfigurations but also logs the response body.
func CallWithLogging_GetVideoEncoderConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoEncoderConfigurations) (media.GetVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationsResponse media.GetVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoEncoderConfigurations")
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurations")
	}
}
