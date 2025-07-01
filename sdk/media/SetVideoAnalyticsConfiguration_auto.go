// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoAnalyticsConfigurationResponse.
func Call_SetVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoAnalyticsConfiguration) (media.SetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoAnalyticsConfigurationResponse media.SetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoAnalyticsConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoAnalyticsConfiguration")
		return reply.Body.SetVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoAnalyticsConfiguration")
	}
}
