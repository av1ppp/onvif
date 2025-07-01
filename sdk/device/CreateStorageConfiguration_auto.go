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

// Call_CreateStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateStorageConfigurationResponse.
func Call_CreateStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.CreateStorageConfiguration) (device.CreateStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateStorageConfigurationResponse device.CreateStorageConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateStorageConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateStorageConfiguration")
	}
	return reply.Body.CreateStorageConfigurationResponse, nil
}

// CallWithLogging_CreateStorageConfiguration works like Call_CreateStorageConfiguration but also logs the response body.
func CallWithLogging_CreateStorageConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.CreateStorageConfiguration) (device.CreateStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateStorageConfigurationResponse device.CreateStorageConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateStorageConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateStorageConfiguration")
	if err != nil {
		return reply.Body.CreateStorageConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateStorageConfiguration")
	}
	return reply.Body.CreateStorageConfigurationResponse, nil
}
