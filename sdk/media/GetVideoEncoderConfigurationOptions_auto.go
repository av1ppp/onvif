// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetVideoEncoderConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetVideoEncoderConfigurationOptionsResponse.
func GetVideoEncoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetVideoEncoderConfigurationOptions]) (media.GetVideoEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationOptionsResponse media.GetVideoEncoderConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurationOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoEncoderConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetVideoEncoderConfigurationOptionsResponse, err
}
