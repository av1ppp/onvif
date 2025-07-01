// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationResponse.
func Call_GetStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.GetStorageConfiguration) (device.GetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationResponse device.GetStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStorageConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfiguration")
		return reply.Body.GetStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetStorageConfiguration")
	}
}
