// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetRuleOptions forwards the call to onvif.Do then parses the payload of the reply as a GetRuleOptionsResponse.
func GetRuleOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetRuleOptions]) (analytics.GetRuleOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRuleOptionsResponse analytics.GetRuleOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetRuleOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetRuleOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetRuleOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetRuleOptionsResponse, err
}
