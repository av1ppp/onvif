// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetAnalyticsModuleOptions forwards the call to onvif.Do then parses the payload of the reply as a GetAnalyticsModuleOptionsResponse.
func GetAnalyticsModuleOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetAnalyticsModuleOptions]) (analytics.GetAnalyticsModuleOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAnalyticsModuleOptionsResponse analytics.GetAnalyticsModuleOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetAnalyticsModuleOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAnalyticsModuleOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAnalyticsModuleOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAnalyticsModuleOptionsResponse, err
}
