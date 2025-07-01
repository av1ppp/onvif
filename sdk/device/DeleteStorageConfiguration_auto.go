// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	}
	return reply.Body.DeleteStorageConfigurationResponse, nil
}

// CallWithLogging_DeleteStorageConfiguration works like Call_DeleteStorageConfiguration but also logs the response body.
func CallWithLogging_DeleteStorageConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.DeleteStorageConfiguration) (device.DeleteStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteStorageConfigurationResponse device.DeleteStorageConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteStorageConfiguration")
	if err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteStorageConfiguration")
	}
	return reply.Body.DeleteStorageConfigurationResponse, nil
}
