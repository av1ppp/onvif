// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetNTP forwards the call to onvif.Do then parses the payload of the reply as a GetNTPResponse.
func GetNTP(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetNTP]) (device.GetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNTPResponse device.GetNTPResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetNTPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNTP")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNTP")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetNTPResponse, err
}
