// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetStorageConfiguration forwards the call to onvif.Do then parses the payload of the reply as a SetStorageConfigurationResponse.
func SetStorageConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetStorageConfiguration]) (device.SetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetStorageConfigurationResponse device.SetStorageConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetStorageConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetStorageConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetStorageConfigurationResponse, err
}
