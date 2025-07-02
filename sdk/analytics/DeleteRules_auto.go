// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// DeleteRules forwards the call to onvif.Do then parses the payload of the reply as a DeleteRulesResponse.
func DeleteRules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.DeleteRules]) (analytics.DeleteRulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteRulesResponse analytics.DeleteRulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.DeleteRulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteRules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteRules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.DeleteRulesResponse, err
}
