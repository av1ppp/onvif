// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetSupportedAnalyticsModules forwards the call to onvif.Do then parses the payload of the reply as a GetSupportedAnalyticsModulesResponse.
func GetSupportedAnalyticsModules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetSupportedAnalyticsModules]) (analytics.GetSupportedAnalyticsModulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSupportedAnalyticsModulesResponse analytics.GetSupportedAnalyticsModulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetSupportedAnalyticsModulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSupportedAnalyticsModules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSupportedAnalyticsModules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetSupportedAnalyticsModulesResponse, err
}
