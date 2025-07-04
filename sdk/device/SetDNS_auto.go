// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetDNS forwards the call to onvif.Do then parses the payload of the reply as a SetDNSResponse.
func SetDNS(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetDNS]) (device.SetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDNSResponse device.SetDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDNS")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDNS")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetDNSResponse, err
}
