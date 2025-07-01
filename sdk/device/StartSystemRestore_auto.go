// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// StartSystemRestore forwards the call to onvif.Do then parses the payload of the reply as a StartSystemRestoreResponse.
func StartSystemRestore(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.StartSystemRestore]) (device.StartSystemRestoreResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartSystemRestoreResponse device.StartSystemRestoreResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "StartSystemRestore")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "StartSystemRestore")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.StartSystemRestoreResponse, err
}
