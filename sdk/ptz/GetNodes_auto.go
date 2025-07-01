// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/logx"

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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNodes")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNodes")
	}
	return reply.Body.GetNodesResponse, nil
}

// CallWithLogging_GetNodes works like Call_GetNodes but also logs the response body.
func CallWithLogging_GetNodes(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetNodes) (ptz.GetNodesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodesResponse ptz.GetNodesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNodes")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNodes")
	if err != nil {
		return reply.Body.GetNodesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNodes")
	}
	return reply.Body.GetNodesResponse, nil
}
