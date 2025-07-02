// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetSupportedRules forwards the call to onvif.Do then parses the payload of the reply as a GetSupportedRulesResponse.
func GetSupportedRules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetSupportedRules]) (analytics.GetSupportedRulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSupportedRulesResponse analytics.GetSupportedRulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetSupportedRulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSupportedRules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSupportedRules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetSupportedRulesResponse, err
}
