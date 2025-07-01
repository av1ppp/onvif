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

// Call_GetSystemBackup forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemBackupResponse.
func Call_GetSystemBackup(ctx context.Context, dev *onvif.Device, request device.GetSystemBackup) (device.GetSystemBackupResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemBackupResponse device.GetSystemBackupResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemBackup")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemBackup")
	}
	return reply.Body.GetSystemBackupResponse, nil
}

// CallWithLogging_GetSystemBackup works like Call_GetSystemBackup but also logs the response body.
func CallWithLogging_GetSystemBackup(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetSystemBackup) (device.GetSystemBackupResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemBackupResponse device.GetSystemBackupResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemBackup")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemBackup")
	if err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSystemBackup")
	}
	return reply.Body.GetSystemBackupResponse, nil
}
