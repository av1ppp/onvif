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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartSystemRestore")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartSystemRestore")
	}
	return reply.Body.StartSystemRestoreResponse, nil
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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartSystemRestore")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartSystemRestore")
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "StartSystemRestore")
	}
	return reply.Body.StartSystemRestoreResponse, nil
}
