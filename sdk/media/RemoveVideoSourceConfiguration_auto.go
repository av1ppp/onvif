// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// RemoveVideoSourceConfiguration forwards the call to onvif.Do then parses the payload of the reply as a RemoveVideoSourceConfigurationResponse.
func RemoveVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.RemoveVideoSourceConfiguration]) (media.RemoveVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoSourceConfigurationResponse media.RemoveVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoSourceConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveVideoSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RemoveVideoSourceConfigurationResponse, err
}
