// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetServiceCapabilities forwards the call to onvif.Do then parses the payload of the reply as a GetServiceCapabilitiesResponse.
func GetServiceCapabilities(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetServiceCapabilities]) (media.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse media.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetServiceCapabilities")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetServiceCapabilities")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetServiceCapabilitiesResponse, err
}
