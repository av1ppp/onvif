// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// RelativeMove forwards the call to onvif.Do then parses the payload of the reply as a RelativeMoveResponse.
func RelativeMove(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.RelativeMove]) (ptz.RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse ptz.RelativeMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RelativeMove")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RelativeMove")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RelativeMoveResponse, err
}
