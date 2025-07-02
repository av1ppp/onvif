// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// CreateAnalyticsModules forwards the call to onvif.Do then parses the payload of the reply as a CreateAnalyticsModulesResponse.
func CreateAnalyticsModules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.CreateAnalyticsModules]) (analytics.CreateAnalyticsModulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateAnalyticsModulesResponse analytics.CreateAnalyticsModulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.CreateAnalyticsModulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateAnalyticsModules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateAnalyticsModules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateAnalyticsModulesResponse, err
}
