// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetVideoSourceConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetVideoSourceConfigurationOptionsResponse.
func GetVideoSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetVideoSourceConfigurationOptions]) (media.GetVideoSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationOptionsResponse media.GetVideoSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceConfigurationOptions")
	}
	return reply.Body.GetVideoSourceConfigurationOptionsResponse, nil
}
