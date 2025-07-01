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

// Call_SetAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioOutputConfigurationResponse.
func Call_SetAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioOutputConfiguration) (media.SetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioOutputConfigurationResponse media.SetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioOutputConfiguration")
	}
}

// CallWithLogging_SetAudioOutputConfiguration works like Call_SetAudioOutputConfiguration but also logs the response body.
func CallWithLogging_SetAudioOutputConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetAudioOutputConfiguration) (media.SetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioOutputConfigurationResponse media.SetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAudioOutputConfiguration")
		return reply.Body.SetAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAudioOutputConfiguration")
	}
}
