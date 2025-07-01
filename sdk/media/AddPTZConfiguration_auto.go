// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// AddPTZConfiguration forwards the call to onvif.Do then parses the payload of the reply as a AddPTZConfigurationResponse.
func AddPTZConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.AddPTZConfiguration]) (media.AddPTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddPTZConfigurationResponse media.AddPTZConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.AddPTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddPTZConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddPTZConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AddPTZConfigurationResponse, err
}
