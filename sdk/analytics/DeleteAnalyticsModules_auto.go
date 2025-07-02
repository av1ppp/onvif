// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// DeleteAnalyticsModules forwards the call to onvif.Do then parses the payload of the reply as a DeleteAnalyticsModulesResponse.
func DeleteAnalyticsModules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.DeleteAnalyticsModules]) (analytics.DeleteAnalyticsModulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteAnalyticsModulesResponse analytics.DeleteAnalyticsModulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.DeleteAnalyticsModulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteAnalyticsModules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteAnalyticsModules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.DeleteAnalyticsModulesResponse, err
}
