// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// CreateStorageConfiguration forwards the call to onvif.Do then parses the payload of the reply as a CreateStorageConfigurationResponse.
func CreateStorageConfiguration(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.CreateStorageConfiguration]) (device.CreateStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateStorageConfigurationResponse device.CreateStorageConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateStorageConfiguration")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateStorageConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateStorageConfigurationResponse, err
}
