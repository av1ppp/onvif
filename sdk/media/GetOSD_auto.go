// Code generated : DO NOT EDIT.

package media

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

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetOSD")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetOSD")
	}
	return reply.Body.GetOSDResponse, nil
}
