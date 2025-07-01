// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	}
	return reply.Body.RemoveVideoAnalyticsConfigurationResponse, nil
}

// CallWithLogging_RemoveVideoAnalyticsConfiguration works like Call_RemoveVideoAnalyticsConfiguration but also logs the response body.
func CallWithLogging_RemoveVideoAnalyticsConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemoveVideoAnalyticsConfiguration) (media.RemoveVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoAnalyticsConfigurationResponse media.RemoveVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveVideoAnalyticsConfiguration")
	if err != nil {
		return reply.Body.RemoveVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoAnalyticsConfiguration")
	}
	return reply.Body.RemoveVideoAnalyticsConfigurationResponse, nil
}
