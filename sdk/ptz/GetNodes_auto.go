// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetNodes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNodesResponse.
func Call_GetNodes(ctx context.Context, dev *onvif.Device, request ptz.GetNodes) (ptz.GetNodesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodesResponse ptz.GetNodesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNodes")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNodes")
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNodes")
	}
}
