// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetNode forwards the call to onvif.Do then parses the payload of the reply as a GetNodeResponse.
func GetNode(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetNode]) (ptz.GetNodeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodeResponse ptz.GetNodeResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetNodeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNode")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNode")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetNodeResponse, err
}
