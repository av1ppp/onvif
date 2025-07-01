// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// AddScopes forwards the call to onvif.Do then parses the payload of the reply as a AddScopesResponse.
func AddScopes(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.AddScopes]) (device.AddScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddScopesResponse device.AddScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddScopes")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.AddScopesResponse, err
}
