// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetPresets forwards the call to onvif.Do then parses the payload of the reply as a GetPresetsResponse.
func GetPresets(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetPresets]) (ptz.GetPresetsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetsResponse ptz.GetPresetsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetPresetsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresets")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPresets")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetPresetsResponse, err
}
