// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetMetadataConfigurationResponse.
func Call_SetMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.SetMetadataConfiguration) (media.SetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetMetadataConfigurationResponse media.SetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetMetadataConfiguration")
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	}
}
