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

// Call_AddAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioSourceConfigurationResponse.
func Call_AddAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioSourceConfiguration) (media.AddAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioSourceConfigurationResponse media.AddAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioSourceConfiguration")
	}
}

// CallWithLogging_AddAudioSourceConfiguration works like Call_AddAudioSourceConfiguration but also logs the response body.
func CallWithLogging_AddAudioSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddAudioSourceConfiguration) (media.AddAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioSourceConfigurationResponse media.AddAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddAudioSourceConfiguration")
		return reply.Body.AddAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddAudioSourceConfiguration")
	}
}
