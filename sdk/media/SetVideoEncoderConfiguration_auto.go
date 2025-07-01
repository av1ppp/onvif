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

// Call_SetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoEncoderConfigurationResponse.
func Call_SetVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoEncoderConfiguration) (media.SetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoEncoderConfigurationResponse media.SetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoEncoderConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoEncoderConfiguration")
	}
	return reply.Body.SetVideoEncoderConfigurationResponse, nil
}

// CallWithLogging_SetVideoEncoderConfiguration works like Call_SetVideoEncoderConfiguration but also logs the response body.
func CallWithLogging_SetVideoEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetVideoEncoderConfiguration) (media.SetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoEncoderConfigurationResponse media.SetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoEncoderConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoEncoderConfiguration")
	if err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoEncoderConfiguration")
	}
	return reply.Body.SetVideoEncoderConfigurationResponse, nil
}
