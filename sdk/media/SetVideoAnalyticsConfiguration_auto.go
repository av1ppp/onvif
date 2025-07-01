// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetVideoAnalyticsConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetVideoAnalyticsConfigurationResponse.
func SetVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetVideoAnalyticsConfiguration]) (media.SetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoAnalyticsConfigurationResponse media.SetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoAnalyticsConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoAnalyticsConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetVideoAnalyticsConfigurationResponse, err
}
