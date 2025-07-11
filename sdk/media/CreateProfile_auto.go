// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// CreateProfile forwards the call to onvif.Do then parses the payload of the reply as a CreateProfileResponse.
func CreateProfile(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.CreateProfile]) (media.CreateProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateProfileResponse media.CreateProfileResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.CreateProfileResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateProfile")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateProfile")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateProfileResponse, err
}
