// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// SetPreset forwards the call to onvif.Do then parses the payload of the reply as a SetPresetResponse.
func SetPreset(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.SetPreset]) (ptz.SetPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetPresetResponse ptz.SetPresetResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetPreset")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetPresetResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetPreset")
	}
	return reply.Body.SetPresetResponse, nil
}
