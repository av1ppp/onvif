// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetDynamicDNS forwards the call to onvif.Do then parses the payload of the reply as a GetDynamicDNSResponse.
func GetDynamicDNS(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetDynamicDNS]) (device.GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse device.GetDynamicDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetDynamicDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDynamicDNS")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDynamicDNS")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetDynamicDNSResponse, err
}
