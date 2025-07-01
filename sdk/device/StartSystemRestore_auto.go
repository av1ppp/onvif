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

// Call_StartSystemRestore forwards the call to dev.CallMethod() then parses the payload of the reply as a StartSystemRestoreResponse.
func Call_StartSystemRestore(ctx context.Context, dev *onvif.Device, request device.StartSystemRestore) (device.StartSystemRestoreResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartSystemRestoreResponse device.StartSystemRestoreResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartSystemRestore")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartSystemRestore")
	}
}

// CallWithLogging_StartSystemRestore works like Call_StartSystemRestore but also logs the response body.
func CallWithLogging_StartSystemRestore(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.StartSystemRestore) (device.StartSystemRestoreResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartSystemRestoreResponse device.StartSystemRestoreResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartSystemRestore")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartSystemRestore")
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartSystemRestore")
	}
}
