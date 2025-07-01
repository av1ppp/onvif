// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemovePTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePTZConfigurationResponse.
func Call_RemovePTZConfiguration(ctx context.Context, dev *onvif.Device, request media.RemovePTZConfiguration) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePTZConfiguration")
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	}
}
