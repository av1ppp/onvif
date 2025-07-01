// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_RemoveVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoAnalyticsConfigurationResponse.
func Call_RemoveVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveVideoAnalyticsConfiguration) (media.RemoveVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoAnalyticsConfigurationResponse media.RemoveVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoAnalyticsConfiguration")
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	}
}
