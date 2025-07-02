// Code generated : DO NOT EDIT.

package sdkanalytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// CreateRules forwards the call to onvif.Do then parses the payload of the reply as a CreateRulesResponse.
func CreateRules(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.CreateRules]) (analytics.CreateRulesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateRulesResponse analytics.CreateRulesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.CreateRulesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateRules")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateRules")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateRulesResponse, err
}
