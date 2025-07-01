// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// GetNodes forwards the call to onvif.Do then parses the payload of the reply as a GetNodesResponse.
func GetNodes(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.GetNodes]) (ptz.GetNodesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodesResponse ptz.GetNodesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNodes")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNodes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNodes")
	}
	return reply.Body.GetNodesResponse, nil
}
