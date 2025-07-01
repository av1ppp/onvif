// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetProfile forwards the call to onvif.Do then parses the payload of the reply as a GetProfileResponse.
func GetProfile(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetProfile]) (media.GetProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfileResponse media.GetProfileResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfile")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetProfile")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GetProfileResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfile")
	}
	return reply.Body.GetProfileResponse, nil
}
