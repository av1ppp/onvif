// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetServices forwards the call to onvif.Do then parses the payload of the reply as a GetServicesResponse.
func GetServices(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetServices]) (device.GetServicesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServicesResponse device.GetServicesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetServicesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServices")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetServices")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GetServicesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetServices")
	}
	return reply.Body.GetServicesResponse, nil
}
