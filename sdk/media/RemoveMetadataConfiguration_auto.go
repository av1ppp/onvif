// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveMetadataConfigurationResponse.
func Call_RemoveMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveMetadataConfiguration) (media.RemoveMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveMetadataConfigurationResponse media.RemoveMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveMetadataConfiguration")
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	}
}
