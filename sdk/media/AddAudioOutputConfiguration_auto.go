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

// Call_AddAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioOutputConfigurationResponse.
func Call_AddAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioOutputConfiguration) (media.AddAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioOutputConfigurationResponse media.AddAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioOutputConfiguration")
	}
}

// CallWithLogging_AddAudioOutputConfiguration works like Call_AddAudioOutputConfiguration but also logs the response body.
func CallWithLogging_AddAudioOutputConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddAudioOutputConfiguration) (media.AddAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioOutputConfigurationResponse media.AddAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioOutputConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddAudioOutputConfiguration")
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioOutputConfiguration")
	}
}
