// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetHostnameFromDHCP forwards the call to onvif.Do then parses the payload of the reply as a SetHostnameFromDHCPResponse.
func SetHostnameFromDHCP(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetHostnameFromDHCP]) (device.SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse device.SetHostnameFromDHCPResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetHostnameFromDHCP")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetHostnameFromDHCPResponse, err
}
