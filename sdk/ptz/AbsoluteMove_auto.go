// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// AbsoluteMove forwards the call to onvif.Do then parses the payload of the reply as a AbsoluteMoveResponse.
func AbsoluteMove(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.AbsoluteMove]) (ptz.AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse ptz.AbsoluteMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AbsoluteMove")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AbsoluteMove")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AbsoluteMoveResponse, err
}
