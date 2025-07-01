// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCompatibleVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoEncoderConfigurationsResponse.
func Call_GetCompatibleVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoEncoderConfigurations) (media.GetCompatibleVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoEncoderConfigurationsResponse media.GetCompatibleVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleVideoEncoderConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoEncoderConfigurations")
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleVideoEncoderConfigurations")
	}
}
