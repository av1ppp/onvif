// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// CreateOSD forwards the call to onvif.Do then parses the payload of the reply as a CreateOSDResponse.
func CreateOSD(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.CreateOSD]) (media.CreateOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateOSDResponse media.CreateOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.CreateOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateOSD")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateOSD")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateOSDResponse, err
}
