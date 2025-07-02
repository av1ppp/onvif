// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetSystemBackup forwards the call to onvif.Do then parses the payload of the reply as a GetSystemBackupResponse.
func GetSystemBackup(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetSystemBackup]) (device.GetSystemBackupResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemBackupResponse device.GetSystemBackupResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSystemBackup")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSystemBackup")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetSystemBackupResponse, err
}
