// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetSystemFactoryDefault forwards the call to onvif.Do then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func SetSystemFactoryDefault(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetSystemFactoryDefault]) (device.SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse device.SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSystemFactoryDefault")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSystemFactoryDefault")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetSystemFactoryDefaultResponse, err
}
