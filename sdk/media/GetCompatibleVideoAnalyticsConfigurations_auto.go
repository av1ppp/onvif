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

// Call_GetCompatibleVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoAnalyticsConfigurationsResponse.
func Call_GetCompatibleVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoAnalyticsConfigurations) (media.GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoAnalyticsConfigurationsResponse media.GetCompatibleVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleVideoAnalyticsConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleVideoAnalyticsConfigurations")
	}
}

// CallWithLogging_GetCompatibleVideoAnalyticsConfigurations works like Call_GetCompatibleVideoAnalyticsConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleVideoAnalyticsConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetCompatibleVideoAnalyticsConfigurations) (media.GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoAnalyticsConfigurationsResponse media.GetCompatibleVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleVideoAnalyticsConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleVideoAnalyticsConfigurations")
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleVideoAnalyticsConfigurations")
	}
}
