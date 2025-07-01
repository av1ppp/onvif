// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetHostname forwards the call to onvif.Do then parses the payload of the reply as a GetHostnameResponse.
func GetHostname(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetHostname]) (device.GetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetHostnameResponse device.GetHostnameResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetHostnameResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetHostname")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetHostname")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetHostnameResponse, err
}
