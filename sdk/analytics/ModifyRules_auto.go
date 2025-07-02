// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// ModifyRules forwards the call to onvif.Do then parses the payload of the reply as a ModifyRulesResponse.
func ModifyRules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.ModifyRules]) (analytics.ModifyRulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyRulesResponse analytics.ModifyRulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.ModifyRulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ModifyRules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ModifyRules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.ModifyRulesResponse, err
}
