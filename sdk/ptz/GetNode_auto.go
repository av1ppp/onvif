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

// Call_GetNode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNodeResponse.
func Call_GetNode(ctx context.Context, dev *onvif.Device, request ptz.GetNode) (ptz.GetNodeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodeResponse ptz.GetNodeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetNodeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetNodeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNode")
	}
	return reply.Body.GetNodeResponse, nil
}

// CallWithLogging_GetNode works like Call_GetNode but also logs the response body.
func CallWithLogging_GetNode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetNode) (ptz.GetNodeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodeResponse ptz.GetNodeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetNodeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetNode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetNode")
	if err != nil {
		return reply.Body.GetNodeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetNode")
	}
	return reply.Body.GetNodeResponse, nil
}
