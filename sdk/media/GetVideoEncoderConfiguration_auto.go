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

// Call_GetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationResponse.
func Call_GetVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetVideoEncoderConfiguration) (media.GetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationResponse media.GetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfiguration")
	}
	return reply.Body.GetVideoEncoderConfigurationResponse, nil
}

// CallWithLogging_GetVideoEncoderConfiguration works like Call_GetVideoEncoderConfiguration but also logs the response body.
func CallWithLogging_GetVideoEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoEncoderConfiguration) (media.GetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationResponse media.GetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoEncoderConfiguration")
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoEncoderConfiguration")
	}
	return reply.Body.GetVideoEncoderConfigurationResponse, nil
}
