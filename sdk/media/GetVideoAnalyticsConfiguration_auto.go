// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoAnalyticsConfigurationResponse.
func Call_GetVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.GetVideoAnalyticsConfiguration) (media.GetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationResponse media.GetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoAnalyticsConfiguration")
		return reply.Body.GetVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfiguration")
	}
}
