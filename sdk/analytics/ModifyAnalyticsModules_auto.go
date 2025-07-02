// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// ModifyAnalyticsModules forwards the call to onvif.Do then parses the payload of the reply as a ModifyAnalyticsModulesResponse.
func ModifyAnalyticsModules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.ModifyAnalyticsModules]) (analytics.ModifyAnalyticsModulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyAnalyticsModulesResponse analytics.ModifyAnalyticsModulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.ModifyAnalyticsModulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyAnalyticsModules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ModifyAnalyticsModules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.ModifyAnalyticsModulesResponse, err
}
