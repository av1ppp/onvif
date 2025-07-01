// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetCompatibleAudioEncoderConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetCompatibleAudioEncoderConfigurationsResponse.
func GetCompatibleAudioEncoderConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetCompatibleAudioEncoderConfigurations]) (media.GetCompatibleAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioEncoderConfigurationsResponse media.GetCompatibleAudioEncoderConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleAudioEncoderConfigurations")
	}
	return reply.Body.GetCompatibleAudioEncoderConfigurationsResponse, nil
}
