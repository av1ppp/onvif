// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetOSD forwards the call to onvif.Do then parses the payload of the reply as a SetOSDResponse.
func SetOSD(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetOSD]) (media.SetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetOSDResponse media.SetOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetOSD")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetOSD")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetOSDResponse, err
}
