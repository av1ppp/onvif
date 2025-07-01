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

// Call_SetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioSourceConfigurationResponse.
func Call_SetAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioSourceConfiguration) (media.SetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioSourceConfigurationResponse media.SetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioSourceConfiguration")
	}
}

// CallWithLogging_SetAudioSourceConfiguration works like Call_SetAudioSourceConfiguration but also logs the response body.
func CallWithLogging_SetAudioSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetAudioSourceConfiguration) (media.SetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioSourceConfigurationResponse media.SetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioSourceConfiguration")
		return reply.Body.SetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioSourceConfiguration")
	}
}
