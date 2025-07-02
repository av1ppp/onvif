// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetAnalyticsModules forwards the call to onvif.Do then parses the payload of the reply as a GetAnalyticsModulesResponse.
func GetAnalyticsModules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetAnalyticsModules]) (analytics.GetAnalyticsModulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAnalyticsModulesResponse analytics.GetAnalyticsModulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetAnalyticsModulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAnalyticsModules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAnalyticsModules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAnalyticsModulesResponse, err
}
