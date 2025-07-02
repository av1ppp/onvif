// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetCompatibleVideoEncoderConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetCompatibleVideoEncoderConfigurationsResponse.
func GetCompatibleVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetCompatibleVideoEncoderConfigurations]) (media.GetCompatibleVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoEncoderConfigurationsResponse media.GetCompatibleVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleVideoEncoderConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleVideoEncoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, err
}
