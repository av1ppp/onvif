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

// Call_GetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationResponse.
func Call_GetAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioSourceConfiguration) (media.GetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationResponse media.GetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfiguration")
	}
}

// CallWithLogging_GetAudioSourceConfiguration works like Call_GetAudioSourceConfiguration but also logs the response body.
func CallWithLogging_GetAudioSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetAudioSourceConfiguration) (media.GetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationResponse media.GetAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioSourceConfiguration")
		return reply.Body.GetAudioSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfiguration")
	}
}
