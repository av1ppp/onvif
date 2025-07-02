// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetRules forwards the call to onvif.Do then parses the payload of the reply as a GetRulesResponse.
func GetRules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetRules]) (analytics.GetRulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRulesResponse analytics.GetRulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetRulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetRulesResponse, err
}
