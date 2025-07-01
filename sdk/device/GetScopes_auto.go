// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetScopes forwards the call to onvif.Do then parses the payload of the reply as a GetScopesResponse.
func GetScopes(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetScopes]) (device.GetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetScopesResponse device.GetScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetScopes")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetScopes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GetScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetScopes")
	}
	return reply.Body.GetScopesResponse, nil
}
