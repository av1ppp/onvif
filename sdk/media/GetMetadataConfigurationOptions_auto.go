// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetMetadataConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetMetadataConfigurationOptionsResponse.
func GetMetadataConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetMetadataConfigurationOptions]) (media.GetMetadataConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationOptionsResponse media.GetMetadataConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetMetadataConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetMetadataConfigurationOptionsResponse, err
}
