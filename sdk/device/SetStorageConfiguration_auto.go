// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetStorageConfigurationResponse.
func Call_SetStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.SetStorageConfiguration) (device.SetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetStorageConfigurationResponse device.SetStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetStorageConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetStorageConfiguration")
		return reply.Body.SetStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetStorageConfiguration")
	}
}
