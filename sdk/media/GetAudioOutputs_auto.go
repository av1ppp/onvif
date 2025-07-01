// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetAudioOutputs forwards the call to onvif.Do then parses the payload of the reply as a GetAudioOutputsResponse.
func GetAudioOutputs(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetAudioOutputs]) (media.GetAudioOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputsResponse media.GetAudioOutputsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAudioOutputs")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAudioOutputs")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAudioOutputsResponse, err
}
