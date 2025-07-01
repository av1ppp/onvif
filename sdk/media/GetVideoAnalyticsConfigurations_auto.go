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

// Call_GetVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoAnalyticsConfigurationsResponse.
func Call_GetVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoAnalyticsConfigurations) (media.GetVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationsResponse media.GetVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	}
	return reply.Body.GetVideoAnalyticsConfigurationsResponse, nil
}

// CallWithLogging_GetVideoAnalyticsConfigurations works like Call_GetVideoAnalyticsConfigurations but also logs the response body.
func CallWithLogging_GetVideoAnalyticsConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoAnalyticsConfigurations) (media.GetVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationsResponse media.GetVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoAnalyticsConfigurations")
	if err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoAnalyticsConfigurations")
	}
	return reply.Body.GetVideoAnalyticsConfigurationsResponse, nil
}
