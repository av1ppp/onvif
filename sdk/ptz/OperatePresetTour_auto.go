// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// OperatePresetTour forwards the call to onvif.Do then parses the payload of the reply as a OperatePresetTourResponse.
func OperatePresetTour(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.OperatePresetTour]) (ptz.OperatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			OperatePresetTourResponse ptz.OperatePresetTourResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.OperatePresetTourResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "OperatePresetTour")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "OperatePresetTour")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.OperatePresetTourResponse, err
}
