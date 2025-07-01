// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoEncoderConfigurationResponse.
func Call_RemoveVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveVideoEncoderConfiguration) (media.RemoveVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoEncoderConfigurationResponse media.RemoveVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoEncoderConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoEncoderConfiguration")
		return reply.Body.RemoveVideoEncoderConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoEncoderConfiguration")
	}
}
