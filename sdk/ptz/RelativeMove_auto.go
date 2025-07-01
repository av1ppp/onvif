// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// RelativeMove forwards the call to onvif.Do then parses the payload of the reply as a RelativeMoveResponse.
func RelativeMove(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.RelativeMove]) (ptz.RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse ptz.RelativeMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RelativeMove")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RelativeMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RelativeMove")
	}
	return reply.Body.RelativeMoveResponse, nil
}
