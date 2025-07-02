// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// ContinuousMove forwards the call to onvif.Do then parses the payload of the reply as a ContinuousMoveResponse.
func ContinuousMove(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.ContinuousMove]) (ptz.ContinuousMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ContinuousMoveResponse ptz.ContinuousMoveResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.ContinuousMoveResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "ContinuousMove")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "ContinuousMove")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.ContinuousMoveResponse, err
}
