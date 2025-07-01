// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// DeleteOSD forwards the call to onvif.Do then parses the payload of the reply as a DeleteOSDResponse.
func DeleteOSD(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.DeleteOSD]) (media.DeleteOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteOSDResponse media.DeleteOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteOSD")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteOSD")
	}
	return reply.Body.DeleteOSDResponse, nil
}
