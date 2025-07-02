// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetSynchronizationPoint forwards the call to onvif.Do then parses the payload of the reply as a SetSynchronizationPointResponse.
func SetSynchronizationPoint(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetSynchronizationPoint]) (media.SetSynchronizationPointResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSynchronizationPointResponse media.SetSynchronizationPointResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetSynchronizationPoint")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetSynchronizationPoint")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetSynchronizationPointResponse, err
}
