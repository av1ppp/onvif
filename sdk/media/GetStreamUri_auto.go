// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetStreamUri forwards the call to onvif.Do then parses the payload of the reply as a GetStreamUriResponse.
func GetStreamUri(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetStreamUri]) (media.GetStreamUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStreamUriResponse media.GetStreamUriResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetStreamUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStreamUri")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetStreamUri")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetStreamUriResponse, err
}
