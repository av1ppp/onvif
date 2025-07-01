// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GeoMove forwards the call to dev.CallMethod() then parses the payload of the reply as a GeoMoveResponse.
func Call_GeoMove(ctx context.Context, dev *onvif.Device, request ptz.GeoMove) (ptz.GeoMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GeoMoveResponse ptz.GeoMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GeoMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GeoMove")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GeoMove")
		return reply.Body.GeoMoveResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GeoMove")
	}
}
