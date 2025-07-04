// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetConfigurationOptions forwards the call to onvif.Do then parses the payload of the reply as a GetConfigurationOptionsResponse.
func GetConfigurationOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetConfigurationOptions]) (ptz.GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse ptz.GetConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetConfigurationOptionsResponse, err
}
