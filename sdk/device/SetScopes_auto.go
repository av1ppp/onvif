// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetScopes forwards the call to onvif.Do then parses the payload of the reply as a SetScopesResponse.
func SetScopes(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetScopes]) (device.SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse device.SetScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetScopes")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetScopesResponse, err
}
