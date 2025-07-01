// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoAnalyticsConfigurationsResponse.
func Call_GetVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoAnalyticsConfigurations) (media.GetVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationsResponse media.GetVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoAnalyticsConfigurations")
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	}
}
