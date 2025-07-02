// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// RemoveScopes forwards the call to onvif.Do then parses the payload of the reply as a RemoveScopesResponse.
func RemoveScopes(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.RemoveScopes]) (device.RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse device.RemoveScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.RemoveScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveScopes")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.RemoveScopesResponse, err
}
