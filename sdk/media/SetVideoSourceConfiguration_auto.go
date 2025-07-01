// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetVideoSourceConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetVideoSourceConfigurationResponse.
func SetVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetVideoSourceConfiguration]) (media.SetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceConfigurationResponse media.SetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetVideoSourceConfigurationResponse, err
}
