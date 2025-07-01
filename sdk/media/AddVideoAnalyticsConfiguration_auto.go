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

// Call_AddVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoAnalyticsConfigurationResponse.
func Call_AddVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoAnalyticsConfiguration) (media.AddVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoAnalyticsConfigurationResponse media.AddVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoAnalyticsConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AddVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoAnalyticsConfiguration")
	}
	return reply.Body.AddVideoAnalyticsConfigurationResponse, nil
}

// CallWithLogging_AddVideoAnalyticsConfiguration works like Call_AddVideoAnalyticsConfiguration but also logs the response body.
func CallWithLogging_AddVideoAnalyticsConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddVideoAnalyticsConfiguration) (media.AddVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoAnalyticsConfigurationResponse media.AddVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.AddVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoAnalyticsConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddVideoAnalyticsConfiguration")
	if err != nil {
		return reply.Body.AddVideoAnalyticsConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoAnalyticsConfiguration")
	}
	return reply.Body.AddVideoAnalyticsConfigurationResponse, nil
}
