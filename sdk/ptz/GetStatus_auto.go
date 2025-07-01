// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetStatus forwards the call to onvif.Do then parses the payload of the reply as a GetStatusResponse.
func GetStatus(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetStatus]) (ptz.GetStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStatusResponse ptz.GetStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStatus")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetStatus")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetStatusResponse, err
}
