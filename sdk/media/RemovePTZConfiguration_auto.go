// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// RemovePTZConfiguration forwards the call to onvif.Do then parses the payload of the reply as a RemovePTZConfigurationResponse.
func RemovePTZConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.RemovePTZConfiguration]) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemovePTZConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemovePTZConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RemovePTZConfigurationResponse, err
}
