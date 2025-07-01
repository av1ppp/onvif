// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetVideoSourceConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetVideoSourceConfigurationsResponse.
func GetVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetVideoSourceConfigurations]) (media.GetVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationsResponse media.GetVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetVideoSourceConfigurationsResponse, err
}
