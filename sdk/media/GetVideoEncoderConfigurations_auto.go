// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetVideoEncoderConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetVideoEncoderConfigurationsResponse.
func GetVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetVideoEncoderConfigurations]) (media.GetVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationsResponse media.GetVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoEncoderConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoEncoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetVideoEncoderConfigurationsResponse, err
}
