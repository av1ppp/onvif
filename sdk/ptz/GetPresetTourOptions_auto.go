// Code generated : DO NOT EDIT.

package sdkptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetPresetTourOptions forwards the call to onvif.Do then parses the payload of the reply as a GetPresetTourOptionsResponse.
func GetPresetTourOptions(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetPresetTourOptions]) (ptz.GetPresetTourOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourOptionsResponse ptz.GetPresetTourOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetPresetTourOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetPresetTourOptions")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetPresetTourOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetPresetTourOptionsResponse, err
}
