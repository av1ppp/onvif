// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetWsdlUrl forwards the call to onvif.Do then parses the payload of the reply as a GetWsdlUrlResponse.
func GetWsdlUrl(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetWsdlUrl]) (device.GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse device.GetWsdlUrlResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetWsdlUrl")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetWsdlUrl")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GetWsdlUrlResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetWsdlUrl")
	}
	return reply.Body.GetWsdlUrlResponse, nil
}
