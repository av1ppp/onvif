// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetStorageConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationsResponse.
func Call_GetStorageConfigurations(ctx context.Context, dev *onvif.Device, request device.GetStorageConfigurations) (device.GetStorageConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationsResponse device.GetStorageConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStorageConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfigurations")
		return reply.Body.GetStorageConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetStorageConfigurations")
	}
}
