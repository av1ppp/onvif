// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetHostname forwards the call to onvif.Do then parses the payload of the reply as a SetHostnameResponse.
func SetHostname(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetHostname]) (device.SetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameResponse device.SetHostnameResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostname")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetHostname")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetHostnameResponse, err
}
