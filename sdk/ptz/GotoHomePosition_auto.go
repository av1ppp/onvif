// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GotoHomePosition forwards the call to onvif.Do then parses the payload of the reply as a GotoHomePositionResponse.
func GotoHomePosition(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GotoHomePosition]) (ptz.GotoHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoHomePositionResponse ptz.GotoHomePositionResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GotoHomePositionResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GotoHomePosition")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GotoHomePosition")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GotoHomePositionResponse, err
}
