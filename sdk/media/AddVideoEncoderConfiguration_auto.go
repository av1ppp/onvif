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

// Call_AddVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoEncoderConfigurationResponse.
func Call_AddVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoEncoderConfiguration) (media.AddVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoEncoderConfigurationResponse media.AddVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoEncoderConfiguration")
	}
}

// CallWithLogging_AddVideoEncoderConfiguration works like Call_AddVideoEncoderConfiguration but also logs the response body.
func CallWithLogging_AddVideoEncoderConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddVideoEncoderConfiguration) (media.AddVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoEncoderConfigurationResponse media.AddVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoEncoderConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddVideoEncoderConfiguration")
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoEncoderConfiguration")
	}
}
