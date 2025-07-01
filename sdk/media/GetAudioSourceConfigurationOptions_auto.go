// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioSourceConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetAudioSourceConfigurationOptionsResponse.
func GetAudioSourceConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioSourceConfigurationOptions]) (media.GetAudioSourceConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationOptionsResponse media.GetAudioSourceConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAudioSourceConfigurationOptions")
	}
	return reply.Body.GetAudioSourceConfigurationOptionsResponse, nil
}
