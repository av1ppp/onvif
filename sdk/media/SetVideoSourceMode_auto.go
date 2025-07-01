// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// SetVideoSourceMode forwards the call to onvif.Do then parses the payload of the reply as a SetVideoSourceModeResponse.
func SetVideoSourceMode(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.SetVideoSourceMode]) (media.SetVideoSourceModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceModeResponse media.SetVideoSourceModeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetVideoSourceModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceMode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoSourceMode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetVideoSourceModeResponse, err
}
