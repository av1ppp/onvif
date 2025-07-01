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

// Call_SetAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioEncoderConfigurationResponse.
func Call_SetAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioEncoderConfiguration) (media.SetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioEncoderConfigurationResponse media.SetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioEncoderConfiguration")
	}
}

// CallWithLogging_SetAudioEncoderConfiguration works like Call_SetAudioEncoderConfiguration but also logs the response body.
func CallWithLogging_SetAudioEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetAudioEncoderConfiguration) (media.SetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioEncoderConfigurationResponse media.SetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioEncoderConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioEncoderConfiguration")
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioEncoderConfiguration")
	}
}
