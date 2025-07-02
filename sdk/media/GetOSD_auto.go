// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetOSD forwards the call to onvif.Do then parses the payload of the reply as a GetOSDResponse.
func GetOSD(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetOSD]) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSD")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetOSD")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetOSDResponse, err
}
