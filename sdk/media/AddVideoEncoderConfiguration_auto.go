// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// AddVideoEncoderConfiguration forwards the call to onvif.Do then parses the payload of the reply as a AddVideoEncoderConfigurationResponse.
func AddVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.AddVideoEncoderConfiguration]) (media.AddVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoEncoderConfigurationResponse media.AddVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.AddVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoEncoderConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddVideoEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AddVideoEncoderConfigurationResponse, err
}
