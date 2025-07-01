// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetDot1XConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetDot1XConfigurationResponse.
func SetDot1XConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetDot1XConfiguration]) (device.SetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDot1XConfigurationResponse device.SetDot1XConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetDot1XConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDot1XConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDot1XConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetDot1XConfigurationResponse, err
}
