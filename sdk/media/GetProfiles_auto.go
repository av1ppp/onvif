// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetProfiles forwards the call to onvif.Do then parses the payload of the reply as a GetProfilesResponse.
func GetProfiles(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetProfiles]) (media.GetProfilesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetProfilesResponse media.GetProfilesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetProfiles")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetProfilesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetProfiles")
	}
	return reply.Body.GetProfilesResponse, nil
}
