// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetVideoAnalyticsConfigurations forwards the call to onvif.Do then parses the payload of the reply as a GetVideoAnalyticsConfigurationsResponse.
func GetVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetVideoAnalyticsConfigurations]) (media.GetVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationsResponse media.GetVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoAnalyticsConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetVideoAnalyticsConfigurationsResponse, err
}
