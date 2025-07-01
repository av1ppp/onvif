// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_DeleteStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteStorageConfigurationResponse.
func Call_DeleteStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.DeleteStorageConfiguration) (device.DeleteStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteStorageConfigurationResponse device.DeleteStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteStorageConfiguration")
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	}
}
