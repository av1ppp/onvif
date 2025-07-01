// Code generated : DO NOT EDIT.

package media

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

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	}
	return reply.Body.GetMetadataConfigurationOptionsResponse, nil
}
